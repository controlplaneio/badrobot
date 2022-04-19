# BadRobot (Kubernetes Operator Audit Tool)

- [Overview](#overview)
  - [Structure](#structure)
  - [Rulesets](#rulesets)
- [Command line Usage](#command-line-usage)

---
## Overview
Badrobot is the internal R&D project codename for a Kubernetes Operator Audit Tool. The purpose of the project is to create an open source tool which audits Public or Private Kubernetes Operators. As Operators can have a large scope the initial version will be focused on specific resources and Operators which conform to the Operator SDK. If successful, the tool could be expanded to identify any resources which are associated with an Operator.

### Structure
The tool is supposed be a static code analyser which can run against a code repository and provide an "opinionated" score on the current Operator configuration. This is similar to KubeSec and will be structured in the same way, that is:
* code to scan a .yaml file
* a package of rules to apply against the .yaml file

### Rulesets

| RuleSet ID | Rule | Risk | Risk Level |
|-----|-----|-----|-----|
| OPR-R1-NS | default Namespace | The Operator is deployed onto the default namespace. Operators should be deployed into a dedicated namespace to reduce the exposure of other sensitive in the event of compromise. The default namespace is where applications and services are deployed if the namespaces is not specified. A compromised application container could pivot to the Operator which may allow an adversary to obtain cluster wide permissions. | Low |
| OPR-R2-NS | kube-system Namespace | The Operator is deployed onto the kube-system namespace. Operators should be deployed into a dedicated namespace to reduce the exposure of other sensitive in the event of compromise. The kube-system namespace is reserved for Kubernetes engine and the Operator should not be deployed here. | High |
| OPR-R3-SC | No securityContext | The Operator is deployed without a securityContext. In the event the Operator is compromised, the adversary would have unrestricted access to resources on the underlying host. Unless the Operator is performing highly permissive cluster configuration and management of resources, it is highly recommended to some restrictions are applied. | High |
| OPR-R4-SC | securityContext set to allowPrivilegeEscalation: true | The Operator is deployed with privilege escalation permissions which in the event of a compromise would allow adversary root access to the underlying host. By default, the Operator-SDK sets allowPrivilegeEscalation: false and must be explictly removed or modified in the deployment manifest. | High |
| OPR-R5-SC | securityContext set to privileged: true | The Operator is deployed with all of the system root’s capabilities. In the event the Operator is compromised, the Adversary would have unrestricted access to resources on the underlying host.  | Medium |
| OPR-R6-SC | securityContext set to readOnlyRootFilesystem: false | The Operator is deployed with write access to the underlying host. In the event the Operator is compromised and the Operator has mount access, an adversary would be able to write to root filesystem to obtain full system compromise. | Medium |
| OPR-R7-SC | securityContext set to runAsNonRoot: false | The Operator is configured to run as root. If the Operator has access to the underlying host it will have the same access as host root account. By default, the Operator-SDK sets runAsNonRoot: true and must be explictly removed or modified in the deployment manifest. | Medium |
| OPR-R8-SC | securityContext set to runAsUser: 0 | The Operator is configured to run as root. If the Operator has access to the underlying host it will have the same access as host root account. By default, the Operator-SDK sets runAsNonRoot: true and must be explictly removed or modified in the deployment manifest. | Medium |
| OPR-R9-SC | securityContext adds CAP_SYS_ADMIN Linux capability | The Operator is configured with CAP_SYS_ADMIN enabled, removing any previously dropped Linux capabilities. CAP_SYS_ADMIN is an overloaded capability allowing system adminstrative operations and can lead to privilege escalation on the host if the Operator is compromised. | High |
| OPR-R10-RBAC | Runs as Cluster Admin | The Operator runs as default cluster role, cluster admin. Even if the Operator requires full cluster administration, this role should not be used and dedicated one instead. It is recommended that the permissions of the Operator are reviewed and redefined. | Critical (failure) |
| OPR-R11-RBAC | ClusterRole has full permissions over all resources | The Operator runs with a cluster role with full access (\*) to all resources (\*). Even if the Operator requires full cluster administration, the cluster role should explicitly define apigroups, resources and verbs it requires access to. It is recommended that the permissions of the Operator are reviewed and redefined. | Critical (failure) |
| OPR-R12-RBAC | ClusterRole has full permissions over all CoreAPI resources  | The Operator runs with a cluster role with full access (\*) to CoreAPI resources (\*). Even if the Operator requires full access to all CoreAPI resources, the cluster role should explicitly define resources and verbs it requires. It is recommended that the permissions of the Operator are reviewed and redefined. | Critical (failure) |
| OPR-R13-RBAC | ClusterRole has full permissions over ClusterRoles and ClusterRoleBindings | The Operator runs with a cluster role which has unrestricted access to ClusterRoles and ClusterRoleBindings. In the event of a compromise, an adversary would be able to rebind the Operator service account with full access to cluster wide resources. | High |
| OPR-R14-RBAC | ClusterRole has access to Kubernetes secrets | The Operator is deployed with access to all secrets across the cluster. Accessing this level of cluster wide secrets may allow an adversary a method of privilege escalation. It is highly recommended that a dedicated role is used for accessing specific secrets the Operator requires to manage and use. | High |
| OPR-R15-RBAC | ClusterRole can exec into Pods | The Operator is deployed with a cluster role that allows remote access to any Pod in the cluster. In the event the Operator is compromised, the adversary would be able to pivot to different containers to attempt escape isolation or access sensitive information. | Medium |
| OPR-R16-RBAC | ClusterRole has escalate permissions | The Operator is deployed with the escalation privilege, allowing the cluster role grant privileges beyond the permissions that are bound to the Operator. Due to this reason, it is recommended that Operators are not given the escalate privilege. | High |
| OPR-R17-RBAC | ClusterRole has bind permissions| The Operator is deployed with the bind privilege, allowing the cluster role to create bindings of roles beyond the permissions granted to the Operator. Due to this reason, it is recommended that Operators are not given the bind privilege.  | High |
| OPR-R18-RBAC | ClusterRole has impersonate permissions | The Operator is deployed with the impersonate privilege, allowing the cluster role to gain the rights of another role. This permission would allow an adversary (with access to the Operator) to masquerade as another as another role to perform malicious actions on cluster resources. Due to this, it is recommended that Operators are not given the impersonate privilege. | Critical |
| OPR-R19-RBAC | ClusterRole can modify pod logs | The Operator is deployed with full permissions over pod logs. The permission can be abused by adversary to remove or overwrite pod logs masking malicious actions performed against pod resources. | Low |
| OPR-R20-RBAC | ClusterRole can remove Kubernetes events | The Operator is deployed with access to deleting Kubernetes events. In the event the Operator is compromised, an Adversary would be able to remove Kubernetes events and hide previous malicious actions. | High |
| OPR-R21-RBAC | ClusterRole has full permissions over any custom resource definitions | The Operator is deployed with full permissions over any custom resource defintion, allowing the modification of custom resources managed by another Operator. The cluster role or role must only be scoped to custom resources managed by the Operator. In the event the Operator is compromised or a malicious Operator is deployed, it would allow an adversary to manipulate custom resources. | High |
| OPR-R22-RBAC | ClusterRole has full permissions over admission controllers | The Operator is deployed with full permissions over Kubernetes admission controllers which can allow an adversary to read or modified submitted resources. | High |
| OPR-R23-RBAC | ClusterRole has permissions over service account token creation | The Operator has full permissions over cluster wide service accounts, allowing the creation of token requests for existing service accounts. This can be abused by an adversary to masquerade as another user to access cluster resources. | High |
| OPR-R24-RBAC | ClusterRole has read, write or delete permissions over persistent volumes | The Operator is deployed with read, write or delete permissions for volume mount, allowing root filesystem access and exposuring sensitive information. | High |
| OPR-R25-RBAC | ClusterRole has read, write or delete permissions over network policies | The Operator is deployed with access to cluster wide network policies, allowing the modification of network routes. An adversary can leverage these permissions to access unauthorised resources. | Low |
| OPR-R26-RBAC | ClusterRole has permissions over the Kubernetes API server proxy | The Operator is deployed with permissions over the proxy sub resource of the node, allowing command execution on every pod on the node via the Kubelet API. An adversary can leverage this permission on the Operator to run custom workloads on several pods on the node. | High |

---
## Command Line Usage
```bash
$ badrobot scan operator-manifest.yaml
```
> Currently there are example operator files to use for Badrobot under test/example-yaml
## Appendix
