# terratestutilities

A number of helper functions to be used in other go modules. For instance checking that appropriate environment variables have been set to authenticate through the Azure Cli.

We particulary use this alongside [Terratest](https://github.com/gruntwork-io/terratest) to control Azure Cli authentication and to set whether a terraform execution should do a plan, apply, destroy or all.

Checkout [terraform-template](https://github.com/hiscox/terraform-template) and inspect the tests folder to see an example setup of how we use this module to control the workflow in a pipeline.

## Getting Started

Add this module as an import:

```go
import (
  util "github.com/hiscox/go-utilities"
)
```

Access methods like so:

```go
cSecret, cID, tenID, subID, err := util.AzCLiAuth()
if err != nil {
  // handle it
}
```

## Build and Test

TODO: add an azure-pipelines.yml to run the tests
