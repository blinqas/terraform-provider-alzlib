# ALZLib Terraform Provider

_This provider is built on the [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk)._

The ALZLib provider uses the [ALZLib](https://github.com/matt-FFFFFF/alzlib) library to provide ALZ archetype data resources to Terraform.

The data sources resources that it produces are complex objects, with nested maps.
This output is designed to be used with a Terraform module that will process this data to deploy resources.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.x
- [Go](https://golang.org/doc/install) >= 1.18

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```sh
go install
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

_TBC_

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

```sh
make testacc
```
