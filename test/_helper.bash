#!/usr/bin/env bash

load './bin/bats-support/load'
load './bin/bats-assert/load'

export TEST_DIR="."

export BIN_DIR='../dist/'

_global_setup() {
    [ ! -f "${BATS_PARENT_TMPNAME}".skip ] || skip "skip remaining tests"
}

_global_teardown() {
    if [ -z "$BATS_TEST_COMPLETED" ]; then
      touch "${BATS_PARENT_TMPNAME}".skip
    fi
}

_test_description_matches_regex() {
  [[ "${BATS_TEST_DESCRIPTION}" =~ ${1} ]]
}

_app() {
  local ARGS="${@:-}"
  if [[ "${BIN_DIR}" != "" ]]; then
    # remove --json flags
    ARGS=$(echo "${ARGS}" | sed -E 's,--json,,g')
  fi
  "${BIN_DIR}"/badrobot scan "${ARGS}";
}

assert_zero_points() {
  assert_output --regexp ".*with a score of 0 points.*"
  assert_failure
}

assert_lt_zero_points() {
  assert_output --regexp ".*\with a score of \-[1-9][0-9]* points.*"
  assert_failure
}

assert_file_not_found() {
  assert_output --regexp ".*File somefile.yaml does not exist.*" \
    || assert_output --regexp ".*no such file or directory.*"  \
    || assert_output --regexp ".*Invalid input.*"
}

assert_invalid_input() {
  assert_output --regexp '  "message": "Invalid input"' \
    || assert_output --regexp ".*Kubernetes kind not found.*" \
    || assert_output --regexp ".*no such file or directory.*" \
    || assert_output --regexp ".*Invalid input.*"
}

assert_failure_local() {
  assert_failure
}
