#!/usr/bin/env bash

load '_helper'

setup() {
    _global_setup
}

teardown() {
  _global_teardown
}

# ---

@test "errors with no filename" {
  run _app
  assert_failure
}

@test "errors with no filename - output logs (local)" {
  run "${BIN_DIR:-}"/badrobot scan
  assert_failure
  assert_line "Error: file path is required"
}

@test "errors with invalid file" {
  run _app somefile.yaml
  assert_failure
  assert_file_not_found
}

@test "errors with empty file" {
  run _app "${TEST_DIR}/asset/reg-empty-file"
  assert_failure
  assert_invalid_input
}

@test "errors with empty file (json)" {
  run _app "${TEST_DIR}/asset/reg-empty-file" --json
  assert_invalid_input
  assert_failure
}

@test "errors with empty JSON (json)" {
  run _app "${TEST_DIR}/asset/reg-empty-json-file" --json
  assert_invalid_input
  assert_failure
}

# TBD KGW - kubekind no longer working for schema validation, needs replacing
# @test "errors with invalid kind" {
#   run _app "${TEST_DIR}/asset/reg-invalid-kind.yaml"
#   assert_invalid_input
#   assert_failure
# }

# TBD KGW - kubekind no longer working for schema validation, needs replacing
# @test "errors with invalid schema" {
#   run _app "${TEST_DIR}/asset/reg-invalid-schema.yaml"
#   assert_invalid_input
#   assert_failure
# }
