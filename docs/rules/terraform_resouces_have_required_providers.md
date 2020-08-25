# terraform_resouces_have_required_providers

Require that all resources have version constraints declared through `required_providers`.

## Configuration

```hcl
rule "terraform_resources_have_required_providers" {
  enabled = true
}
```

## Examples

```hcl
terraform {
 required_providers {
 }
}

resource "template_test" "example" {
}

```

```
$ tflint
1 issue(s) found:

Warning: `Missing version constraint for provider "template" in "required_providers"

  on main.tf line 6:
   1: provider "template" {}

Reference: https://github.com/terraform-linters/tflint/blob/v0.18.0/docs/rules/terraform_resources_have_required_providers.md
```


## Why

Providers are plugins released on a separate rhythm from Terraform itself, and so they have their own version numbers. For production use, you should constrain the acceptable provider versions via configuration, to ensure that new versions with breaking changes will not be automatically installed by `terraform init` in future.

## How To Fix

Add the [`required_providers`](https://www.terraform.io/docs/configuration/terraform.html#specifying-required-provider-versions) block to the `terraform` configuration block and include current versions for all providers. For example:

```tf
terraform {
  required_providers {
    template = "~> 2.0"
  }
}
```

Provider version constraints can be specified using a [version argument within a provider block](https://www.terraform.io/docs/configuration/providers.html#provider-versions) for backwards compatability. This approach is now discouraged, particularly for child modules. This usage is unsupported by this rule.