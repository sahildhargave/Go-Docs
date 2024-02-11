# MOD documentation

## Project Structure of most of the Go Projects

# project name/
|-- go.mod
|-- go.sum
|-- main.go
|-- 02variable/
| |-- mod.go
|-- other_directories/
| |-- ...

- **`go.mod`**: The main Go module file.
- **`go.sum`**: The checksum file to lock dependencies versions.
- **`main.go`**: The entry point of your application.
- **`02variable/`**: Directory containing Go code related to variables (example).
- **`other_directories/`**: Additional directories as needed.

## Getting Started

 ## Versioning

We use [Semantic Versioning (SemVer)](https://semver.org/) for versioning. The format is `MAJOR.MINOR.PATCH`.

| Version | Description |
| ------- | ----------- |
| 1.0.0   | Initial release with basic features |
| 1.1.0   | Added feature X and Y |
| 1.1.1   | Bugfix for issue Z |
| 2.0.0   | Major update with breaking changes |
| ...     | More details about each version |

# Table of Contents

1. [Introduction](#introduction)
2. [Modules, Packages, and Versions](#modules-packages-and-versions)
   - [Module Paths](#module-paths)
   - [Versions](#versions)
   - [Pseudo-Versions](#pseudo-versions)
   - [Major Version Suffixes](#major-version-suffixes)
3. [Resolving a Package to a Module](#resolving-a-package-to-a-module)
4. [go.mod Files](#gomod-files)
   - [Lexical Elements](#lexical-elements)
   - [Module Paths and Versions](#module-paths-and-versions)
   - [Grammar](#grammar)
   - [Module Directive](#module-directive)
   - [go Directive](#go-directive)
   - [toolchain Directive](#toolchain-directive)
   - [require Directive](#require-directive)
   - [exclude Directive](#exclude-directive)
   - [replace Directive](#replace-directive)
   - [retract Directive](#retract-directive)
5. [Automatic Updates](#automatic-updates)
   - [Minimal Version Selection (MVS)](#minimal-version-selection-mvs)
   - [Replacement](#replacement)
   - [Exclusion](#exclusion)
   - [Upgrades](#upgrades)
   - [Downgrade](#downgrade)
6. [Module Graph Pruning](#module-graph-pruning)
   - [Lazy Module Loading](#lazy-module-loading)
7. [Workspaces](#workspaces)
   - [go.work Files](#gowork-files)
   - [Lexical Elements](#lexical-elements-1)
   - [Grammar](#grammar-1)
   - [go Directive](#go-directive-1)
   - [toolchain Directive](#toolchain-directive-1)
   - [use Directive](#use-directive)
   - [replace Directive](#replace-directive-1)
8. [Compatibility with Non-Module Repositories](#compatibility-with-non-module-repositories)
   - [+incompatible Versions](#incompatible-versions)
   - [Minimal Module Compatibility](#minimal-module-compatibility)
9. [Module-Aware Commands](#module-aware-commands)
   - [Build Commands](#build-commands)
   - [Vendoring](#vendoring)
   - [go get](#go-get)
   - [go install](#go-install)
   - [go list -m](#go-list--m)
   - [go mod download](#go-mod-download)
   - [go mod edit](#go-mod-edit)
   - [go mod graph](#go-mod-graph)
   - [go mod init](#go-mod-init)
   - [go mod tidy](#go-mod-tidy)
   - [go mod vendor](#go-mod-vendor)
   - [go mod verify](#go-mod-verify)
   - [go mod why](#go-mod-why)
   - [go version -m](#go-version--m)
   - [go clean -modcache](#go-clean--modcache)
10. [Version Queries](#version-queries)
11. [Module Commands Outside a Module](#module-commands-outside-a-module)
   - [go work init](#go-work-init)
   - [go work edit](#go-work-edit)
   - [go work use](#go-work-use)
   - [go work sync](#go-work-sync)
12. [Module Proxies](#module-proxies)
   - [GOPROXY Protocol](#goproxy-protocol)
   - [Communicating with Proxies](#communicating-with-proxies)
   - [Serving Modules Directly from a Proxy](#serving-modules-directly-from-a-proxy)
13. [Version Control Systems](#version-control-systems)
   - [Finding a Repository for a Module Path](#finding-a-repository-for-a-module-path)
   - [Mapping Versions to Commits](#mapping-versions-to-commits)
   - [Mapping Pseudo-Versions to Commits](#mapping-pseudo-versions-to-commits)
   - [Mapping Branches and Commits to Versions](#mapping-branches-and-commits-to-versions)
   - [Module Directories Within a Repository](#module-directories-within-a-repository)
   - [Special Case for LICENSE Files](#special-case-for-license-files)
   - [Controlling Version Control Tools with GOVCS](#controlling-version-control-tools-with-govcs)
14. [Module Zip Files](#module-zip-files)
   - [File Path and Size Constraints](#file-path-and-size-constraints)
15. [Private Modules](#private-modules)
   - [Private Proxy Serving All Modules](#private-proxy-serving-all-modules)
   - [Private Proxy Serving Private Modules](#private-proxy-serving-private-modules)
   - [Direct Access to Private Modules](#direct-access-to-private-modules)
   - [Passing Credentials to Private Proxies](#passing-credentials-to-private-proxies)
   - [Passing Credentials to Private Repositories](#passing-credentials-to-private-repositories)
16. [Privacy](#privacy)
17. [Module Cache](#module-cache)
   - [Authenticating Modules](#authenticating-modules)
   - [go.sum Files](#gosum-files)
   - [Checksum Database](#checksum-database)
18. [Environment Variables](#environment-variables)
19. [Glossary](#glossary)


Gorilla Mux is a powerful HTTP router for Go applications, providing a flexible and expressive way to define and handle HTTP routes. This project demonstrates how to use Gorilla Mux in a Go web application.

## Features

- **Router**: Gorilla Mux serves as a router for HTTP requests, enabling the definition of routes and association with corresponding handler functions.

- **URL Patterns**: Supports complex URL patterns, including path variables and regular expressions, offering flexibility in defining API endpoints.

- **Middleware Support**: Allows the integration of middleware for pre-processing or post-processing tasks, such as authentication, logging, or request/response manipulation.

- **Subrouters**: Enables the creation of subrouters, allowing the organization of routes into modular and reusable components.

## How to Use Gorilla Mux?

### Installation

You can install Gorilla Mux using the following command:

## **Repository:**

```bash
go get -u github.com/gorilla/mux

### Router Initialization
2. **Initialize Go Workspace:**

    ```bash
    go work init
    ```

3. **Use the Workspace:**

    ```bash
    go work use .
    ```

4. **Build the Project:**

    ```bash
    go build .
    ```

5. **Verify Modules:**

    ```bash
    go mod verify
    ```

6. **List Modules:**

    ```bash
    go list all
    ```

7. **List Modules (Module Mode):**

    ```bash
    go list -m all
    ```

8. **List Module Versions:**

    ```bash
    go list -m -versions github.com/gorilla/mux
    ```

9. **Tidy Modules:**

    ```bash
    go mod tidy
    ```

10. **Why a Module:**

    ```bash
    go mod why github.com/gorilla/mux
    ```

11. **Module Graph:**

    ```bash
    go mod graph
    ```

12. **Edit Modules:**

    ```bash
    go mod edit
    ```

13. **Vendor Modules:**

    ```bash
    go mod vendor
    ```

14. **Run with Vendor Modules:**

    ```bash
    go run -mod=vendor main.go
    ```

