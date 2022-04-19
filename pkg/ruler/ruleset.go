package ruler

import (
	"bytes"
	// "crypto/sha256"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"sync"

	"github.com/controlplaneio/badrobot/pkg/rules"
	"github.com/ghodss/yaml"

	// "github.com/in-toto/in-toto-golang/in_toto"

	"github.com/thedevsaddam/gojsonq/v2"
	"go.uber.org/zap"
)

type Ruleset struct {
	Rules  []Rule
	logger *zap.SugaredLogger
}

type InvalidInputError struct {
}

func (e *InvalidInputError) Error() string {
	return "Invalid input"
}

func NewRuleset(logger *zap.SugaredLogger) *Ruleset {
	list := make([]Rule, 0)

	// OPR-R1-NS - default namespace
	defaultNamespaceRule := Rule{
		Predicate: rules.DefaultNamespace,
		ID:        "DefaultNamespace",
		Selector:  ".metadata .name == default .subjects .namespace == default",
		Reason:    "Operator is deployed into the default namespace.",
		Kinds:     []string{"Namespace", "Deployment", "ClusterRoleBinding"},
		Points:    -3,
	}
	list = append(list, defaultNamespaceRule)

	// OPR-R2-NS - kube-system namespace
	kubesystemNamespaceRule := Rule{
		Predicate: rules.KubeSystemNamespace,
		ID:        "KubeSystemNamespace",
		Selector:  ".metadata .name == kube-system .subjects .namespace == kube-system",
		Reason:    "Operator is deployed into the kube-system namespace.",
		Kinds:     []string{"Namespace", "Deployment", "ClusterRoleBinding"},
		Points:    -9,
	}
	list = append(list, kubesystemNamespaceRule)

	// OPR-R3-SC - No securityContext
	noSecurityContextRule := Rule{
		Predicate: rules.NoSecurityContext,
		ID:        "NoSecurityContext",
		Selector:  ".spec .template .spec .securityContext .containers[] ",
		Reason:    "Operators should be deployed with securityContextApplied",
		Kinds:     []string{"Pod", "Deployment", "StatefulSet", "DaemonSet"},
		Points:    -9,
	}
	list = append(list, noSecurityContextRule)

	// OPR-R4-SC - securityContext set to allowPrivilegeEscalation: true
	allowPrivilegeEscalation := Rule{
		Predicate: rules.AllowPrivilegeEscalation,
		ID:        "AllowPrivilegeEscalation",
		Selector:  ".spec .containers[] .securityContext .allowPrivilegeEscalation == true",
		Reason:    "Operators should not deploy with allowPrivilegeEscalation: true",
		Kinds:     []string{"Pod", "Deployment", "StatefulSet", "DaemonSet"},
		Points:    -7,
	}
	list = append(list, allowPrivilegeEscalation)

	// OPR-R5-SC - securityContext set to privileged: true
	privilegedRule := Rule{
		Predicate: rules.Privileged,
		ID:        "Privileged",
		Selector:  ".spec .containers[] .securityContext .privileged == true",
		Reason:    "Operators should not deploy with privileged: true",
		Kinds:     []string{"Pod", "Deployment", "StatefulSet", "DaemonSet"},
		Points:    -30,
	}
	list = append(list, privilegedRule)

	// OPR-R6-SC - securityContext set to readOnlyRootFilesystem: false
	readOnlyRootFilesystemRule := Rule{
		Predicate: rules.ReadOnlyRootFilesystem,
		ID:        "ReadOnlyRootFilesystem",
		Selector:  ".spec .containers[] .securityContext .readOnlyRootFilesystem == false",
		Reason:    "Operators should not deploy with readOnlyRootFilesystem: true",
		Kinds:     []string{"Pod", "Deployment", "StatefulSet", "DaemonSet"},
		Points:    -7,
	}
	list = append(list, readOnlyRootFilesystemRule)

	// OPR-R7-SC - securityContext set to runAsNonRoot: false
	runAsNonRootRule := Rule{
		Predicate: rules.RunAsNonRoot,
		ID:        "RunAsNonRoot",
		Selector:  ".spec .containers[] .securityContext .runAsNonRoot == false",
		Reason:    "Operators should not run as the root user",
		Kinds:     []string{"Pod", "Deployment", "StatefulSet", "DaemonSet"},
		Points:    -7,
	}
	list = append(list, runAsNonRootRule)

	// OPR-R8-SC - securityContext set to runAsUser: 0
	runAsUserRule := Rule{
		Predicate: rules.RunAsUser,
		ID:        "RunAsUser",
		Selector:  ".spec containers[] .securityContext .runAsUser -gt 0",
		Reason:    "Operators should not run as the root user (UID = 0)",
		Kinds:     []string{"Pod", "Deployment", "StatefulSet", "DaemonSet"},
		Points:    -7,
	}
	list = append(list, runAsUserRule)

	// OPR-R9-SC - securityContext adds CAP_SYS_ADMIN Linux capability
	capSysAdminRule := Rule{
		Predicate: rules.CapSysAdmin,
		ID:        "CapSysAdmin",
		Selector:  "containers[] .securityContext .capabilities .add == SYS_ADMIN",
		Reason:    "CAP_SYS_ADMIN is the most privileged capability and where possible disabled for Operators",
		Kinds:     []string{"Pod", "Deployment", "StatefulSet", "DaemonSet"},
		Points:    -30,
	}
	list = append(list, capSysAdminRule)

	// OPR-R10-RBAC - Runs as Cluster Admin
	clusterAdminRule := Rule{
		Predicate: rules.ClusterAdmin,
		ID:        "ClusterAdmin",
		Selector:  ".roleRef .name",
		Reason:    "The Operator is using Kubernetes native cluster admin role. Operators must use a dedicated cluster role",
		Kinds:     []string{"ClusterRoleBinding"},
		Points:    -30,
	}
	list = append(list, clusterAdminRule)

	// OPR-R11-RBAC - ClusterRole has full permissions over all resources
	starAllClusterRoleRule := Rule{
		Predicate: rules.StarAllClusterRole,
		ID:        "StarAllClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has full permissions on all resources in the cluster",
		Kinds:     []string{"ClusterRole"},
		Points:    -30,
	}
	list = append(list, starAllClusterRoleRule)

	// OPR-R12-RBAC - ClusterRole has full permissions over all CoreAPI resources
	starAllCoreAPIClusterRoleRule := Rule{
		Predicate: rules.StarAllCoreAPIClusterRole,
		ID:        "StarAllCoreAPIClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has full permissions on all CoreAPI resources in the cluster",
		Kinds:     []string{"ClusterRole"},
		Points:    -30,
	}
	list = append(list, starAllCoreAPIClusterRoleRule)

	// OPR-R13-RBAC - ClusterRole has full permissions over ClusterRoles and ClusterRoleBindings
	starClusterRoleAndBindingsRule := Rule{
		Predicate: rules.StarClusterRoleAndBindings,
		ID:        "StarClusterRoleAndBindings",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has full permissions over ClusterRoles and ClusterRoleBindings",
		Kinds:     []string{"ClusterRole"},
		Points:    -9,
	}
	list = append(list, starClusterRoleAndBindingsRule)

	// OPR-R14-RBAC - ClusterRole has access to Kubernetes secrets
	secretsClusterRoleRule := Rule{
		Predicate: rules.SecretsClusterRole,
		ID:        "SecretsClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has access to all secrets",
		Kinds:     []string{"ClusterRole"},
		Points:    -9,
	}
	list = append(list, secretsClusterRoleRule)

	// OPR-R15-RBAC - ClusterRole can exec into Pods
	execPodsClusterRoleRule := Rule{
		Predicate: rules.ExecPodsClusterRole,
		ID:        "ExecPodsClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has permissions to exec into any pod in the cluster",
		Kinds:     []string{"ClusterRole"},
		Points:    -9,
	}
	list = append(list, execPodsClusterRoleRule)

	// OPR-R16-RBAC - ClusterRole has escalate permissions
	escalateClusterRoleRule := Rule{
		Predicate: rules.EscalateClusterRole,
		ID:        "EscalateClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has escalate permissions",
		Kinds:     []string{"ClusterRole"},
		Points:    -9,
	}
	list = append(list, escalateClusterRoleRule)

	// OPR-R17-RBAC - ClusterRole has bind permissions
	bindClusterRoleRule := Rule{
		Predicate: rules.BindClusterRole,
		ID:        "BindClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has bind permissions",
		Kinds:     []string{"ClusterRole"},
		Points:    -9,
	}
	list = append(list, bindClusterRoleRule)

	// OPR-R18-RBAC - ClusterRole has impersonate permissions
	impersonateClusterRoleRule := Rule{
		Predicate: rules.ImpersonateClusterRole,
		ID:        "ImpersonateClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has impersonate permissions",
		Kinds:     []string{"ClusterRole"},
		Points:    -30,
	}
	list = append(list, impersonateClusterRoleRule)

	// OPR-R19-RBAC - ClusterRole can modify pod logs
	modifyPodLogsClusterRoleRule := Rule{
		Predicate: rules.ModifyPodLogsClusterRole,
		ID:        "ModifyPodLogsClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has permissions to modify pod logs",
		Kinds:     []string{"ClusterRole"},
		Points:    -1,
	}
	list = append(list, modifyPodLogsClusterRoleRule)

	// OPR-R20-RBAC - ClusterRole can remove Kubernetes events
	removeEventsClusterRoleRule := Rule{
		Predicate: rules.RemoveEventsClusterRole,
		ID:        "RemoveEventsClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has permissions to delete Kubernetes Events",
		Kinds:     []string{"ClusterRole"},
		Points:    -9,
	}
	list = append(list, removeEventsClusterRoleRule)

	// OPR-R21-RBAC - ClusterRole has full permissions over any custom resource definitions
	customResourceClusterRoleRule := Rule{
		Predicate: rules.CustomResourceClusterRole,
		ID:        "CustomResourceClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has permissions over any Custom Resource",
		Kinds:     []string{"ClusterRole"},
		Points:    -1,
	}
	list = append(list, customResourceClusterRoleRule)

	// OPR-R22-RBAC - ClusterRole has full permissions over admission controllers
	admissionControllerClusterRoleRule := Rule{
		Predicate: rules.AdmissionControllerClusterRole,
		ID:        "AdmissionControllerClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has full permissions over Admission Controllers",
		Kinds:     []string{"ClusterRole"},
		Points:    -9,
	}
	list = append(list, admissionControllerClusterRoleRule)

	// OPR-R23-RBAC - ClusterRole has permissions over service account token creation
	serviceAccountClusterRoleRule := Rule{
		Predicate: rules.ServiceAccountClusterRole,
		ID:        "ServiceAccountClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has permissions over service accounts to create token requests for existing service accounts",
		Kinds:     []string{"ClusterRole"},
		Points:    -9,
	}
	list = append(list, serviceAccountClusterRoleRule)

	// OPR-R24-RBAC - ClusterRole has read, write or delete permissions over persistent volumes
	persistentVolumeClusterRoleRule := Rule{
		Predicate: rules.PersistentVolumeClusterRole,
		ID:        "PersistentVolumeClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has read, write or delete permissions over persistent volumes",
		Kinds:     []string{"ClusterRole"},
		Points:    -9,
	}
	list = append(list, persistentVolumeClusterRoleRule)

	// OPR-R25-RBAC - ClusterRole has read, write or delete permissions over network policies
	networkPolicyClusterRoleRule := Rule{
		Predicate: rules.NetworkPolicyClusterRole,
		ID:        "NetworkPolicyClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has full permissions over network policies",
		Kinds:     []string{"ClusterRole"},
		Points:    -9,
	}
	list = append(list, networkPolicyClusterRoleRule)

	// OPR-R26-RBAC - ClusterRole has permissions over the Kubernetes API server proxy
	nodeProxyClusterRoleRule := Rule{
		Predicate: rules.NodeProxyClusterRole,
		ID:        "NodeProxyClusterRole",
		Selector:  ".rules .apiGroups .resources .verbs",
		Reason:    "The Operator SA cluster role has permissions the Kubernetes API server proxy",
		Kinds:     []string{"ClusterRole"},
		Points:    -9,
	}
	list = append(list, nodeProxyClusterRoleRule)

	return &Ruleset{
		Rules:  list,
		logger: logger,
	}
}

func (rs *Ruleset) Run(fileName string, fileBytes []byte, schemaDir string) ([]Report, error) {
	reports := make([]Report, 0)

	isJSON := json.Valid(fileBytes)
	if isJSON {
		report := rs.generateReport(fileName, fileBytes, schemaDir)
		reports = append(reports, report)
	} else {
		lineBreak := detectLineBreak(fileBytes)
		bits := bytes.Split(fileBytes, []byte(lineBreak+"---"+lineBreak))
		for i, d := range bits {
			doc := bytes.TrimSpace(d)

			// If empty or just a header
			if len(doc) == 0 || (len(doc) == 3 && string(doc) == "---") {
				// if we're at the end and there are no reports
				if len(bits) == i+1 && len(reports) == 0 {
					rs.logger.Debugf("empty and no records, erroring")
					return nil, &InvalidInputError{}
				}
				rs.logger.Debugf("empty but still more docs, continuing")
				continue
			}
			data, err := yaml.YAMLToJSON(doc)
			if err != nil {
				return reports, err
			}
			report := rs.generateReport(fileName, data, schemaDir)
			reports = append(reports, report)
		}
	}

	return reports, nil
}

// func GenerateInTotoLink(reports []Report, fileBytes []byte) in_toto.Metablock {

// 	var linkMb in_toto.Metablock

// 	materials := make(map[string]interface{})
// 	request := make(map[string]interface{})

// 	// INFO: it appears that the last newline of the yaml is removed when
// 	// receiving, which makes the integrity check fail on other implementations
// 	fileBytes = append(fileBytes, 10)

// 	request["sha256"] = fmt.Sprintf("%x", sha256.Sum256([]uint8(fileBytes)))

// 	// TODO: the filename should be a parameter passed to the report (as it is
// 	// very likely other filenames will exist in supply chains)
// 	materials["deployment.yml"] = request

// 	products := make(map[string]interface{})
// 	for _, report := range reports {
// 		reportArtifact := make(map[string]interface{})
// 		// FIXME: encoding as json now for integrity check, this is the wrong way
// 		// to compute the hash over the result. Also, some error checking would be
// 		// more than ideal.
// 		reportValue, _ := json.Marshal(report)
// 		reportArtifact["sha256"] =
// 			fmt.Sprintf("%x", sha256.Sum256([]uint8(reportValue)))
// 		products[report.Object] = reportArtifact
// 	}

// 	linkMb.Signatures = []in_toto.Signature{}
// 	linkMb.Signed = in_toto.Link{
// 		Type:       "link",
// 		Name:       "kubesec",
// 		Materials:  materials,
// 		Products:   products,
// 		ByProducts: map[string]interface{}{},
// 		// FIXME: the command should include whether this is called through the
// 		// server or a standalone tool.
// 		Command:     []string{},
// 		Environment: map[string]interface{}{},
// 	}

// 	return linkMb
// }

func appendUniqueRule(uniqueRules []RuleRef, newRule RuleRef) []RuleRef {
	if !containsRule(uniqueRules[:], newRule) {
		uniqueRules = append(uniqueRules, newRule)
	}
	return uniqueRules
}

func containsRule(rules []RuleRef, newRule RuleRef) bool {
	for _, rule := range rules {
		if reflect.DeepEqual(rule, newRule) {
			return true
		}
	}
	return false
}

func (rs *Ruleset) generateReport(fileName string, json []byte, schemaDir string) Report {
	report := Report{
		Object:   "Unknown",
		FileName: fileName,
		Score:    0,
		Rules:    make([]RuleRef, 0),
		Scoring: RuleScoring{
			Advise:   make([]RuleRef, 0),
			Passed:   make([]RuleRef, 0),
			Critical: make([]RuleRef, 0),
		},
	}

	report.Object = getObjectName(json)

	// KGW removed kubeval due to out of date schema validation breaking rule checks

	// validate resource with kubeval
	// cfg := kubeval.NewDefaultConfig()
	// cfg.FileName = fileName
	// cfg.Strict = true

	// if schemaDir != "" {
	// 	cfg.SchemaLocation = "file://" + schemaDir
	// } else if _, err := os.Stat("/schemas/kubernetes-json-schema/master/master-standalone"); !os.IsNotExist(err) {
	// 	cfg.SchemaLocation = "file:///schemas"
	// }

	// results, err := kubeval.Validate(json, cfg)
	// if err != nil {
	// 	if strings.Contains(err.Error(), "404 Not Found") {
	// 		report.Message = "This resource is invalid, unknown schema"
	// 	} else {
	// 		report.Message = err.Error()
	// 	}
	// 	return report
	// }

	// for _, result := range results {
	// 	if len(result.Errors) > 0 {
	// 		for _, desc := range result.Errors {
	// 			report.Message += desc.String() + " "
	// 		}
	// 	} else if result.Kind == "" {
	// 		report.Message += "This resource is invalid, Kubernetes kind not found"
	// 	}
	// }

	// if len(report.Message) > 0 {
	// 	return report
	// }
	report.Valid = true

	// run rules in parallel
	ch := make(chan RuleRef, len(rs.Rules))
	var wg sync.WaitGroup
	for _, rule := range rs.Rules {
		wg.Add(1)
		go eval(json, rule, ch, &wg)
	}
	wg.Wait()
	close(ch)

	// collect results
	var appliedRules int
	for ruleRef := range ch {
		appliedRules++

		report.Rules = appendUniqueRule(report.Rules, ruleRef)

		if ruleRef.Containers > 0 {
			if ruleRef.Points >= 0 {
				rs.logger.Debugf("positive score rule matched %v (%v points)", ruleRef.Selector, ruleRef.Points)
				report.Score += ruleRef.Points
				report.Scoring.Passed = append(report.Scoring.Passed, ruleRef)
			}

			if ruleRef.Points < 0 {
				rs.logger.Debugf("negative score rule matched %v (%v points)", ruleRef.Selector, ruleRef.Points)
				report.Score += ruleRef.Points
				report.Scoring.Critical = append(report.Scoring.Critical, ruleRef)
			}
		} else if ruleRef.Points >= 0 {
			rs.logger.Debugf("positive score rule failed %v (%v points)", ruleRef.Selector, ruleRef.Points)
			report.Scoring.Advise = append(report.Scoring.Advise, ruleRef)
		}
	}

	if appliedRules < 1 {
		report.Message = "This resource kind is not supported by badrobot"
	} else if report.Score >= 0 {
		report.Message = fmt.Sprintf("Passed with a score of %v points", report.Score)
	} else {
		report.Message = fmt.Sprintf("Failed with a score of %v points", report.Score)
	}

	// sort results into priority order
	sort.Sort(RuleRefCustomOrder(report.Scoring.Critical))
	sort.Sort(RuleRefCustomOrder(report.Scoring.Passed))
	sort.Sort(RuleRefCustomOrder(report.Scoring.Advise))

	return report
}

func eval(json []byte, rule Rule, ch chan RuleRef, wg *sync.WaitGroup) {
	defer wg.Done()

	containers, err := rule.Eval(json)

	// skip rule if it doesn't apply to object kind
	switch err.(type) {
	case *NotSupportedError:
		return
	}

	result := RuleRef{
		Containers: containers,
		ID:         rule.ID,
		Points:     rule.Points,
		Reason:     rule.Reason,
		Selector:   rule.Selector,
		Weight:     rule.Weight,
		Link:       rule.Link,
	}

	ch <- result
}

// getObjectName returns <kind>/<name>.<namespace>
func getObjectName(json []byte) string {
	jq := gojsonq.New().Reader(bytes.NewReader(json))
	if len(jq.Errors()) > 0 {
		return "Unknown"
	}

	kind := jq.Copy().From("kind").Get()
	if kind == nil {
		return "Unknown"
	}
	object := fmt.Sprintf("%v", kind)

	name := jq.Copy().From("metadata.name").Get()
	if name == nil {
		object += "/undefined"
	} else {
		object += fmt.Sprintf("/%v", name)
	}

	namespace := jq.Copy().From("metadata.namespace").Get()
	if namespace == nil {
		object += ".default"
	} else {
		object += fmt.Sprintf(".%v", namespace)
	}

	return object
}

func detectLineBreak(haystack []byte) string {
	windowsLineEnding := bytes.Contains(haystack, []byte("\r\n"))
	if windowsLineEnding && runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}
