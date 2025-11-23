# Go Project Structure

This is a typical Go project layout template. The directories below form the foundation of any Go project.

## Core Directories

### `/cmd`
Contains the application's main package and `main.go` file. This package should use minimal business logic or none at all. It serves to configure the entire application and define build rules and arguments. In other words, this is the technical entry point.

### `/internal`
Code that directly relates to the current application and won't be reused in other applications. This is the application's business logic - handlers, routes, migrations, etc.

**Important**: The `internal` directory has special functionality in Go. Packages inside `internal` cannot be imported via `$GOPATH` (or any other way) into other applications, as this code is intended only for this project and nowhere else. It's non-reusable by design.

### `/pkg`
Contains library code that doesn't directly depend on the current application. This is the opposite of `internal`. Place packages here that can be useful in this application and any other. Reusable code like HTTP helpers, message broker utilities, database tools, etc. This code can be used in other applications and has no dependency on the current one.

### `/vendor`
The name speaks for itself. This directory, often not tracked by git, contains libraries pulled and managed by a dependency manager. Code in this directory cannot be modified.

This directory may not exist in Go projects using `go mod`, since Go modules by default don't use vendor mode and can store all libraries outside the current project at `$GOPATH/pkg/mod`. Projects using `go mod` in standard mode (not vendor mode) won't have this directory.

## API & Web

### `/api`
Files related to the public API - Swagger, Proto, and other similar schemas.

### `/web`
Any static assets for web applications: JS, CSS, HTML. Also Go templates or components for JS frameworks, etc.

## Configuration & Scripts

### `/config`
Any configuration files required for the application.

### `/scripts`
Scripts that may be necessary for building, analyzing, or installing the project.

### `/build`
Files for build tools or similar purposes. For example, CI/CD configs or Dockerfiles.

## Deployment & Testing

### `/deployments`
Files for direct deployment of the application to orchestration systems: docker-compose, kubernetes/helm, mesos, terraform, bosh.

### `/test`
Can contain any tests and auxiliary files for them.

## Documentation & Examples

### `/docs`
Directory containing markdown descriptions for your application, documentation, or other useful information.

### `/examples`
Examples of using your application or library.

## Utilities & Assets

### `/tools`
Code for auxiliary tools that doesn't depend on packages in the `pkg` and `internal` directories.

### `/third_party`
Additional utilities like Swagger UI, GraphQL console, or other tools that don't directly relate to the project but are necessary for use in it. May also contain forked modified libraries.

### `/assets`
Images, icons, logos, etc.