# Complete Example

This example demonstrates comprehensive usage of the `tf-aws-module_primitive-vpc_security_group_egress_rule` module with multiple egress rules.

## Features Demonstrated

- Multiple egress rules on the same security group
- SSH access (port 22) from IPv4 CIDR
- HTTPS access (port 443) from IPv4 CIDR
- Custom port range (8080-8090) from IPv4 CIDR

## Usage

```bash
terraform init
terraform plan -var-file=test.tfvars
terraform apply -var-file=test.tfvars
terraform destroy -var-file=test.tfvars
```

## Resources Created

- 1 VPC
- 1 Security Group
- 3 Security Group Egress Rules

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.100 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_egress_custom_range"></a> [egress\_custom\_range](#module\_egress\_custom\_range) | ../../ | n/a |
| <a name="module_egress_https_ipv4"></a> [egress\_https\_ipv4](#module\_egress\_https\_ipv4) | ../../ | n/a |
| <a name="module_egress_ssh_ipv4"></a> [egress\_ssh\_ipv4](#module\_egress\_ssh\_ipv4) | ../../ | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_default_security_group.default](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_security_group) | resource |
| [aws_security_group.test](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group) | resource |
| [aws_vpc.test](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_custom_cidr_ipv4"></a> [custom\_cidr\_ipv4](#input\_custom\_cidr\_ipv4) | IPv4 CIDR block allowed for custom port range | `string` | `"10.0.3.0/24"` | no |
| <a name="input_https_cidr_ipv4"></a> [https\_cidr\_ipv4](#input\_https\_cidr\_ipv4) | IPv4 CIDR block allowed for HTTPS access | `string` | `"10.0.2.0/24"` | no |
| <a name="input_resource_name_prefix"></a> [resource\_name\_prefix](#input\_resource\_name\_prefix) | Prefix for resource names | `string` | `"sger-complete"` | no |
| <a name="input_ssh_cidr_ipv4"></a> [ssh\_cidr\_ipv4](#input\_ssh\_cidr\_ipv4) | IPv4 CIDR block allowed for SSH access | `string` | `"10.0.1.0/24"` | no |
| <a name="input_vpc_cidr"></a> [vpc\_cidr](#input\_vpc\_cidr) | CIDR block for the VPC | `string` | `"10.0.0.0/16"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_custom_range_egress_rule_id"></a> [custom\_range\_egress\_rule\_id](#output\_custom\_range\_egress\_rule\_id) | ID of the custom range egress rule |
| <a name="output_effective_source"></a> [effective\_source](#output\_effective\_source) | Effective destination for SSH egress rule |
| <a name="output_egress_rule_arn"></a> [egress\_rule\_arn](#output\_egress\_rule\_arn) | ARN of the SSH egress rule |
| <a name="output_egress_rule_id"></a> [egress\_rule\_id](#output\_egress\_rule\_id) | ID of the SSH egress rule |
| <a name="output_https_egress_rule_id"></a> [https\_egress\_rule\_id](#output\_https\_egress\_rule\_id) | ID of the HTTPS egress rule |
| <a name="output_security_group_id"></a> [security\_group\_id](#output\_security\_group\_id) | ID of the test security group |
<!-- END_TF_DOCS -->
