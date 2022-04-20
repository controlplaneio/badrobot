#!/usr/bin/env bash

load '_helper'

setup() {
  _global_setup
}

teardown() {
  _global_teardown
}

# ------------------------------------#
# NAMESPACE RULESET TESTS             #
# ------------------------------------#

# OPR-R1-NS
@test "fails default namespace defined (kind: Namespace)" {
  run _app "${TEST_DIR}/asset/ns-default.yaml"
  assert_lt_zero_points
}

# OPR-R1-NS
@test "fails default namespace defined (kind: Deployment)" {
  run _app "${TEST_DIR}/asset/deploy-ns-default.yaml"
  assert_lt_zero_points
}

# OPR-R1-NS
@test "fails default namespace defined (kind: ClusterRoleBinding)" {
  run _app "${TEST_DIR}/asset/crb-ns-default.yaml"
  assert_lt_zero_points
}

# Dedicated NS
@test "passes dedicated namespace defined (kind: Namespace)" {
  run _app "${TEST_DIR}/asset/ns-dedicated.yaml"
  assert_zero_points
}

# Dedicated NS
@test "passes dedicated namespace defined (kind: Deployment)" {
  run _app "${TEST_DIR}/asset/deploy-ns-dedicated.yaml"
  assert_zero_points
}

# Dedicated NS
@test "passes dedicated namespace defined (kind: ClusterRoleBinding)" {
  run _app "${TEST_DIR}/asset/crb-ns-dedicated.yaml"
  assert_zero_points
}

# OPR-R2-NS
@test "fails kube-system namespace defined (kind: Namespace)" {
  run _app "${TEST_DIR}/asset/ns-kubesystem.yaml"
  assert_lt_zero_points
}

# OPR-R2-NS
@test "fails kube-system namespace defined (kind: Deployment)" {
  run _app "${TEST_DIR}/asset/deploy-ns-kubesystem.yaml"
  assert_lt_zero_points
}

# OPR-R2-NS
@test "fails kube-system namespace defined (kind: ClusterRoleBinding)" {
  run _app "${TEST_DIR}/asset/crb-ns-kubesystem.yaml"
  assert_lt_zero_points
}

# ------------------------------------#
# SECURITYCONTEXT RULESET TESTS       #
# ------------------------------------#

# All securityContexts
@test "passes all security contexts defined" {
  run _app "${TEST_DIR}/asset/deploy-sc-both-all.yaml"
  assert_zero_points
}

# All securityContexts under containers
@test "passes all security contexts defined under containers" {
  run _app "${TEST_DIR}/asset/deploy-sc-containers-all.yaml"
  assert_zero_points
}

# All securityContexts under spec
@test "passes all security contexts defined under spec" {
  run _app "${TEST_DIR}/asset/deploy-sc-spec-all.yaml"
  assert_zero_points
}

# OPR-R3-SC
@test "fails no security contexts defined" {
  run _app "${TEST_DIR}/asset/deploy-sc-none.yaml"
  assert_lt_zero_points
}

# OPR-R4-SC
@test "fails allowPrivilegeEscalation set to true (containers)" {
  run _app "${TEST_DIR}/asset/deploy-sc-containers-allow-priv.yaml"
  assert_lt_zero_points
}

# OPR-R4-SC
@test "fails allowPrivilegeEscalation set to true (spec)" {
  run _app "${TEST_DIR}/asset/deploy-sc-spec-allow-priv.yaml"
  assert_lt_zero_points
}

# OPR-R5-SC
@test "fails privileged set to true (containers)" {
  run _app "${TEST_DIR}/asset/deploy-sc-containers-priv.yaml"
  assert_lt_zero_points
}

# OPR-R5-SC
@test "fails privileged set to true (spec)" {
  run _app "${TEST_DIR}/asset/deploy-sc-spec-priv.yaml"
  assert_lt_zero_points
}

# OPR-R6-SC
@test "fails readOnlyRootFilesystem set to false (containers)" {
  run _app "${TEST_DIR}/asset/deploy-sc-containers-readonly-root.yaml"
  assert_lt_zero_points
}

# OPR-R6-SC
@test "fails readOnlyRootFilesystem set to false (spec)" {
  run _app "${TEST_DIR}/asset/deploy-sc-spec-readonly-root.yaml"
  assert_lt_zero_points
}

# OPR-R7-SC
@test "fails runAsNonRoot set to false (containers)" {
  run _app "${TEST_DIR}/asset/deploy-sc-containers-nonroot.yaml"
  assert_lt_zero_points
}

# OPR-R7-SC
@test "fails runAsNonRoot set to false (spec)" {
  run _app "${TEST_DIR}/asset/deploy-sc-spec-nonroot.yaml"
  assert_lt_zero_points
}

# OPR-R8-SC
@test "fails runAsUser set to 0 (containers)" {
  run _app "${TEST_DIR}/asset/deploy-sc-containers-rootuser.yaml"
  assert_lt_zero_points
}

# OPR-R8-SC
@test "fails runAsUser set to 0 (spec)" {
  run _app "${TEST_DIR}/asset/deploy-sc-spec-rootuser.yaml"
  assert_lt_zero_points
}

# OPR-R9-SC
@test "fails CAP_SYS_ADMIN added (containers)" {
  run _app "${TEST_DIR}/asset/deploy-sc-containers-cap-sysadmin.yaml"
  assert_lt_zero_points
}

# OPR-R9-SC
@test "fails CAP_SYS_ADMIN added (spec)" {
  run _app "${TEST_DIR}/asset/deploy-sc-spec-cap-sysadmin.yaml"
  assert_lt_zero_points
}

# ------------------------------------#
# CLUSTERROLEBINDING RULESET TESTS    #
# ------------------------------------#

# Runs as dedicated Cluster Role
@test "passes dedicated cluster role defined" {
  run _app "${TEST_DIR}/asset/crb-dedicated-cr.yaml"
  assert_zero_points
}

# Runs as dedicated Cluster Role (cluster-admin-suffix)
@test "passes dedicated cluster role defined (suffix)" {
  run _app "${TEST_DIR}/asset/crb-cluster-admin-suffix.yaml"
  assert_zero_points
}

# Runs as dedicated Cluster Role (prefix-cluster-admin)
@test "passes dedicated cluster role defined (prefix)" {
  run _app "${TEST_DIR}/asset/crb-prefix-cluster-admin.yaml"
  assert_zero_points
}

# OPR-R10-RBAC
@test "fails cluster-admin defined" {
  run _app "${TEST_DIR}/asset/crb-cluster-admin.yaml"
  assert_lt_zero_points
}

# ------------------------------------#
# CLUSTERROLE RULESET TESTS           #
# ------------------------------------#

# OPR-R11-RBAC
@test "fails ClusterRole has full access to all resources defined (*)" {
  run _app "${TEST_DIR}/asset/cr-all-star.yaml"
  assert_lt_zero_points
}

# OPR-R11-RBAC
@test "fails ClusterRole has full access to all resources defined (all verbs)" {
  run _app "${TEST_DIR}/asset/cr-all-verbs.yaml"
  assert_lt_zero_points
}

# CoreAPI Limited Resources
@test "passes ClusterRole has access to CoreAPI with limited resources defined" {
  run _app "${TEST_DIR}/asset/cr-coreapi-limited-resources.yaml"
  assert_zero_points
}

# CoreAPI Limited verbs
@test "passes ClusterRole has access to CoreAPI with limited verbs defined" {
  run _app "${TEST_DIR}/asset/cr-coreapi-limited-verbs.yaml"
  assert_zero_points
}

# Non-CoreAPI Limited Resources
@test "passes ClusterRole has access to Non-CoreAPI with limited resources defined" {
  run _app "${TEST_DIR}/asset/cr-noncoreapi-limited.yaml"
  assert_zero_points
}

# Non-CoreAPI All Resources
@test "passes ClusterRole has access to Non-CoreAPI with all resources defined" {
  run _app "${TEST_DIR}/asset/cr-noncoreapi-star.yaml"
  assert_zero_points
}

# OPR-R12-RBAC
@test "fails ClusterRole has full access to CoreAPI defined (*)" {
  run _app "${TEST_DIR}/asset/cr-coreapi-star.yaml"
  assert_lt_zero_points
}

# OPR-R12-RBAC
@test "fails ClusterRole has full access to CoreAPI defined (all verbs)" {
  run _app "${TEST_DIR}/asset/cr-coreapi-all-verbs.yaml"
  assert_lt_zero_points
}

# Only full access to ClusterRoles
@test "passes ClusterRole only has full access to ClusterRoles" {
  run _app "${TEST_DIR}/asset/cr-all-clusterroles-only.yaml"
  assert_zero_points
}

# Only full access to ClusterRoleBindings
@test "passes ClusterRole only has full access to ClusterRoleBindings" {
  run _app "${TEST_DIR}/asset/cr-all-clusterrolebindings-only.yaml"
  assert_zero_points
}

# OPR-R13-RBAC
@test "fails ClusterRole has full access to ClusterRoles and Bindings (*)" {
  run _app "${TEST_DIR}/asset/cr-all-clusterrolesandbindings-star.yaml"
  assert_lt_zero_points
}

# OPR-R13-RBAC
@test "fails ClusterRole has full access to ClusterRoles and Bindings (all verbs)" {
  run _app "${TEST_DIR}/asset/cr-all-clusterrolesandbindings-verbs.yaml"
  assert_lt_zero_points
}

# TBD KGW Split resources test. Requires additional logic in the ruleset
# # OPR-R13-RBAC
# @test "fails ClusterRole has full access to ClusterRoles and Bindings (separate)" {
#   run _app "${TEST_DIR}/asset/cr-all-crbs-separate.yaml"
#   assert_lt_zero_points
# }

# OPR-R14-RBAC
@test "fails ClusterRole has access to Kubernetes secrets (star)" {
  run _app "${TEST_DIR}/asset/cr-secrets-star.yaml"
  assert_lt_zero_points
}

# OPR-R14-RBAC
@test "fails ClusterRole has access to Kubernetes secrets (all verbs)" {
  run _app "${TEST_DIR}/asset/cr-secrets-all-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R15-RBAC
@test "fails ClusterRole has full access to pods (star)" {
  run _app "${TEST_DIR}/asset/cr-pods-star.yaml"
  assert_lt_zero_points
}

# OPR-R15-RBAC
@test "fails ClusterRole has get and create permissions on pods (verbs)" {
  run _app "${TEST_DIR}/asset/cr-pods-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R15-RBAC
@test "fails ClusterRole has full access to pods/exec (star)" {
  run _app "${TEST_DIR}/asset/cr-podsexec-star.yaml"
  assert_lt_zero_points
}

# OPR-R15-RBAC
@test "fails ClusterRole has get and create permissions on pods/exec (verbs)" {
  run _app "${TEST_DIR}/asset/cr-podsexec-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R16-RBAC
@test "fails ClusterRole has escalate permissions" {
  run _app "${TEST_DIR}/asset/cr-escalate.yaml"
  assert_lt_zero_points
}

# OPR-R17-RBAC
@test "fails ClusterRole has bind permissions" {
  run _app "${TEST_DIR}/asset/cr-bind.yaml"
  assert_lt_zero_points
}

# OPR-R18-RBAC
@test "fails ClusterRole has impersonate permissions" {
  run _app "${TEST_DIR}/asset/cr-impersonate.yaml"
  assert_lt_zero_points
}

# OPR-R19-RBAC
@test "fails ClusterRole has full permissions to modify pod logs (star)" {
  run _app "${TEST_DIR}/asset/cr-podslog-star.yaml"
  assert_lt_zero_points
}

# OPR-R19-RBAC
@test "fails ClusterRole has full permissions to modify pod logs (verbs)" {
  run _app "${TEST_DIR}/asset/cr-podslog-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R20-RBAC
@test "fails ClusterRole has permissions to delete Kubernetes events (star)" {
  run _app "${TEST_DIR}/asset/cr-remove-events-star.yaml"
  assert_lt_zero_points
}

# OPR-R20-RBAC
@test "fails ClusterRole has permissions to delete Kubernetes events (verbs)" {
  run _app "${TEST_DIR}/asset/cr-remove-events-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R21-RBAC
@test "fails ClusterRole has full permissions over ANY crd (star)" {
  run _app "${TEST_DIR}/asset/cr-custom-resource-star.yaml"
  assert_lt_zero_points
}

# OPR-R21-RBAC
@test "fails ClusterRole has full permissions over ANY crd (verbs)" {
  run _app "${TEST_DIR}/asset/cr-custom-resource-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R22-RBAC
@test "fails ClusterRole has full permissions over mutating admission controllers (star)" {
  run _app "${TEST_DIR}/asset/cr-ad-mutating-webhook-star.yaml"
  assert_lt_zero_points
}

# OPR-R22-RBAC
@test "fails ClusterRole has full permissions over mutating admission controllers (verbs)" {
  run _app "${TEST_DIR}/asset/cr-ad-mutating-webhook-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R22-RBAC
@test "fails ClusterRole has full permissions over validating admission controllers (star)" {
  run _app "${TEST_DIR}/asset/cr-ad-validating-webhook-star.yaml"
  assert_lt_zero_points
}

# OPR-R22-RBAC
@test "fails ClusterRole has full permissions over validating admission controllers (verbs)" {
  run _app "${TEST_DIR}/asset/cr-ad-validating-webhook-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R22-RBAC
@test "fails ClusterRole has full permissions over both admission controllers (star)" {
  run _app "${TEST_DIR}/asset/cr-ad-both-star.yaml"
  assert_lt_zero_points
}

# OPR-R22-RBAC
@test "fails ClusterRole has full permissions over both admission controllers (verbs)" {
  run _app "${TEST_DIR}/asset/cr-ad-both-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R23-RBAC
@test "pass ClusterRole has get permissions over service accounts" {
  run _app "${TEST_DIR}/asset/cr-sa-only-get.yaml"
  assert_zero_points
}

# OPR-R23-RBAC
@test "pass ClusterRole has get permissions over service accounts" {
  run _app "${TEST_DIR}/asset/cr-sa-token-only-get.yaml"
  assert_zero_points
}

# OPR-R23-RBAC
@test "fails ClusterRole has full permissions over service accounts (star)" {
  run _app "${TEST_DIR}/asset/cr-sa-star.yaml"
  assert_lt_zero_points
}

# OPR-R23-RBAC
@test "fails ClusterRole has full permissions over service accounts (verbs)" {
  run _app "${TEST_DIR}/asset/cr-sa-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R23-RBAC
@test "fails ClusterRole has permissions to access service account tokens (star)" {
  run _app "${TEST_DIR}/asset/cr-sa-token-star.yaml"
  assert_lt_zero_points
}

# OPR-R23-RBAC
@test "fails ClusterRole has permissions to access service account tokens (verbs)" {
  run _app "${TEST_DIR}/asset/cr-sa-token-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R24-RBAC
@test "fails ClusterRole has full permissions over persistent volume claims (star)" {
  run _app "${TEST_DIR}/asset/cr-pvc-star.yaml"
  assert_lt_zero_points
}

# OPR-R24-RBAC
@test "fails ClusterRole has read, write or delete permissions over persistent volume claim (verbs)" {
  run _app "${TEST_DIR}/asset/cr-pvc-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R25-RBAC
@test "fails ClusterRole has full permissions over network (star)" {
  run _app "${TEST_DIR}/asset/cr-network-star.yaml"
  assert_lt_zero_points
}

# OPR-R25-RBAC
@test "fails ClusterRole has read, write or delete permissions over network (verbs)" {
  run _app "${TEST_DIR}/asset/cr-network-verbs.yaml"
  assert_lt_zero_points
}

# OPR-R25-RBAC
@test "fails ClusterRole has full permissions over network policies (star)" {
  run _app "${TEST_DIR}/asset/cr-network-policy-star.yaml"
  assert_lt_zero_points
}

# OPR-R25-RBAC
@test "fails ClusterRole has read, write or delete permissions over network policies (verbs)" {
  run _app "${TEST_DIR}/asset/cr-network-policy-verbs.yaml"
  assert_lt_zero_points
}

# @test "fails Pod with unconfined seccomp" {
#   run _app "${TEST_DIR}/asset/score-0-pod-seccomp-unconfined.yml"
#   assert_lt_zero_points
# }

# @test "fails with CAP_SYS_ADMIN" {
#   run _app "${TEST_DIR}/asset/score-0-cap-sys-admin.yml"
#   assert_lt_zero_points
# }

# @test "fails with CAP_CHOWN" {
#   run _app "${TEST_DIR}/asset/score-0-cap-chown.yml"
#   assert_zero_points
# }

# @test "fails with CAP_SYS_ADMIN and CAP_CHOWN" {
#   run _app "${TEST_DIR}/asset/score-0-cap-sys-admin-and-cap-chown.yml"
#   assert_lt_zero_points
# }

# @test "passes with securityContext capabilities drop all" {
#   run _app "${TEST_DIR}/asset/score-1-cap-drop-all.yml"
#   assert_gt_zero_points
# }

# @test "passes deployment with securitycontext readOnlyRootFilesystem" {
#   run _app "${TEST_DIR}/asset/score-1-dep-ro-root-fs.yml"
#   assert_gt_zero_points
# }

# @test "passes deployment with securitycontext runAsNonRoot" {
#   run _app "${TEST_DIR}/asset/score-1-dep-seccon-run-as-non-root.yml"
#   assert_gt_zero_points
# }

# @test "fails deployment with securitycontext runAsUser 1" {
#   run _app "${TEST_DIR}/asset/score-1-dep-seccon-run-as-user-1.yml"
#   assert_zero_points
# }

# @test "passes deployment with securitycontext runAsUser > 10000" {
#   run _app "${TEST_DIR}/asset/score-1-dep-seccon-run-as-user-10001.yml"
#   assert_gt_zero_points
# }

# @test "fails deployment with empty security context" {
#   run _app "${TEST_DIR}/asset/score-1-dep-empty-security-context.yml"
#   assert_zero_points
# }

# @test "fails deployment with invalid security context" {
#   run _app "${TEST_DIR}/asset/score-1-dep-invalid-security-context.yml"

#   run jq -r .[].message <<<"${output}"

#   assert_output --partial 'Additional property fake is not allowed'
# }

# @test "passes deployment with cgroup resource limits" {
#   run _app "${TEST_DIR}/asset/score-1-dep-resource-limit-cpu.yml"
#   assert_gt_zero_points
# }

# @test "passes deployment with cgroup memory limits" {
#   run _app "${TEST_DIR}/asset/score-1-dep-resource-limit-memory.yml"
#   assert_gt_zero_points
# }

# @test "passes StatefulSet with volumeClaimTemplate" {
#   run _app "${TEST_DIR}/asset/score-1-statefulset-volumeclaimtemplate.yml"
#   assert_gt_zero_points
# }

# @test "fails StatefulSet with no security" {
#   run _app "${TEST_DIR}/asset/score-0-statefulset-no-sec.yml"
#   assert_zero_points
# }

# @test "fails DaemonSet with securityContext.privileged = true" {
#   run _app "${TEST_DIR}/asset/score-0-daemonset-securitycontext-privileged.yml"
#   assert_lt_zero_points
# }

# @test "fails DaemonSet with mounted host docker.sock" {
#   run _app "${TEST_DIR}/asset/score-0-daemonset-mount-docker-socket.yml"
#   assert_lt_zero_points
# }

# @test "passes Pod with apparmor annotation" {
#   run _app "${TEST_DIR}/asset/score-3-pod-apparmor.yaml"
#   assert_gt_zero_points
# }

# @test "fails Pod with unconfined seccomp for all containers" {
#   run _app "${TEST_DIR}/asset/score-0-pod-seccomp-unconfined.yml"
#   assert_lt_zero_points
# }

# @test "passes Pod with non-unconfined seccomp for all containers" {
#   run _app "${TEST_DIR}/asset/score-0-pod-seccomp-non-unconfined.yml"
#   assert_gt_zero_points
# }

# @test "fails DaemonSet with hostNetwork" {
#   run _app "${TEST_DIR}/asset/score-0-daemonset-host-network.yml"
#   assert_lt_zero_points
# }

# @test "fails DaemonSet with hostPid" {
#   run _app "${TEST_DIR}/asset/score-0-daemonset-host-pid.yml"
#   assert_lt_zero_points
# }

# @test "fails DaemonSet with host docker.socket" {
#   run _app "${TEST_DIR}/asset/score-0-daemonset-volume-host-docker-socket.yml"
#   assert_lt_zero_points
# }

# @test "passes Deployment with serviceaccountname" {
#   run _app "${TEST_DIR}/asset/score-2-dep-serviceaccount.yml"
#   assert_gt_zero_points
# }

# @test "passes pod with serviceaccountname" {
#   run _app "${TEST_DIR}/asset/score-2-pod-serviceaccount.yml"
#   assert_gt_zero_points
# }

# @test "fails deployment with allowPrivilegeEscalation" {
#   run _app "${TEST_DIR}/asset/allowPrivilegeEscalation.yaml"
#   assert_lt_zero_points
# }

# @test "returns integer point score for each advice element" {
#   run _app "${TEST_DIR}/asset/score-2-pod-serviceaccount.yml"
#   assert_success

#   run jq -r .[].scoring.advise[].points <<<"${output}"

#   for SCORE in ${output}; do
#     assert bash -c "[[ ${SCORE} =~ ^[0-9]+$ ]]"
#   done
# }

# @test "returns an ordered point score for all advice" {
#   run _app "${TEST_DIR}/asset/score-2-pod-serviceaccount.yml"
#   assert_success

#   run jq -r .[].scoring.advise[].points <<<"${output}"

#   PREVIOUS=""
#   for CURRENT in ${output}; do
#     [ "${PREVIOUS}" = "" ] || assert [ "${CURRENT}" -le "${PREVIOUS}" ]
#     PREVIOUS="${CURRENT}"
#   done
# }

# @test "returns integer point score for each pass element" {
#   run _app "${TEST_DIR}/asset/score-5-pod-serviceaccount.yml"
#   assert_success

#   run jq -r .[].scoring.passed[].points <<<"${output}"

#   for SCORE in ${output}; do
#     assert bash -c "[[ ${SCORE} =~ ^[0-9]+$ ]]"
#   done
# }

# @test "returns an ordered point score for all passed" {
#   run _app "${TEST_DIR}/asset/score-5-pod-serviceaccount.yml"

#   run jq -r .[].scoring.passed[].points <<<"${output}"

#   PREVIOUS=""
#   for CURRENT in ${output}; do
#     [ "${PREVIOUS}" = "" ] || assert [ "${CURRENT}" -le "${PREVIOUS}" ]
#     PREVIOUS="${CURRENT}"
#   done
# }

# @test "check critical and advisory points listed by magnitude" {
#   run _app "${TEST_DIR}/asset/critical-double.yml"

#   # criticals - magnitude sort/lowest number first
#   CRITICAL_FIRST=$(jq -r .[].scoring.critical[0].points <<<"${output}")
#   CRITICAL_SECOND=$(jq -r .[].scoring.critical[1].points <<<"${output}")
#   (( CRITICAL_FIRST <= CRITICAL_SECOND ))

#   # advisories - magnitude sort/highest number first
#   ADVISE_FIRST=$(jq -r .[].scoring.advise[0].points <<<"${output}")
#   ADVISE_SECOND=$(jq -r .[].scoring.advise[1].points <<<"${output}")
#   ADVISE_THIRD=$(jq -r .[].scoring.advise[2].points <<<"${output}")
#   (( ADVISE_FIRST >= ADVISE_SECOND >= ADVISE_THIRD ))
# }

# @test "check critical and advisory points as multi-yaml" {
#   run _app "${TEST_DIR}/asset/critical-double-multiple.yml"

#   # report 1 - criticals - magnitude sort/lowest number first
#   CRITICAL_FIRST_FIRST=$(jq -r .[0].scoring.critical[0].points <<<"${output}")
#   CRITICAL_FIRST_SECOND=$(jq -r .[0].scoring.critical[1].points <<<"${output}")
#   (( CRITICAL_FIRST_FIRST <= CRITICAL_FIRST_SECOND ))

#   # report 1 - advisories - magnitude sort/highest number first
#   ADVISE_FIRST_FIRST=$(jq -r .[0].scoring.advise[0].points <<<"${output}")
#   ADVISE_FIRST_SECOND=$(jq -r .[0].scoring.advise[1].points <<<"${output}")
#   ADVISE_FIRST_THIRD=$(jq -r .[0].scoring.advise[2].points <<<"${output}")
#   (( ADVISE_FIRST_FIRST >= ADVISE_FIRST_SECOND >= ADVISE_FIRST_THIRD ))

#   # report 2 - criticals - magnitude sort/lowest number first
#   CRITICAL_SECOND_FIRST=$(jq -r .[1].scoring.critical[0].points <<<"${output}")
#   CRITICAL_SECOND_SECOND=$(jq -r .[1].scoring.critical[1].points <<<"${output}")
#   (( CRITICAL_SECOND_FIRST <= CRITICAL_SECOND_SECOND ))

#   # report 2 - advisories - magnitude sort/highest number first
#   ADVISE_SECOND_FIRST=$(jq -r .[1].scoring.advise[0].points <<<"${output}")
#   ADVISE_SECOND_SECOND=$(jq -r .[1].scoring.advise[1].points <<<"${output}")
#   ADVISE_SECOND_THIRD=$(jq -r .[1].scoring.advise[2].points <<<"${output}")
#   (( ADVISE_SECOND_FIRST >= ADVISE_SECOND_SECOND >= ADVISE_SECOND_THIRD ))
# }

# @test "returns deterministic report output" {
#   run _app "${TEST_DIR}/asset/score-2-pod-serviceaccount.yml"
#   assert_success

#   RUN_1_SIGNATURE=$(echo "${output}" | sha1sum)

#   run _app "${TEST_DIR}/asset/score-2-pod-serviceaccount.yml"
#   assert_success

#   RUN_2_SIGNATURE=$(echo "${output}" | sha1sum)

#   run _app "${TEST_DIR}/asset/score-2-pod-serviceaccount.yml"
#   assert_success

#   RUN_3_SIGNATURE=$(echo "${output}" | sha1sum)

#   assert [ "${RUN_1_SIGNATURE}" = "${RUN_2_SIGNATURE}" ]
#   assert [ "${RUN_1_SIGNATURE}" = "${RUN_3_SIGNATURE}" ]
# }
