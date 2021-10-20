# BadRobot (Kubernetes Operator Audit Tool)

---

- [Badrobot Overview](#overview)
  - [Structure](#structure)
  - [Roadmap](#roadmap)
- [Command line Usage](#command-line-usage)

## Overview
Badrobot is the internal R&D project codename for a Kubernetes Operator Audit Tool. The purpose of the project is to create an open source tool which audits Public or Private Kubernetes Operators. As Operators can have a large scope the initial version will be focused on specific resources and Operators which conform to the Operator SDK. If successful, the tool could be expanded to identify any resources which are associated with an Operator.

### Structure
The tool is supposed be a static code analyser which can run against a code repository and provide an "opinionated" score on the current Operator configuration. This is similar to KubeSec and will be structured in the same way, that is:
* code to scan a .yaml file
* a package of rules to apply against the .yaml file

### Roadmap
Phase 1 is focussed on service account and cluster-wide role permissions

## Command Line Usage
```bash
$ badrobot scan operator-manifest.yaml
```
> Currently there are example operator files to use for Badrobot under test/example-yaml
## Appendix
