# Environment Provider

The Environment provider is a utility provider for including local environment variables as part of a Terraform configuration.  The provider itself requires no configuration.

## Installation

1. Download an appropriate binary from [Releases page](https://github.com/gaarutyunov/terraform-provider-environment/releases)

Or Install:

```shell
go install github.com/gaarutyunov/terraform-provider-environment@<Version of release>
```

2. Copy:

```shell
mkdir -p .terraform/plugins/registry.terraform.io/hashicorp/environment/<Version of downloaded release>/<GOOS>_<GOARCH>
cp <Path to binary> .terraform/plugins/registry.terraform.io/hashicorp/environment/1.0.0/<GOOS>_<GOARCH>
```

For example:
```shell
mkdir -p .terraform/plugins/registry.terraform.io/hashicorp/environment/1.0.0/darwin_amd64
cp /usr/local/go/bin/terraform-provider-environment .terraform/plugins/registry.terraform.io/hashicorp/environment/1.0.0/darwin_amd64
```

## environment_variable Data Source

 The environment_variable data source provides a mechanism for a Terraform configuration to read values of local environment variables and incorporate them into a Terraform configuration.  This helps make it easier to make the Terraform script machine independent.  For example, when configuring a managed Kubernetes cluster on AWS, the Kubernetes config file needs to be modified to allow Kubernetes commands to successfully connect and authenticate to the Kubernetes API server.  The config file is typically written to `~/.kube/` directory, but this path does not work on Windows.  By providing access to the HOME environment variable, it is possible to compute the path based on the user's home directory without assuming an operating system.
 
 ```hcl
terraform {
  required_providers {
    environment = {
      source = "terraform-provider-environment"
      version = ">= 1.0.0"
    }
  }
}
 
provider "environment" {}

data "environment_variable" "HOME" {
  name = "HOME"
  fail_if_empty = true
  normalize_file_path = true
}

provider "kubernetes" {
  config_file = "${data.environment_variable.HOME.value}/.kube/my-cluster-config"
}
```

### Argument Reference
___
The following arguments are supported:

- `name` - (Required) The name of the environment variable whose value should be used.
- `default` - (Optional) The default value to use should the environment variable not be set or set to an empty string.
- `fail_if_empty` - (Optional) Whether or not the data source read should fail if the final value of the environment variable (after applying the `default`, if specified), if empty.  The default value is `false`.
- `normalize_file_path` - (Optional) Whether or not to treat the final value as a file path, which means making sure to quote any backslashes in the path when running on the Windows platform.  The default value is `false`.

### Attributes
___
- `value` - The final value of the environment variable (after applying the `default`, if specified.
