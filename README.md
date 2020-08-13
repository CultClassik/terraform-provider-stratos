# Terraform Provider Stratos

Run the following command to build the provider

```shell
go build -o terraform-provider-stratos
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
terraform init && terraform apply
```

## Resources

1. server_role
2. server_support_group
3. stratos_config

 ## Data Sources

 1. server_role
 2. server_support_group
 