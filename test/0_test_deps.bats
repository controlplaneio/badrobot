#!/usr/bin/env bash

load '_helper'

setup() {
  _global_setup
}

teardown() {
  _global_teardown
}

@test "test dep - jq is installed" {
  run command -v jq

  assert_success
}
