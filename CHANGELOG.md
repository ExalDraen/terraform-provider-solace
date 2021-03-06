# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## Unreleased

### Known Issues

### Added

### Changed

## 0.1.1

### Changed

* Switch to [Go modules](https://github.com/golang/go/wiki/Modules) for dependencies
* Vendor dependencies to avoid issues with missing dependencies at build time.
* Switch to Go `1.11.x` in TravisCI
* Use recommended method of running `goreleaser` in TravisCI.

## 0.1.0

### Added

* Added basic acceptance test for `solace_msgvpn`. Does not yet run on CI.

### Changed

* Remove unused `ClientAndAuth` struct.
* Update the terraform helper dependency to `0.12.7`. This means the provider now requires Terraform > 0.12.

## 0.0.2

### Changed

* Cleaned up source formatting with `gofmt -s`
* Updated `semp-client` to allow updating resource properties to their "null" value (solving the Known Issue in `v.0.0.1`).
* Fleshed out config examples.

## 0.0.1

### Known Issues

The models used for parameters by the `semp-client` library are tagged with the label `omitempty`. Per [the json docs](https://golang.org/pkg/encoding/json/#Marshal) this means that these fields will *not* be included in a request if have an "empty value".
This is a problem when trying to update an attribute from a *non-empty* to an *empty* value, for example to change a boolean from `false` to `true`: such a value is recognized as an "empty value" and not included in any update (`PATCH`) requests.

### Added

* Basic scaffolding for terraform provider
* Very basic MSG VPN resource with a few select attributes, CRUD and import capability
* ACL Profile resource with CRUD
* ACL profile client connection exceptions
* ACL profile publish, & subscribe exceptions
* Basic Client profile resource with a few select attributes, CRUD and import capability
* Basic Client username resource with a few select attributes, CRUD and import capability
* Makefile to build, lint and test (but no tests yet). Includes a slew of `gometalinter` linters.
