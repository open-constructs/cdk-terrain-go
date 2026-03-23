# CDK Terrain Go

CDK Terrain (CDKTN) is a community fork of the Cloud Development Kit for Terraform (CDKTF).
CDKTN allows you to use familiar programming languages to define cloud infrastructure
and provision it through HashiCorp Terraform or OpenTofu. This gives you access to the
entire Terraform/OpenTofu ecosystem without learning HashiCorp Configuration Language (HCL)
and lets you leverage the power of your existing toolchain for testing, dependency management, etc.

This repository contains the Go bindings for CDK Terrain.

## Installation

```bash
go get github.com/open-constructs/cdk-terrain-go/cdktn
```

## Usage

```go
import "github.com/open-constructs/cdk-terrain-go/cdktn"
```

## Documentation

Refer to the [CDKTN documentation](https://cdktn.io/docs) for more detail about how to build and manage CDKTN applications, including:

- [Application Architecture](https://cdktn.io/docs/concepts/cdktn-architecture): Learn the tools and processes that CDKTN uses to leverage the Terraform ecosystem and convert code into Terraform configuration files.

- [Project Setup](https://cdktn.io/docs/create-and-deploy/project-setup): Learn how to create a new CDKTN project from a pre-built or custom template.

- [Examples](https://cdktn.io/docs/examples-and-guides/examples): Reference example projects in every supported language.

## Community

- Ask a question on the [cdk.dev - #cdk-terrain channel](https://cdk.dev).
- Report a [bug](https://github.com/open-constructs/cdk-terrain/issues/new?assignees=&labels=bug&template=bug-report.md&title=) or request a new [feature](https://github.com/open-constructs/cdk-terrain/issues/new?assignees=&labels=enhancement&template=feature-request.md&title=).
- Browse all [open issues](https://github.com/open-constructs/cdk-terrain/issues).

## Source

This module is generated from [open-constructs/cdk-terrain](https://github.com/open-constructs/cdk-terrain).

## License

This project is licensed under the [MPL-2.0](./LICENSE).
