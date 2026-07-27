# tf-aws-module_primitive-vpc_security_group_egress_rule

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Tests](https://github.com/launchbynttdata/tf-aws-module_primitive-vpc_security_group_egress_rule/actions/workflows/tests.yml/badge.svg)](https://github.com/launchbynttdata/tf-aws-module_primitive-vpc_security_group_egress_rule/actions/workflows/tests.yml)
[![Static Analysis](https://github.com/launchbynttdata/tf-aws-module_primitive-vpc_security_group_egress_rule/actions/workflows/static-analysis.yml/badge.svg)](https://github.com/launchbynttdata/tf-aws-module_primitive-vpc_security_group_egress_rule/actions/workflows/static-analysis.yml)

## Overview

This Terraform primitive module manages a single AWS VPC security group **egress rule**. It provides a lightweight wrapper around the [`aws_vpc_security_group_egress_rule`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_security_group_egress_rule) resource with sensible defaults and basic validation.

## Features

- **Single egress rule management** - Creates and manages one security group egress rule
- **Multiple destination types** - Supports IPv4/IPv6 CIDR blocks, prefix lists, and security group references
- **Protocol flexibility** - Handles TCP, UDP, ICMP, and custom protocols
- **Lightweight validation** - Basic input validation with clear error messages
- **Terraform 1.5+ compatibility** - Uses modern validation patterns and check blocks

## Usage

### Basic IPv4 CIDR Example

```hcl
module "ssh_egress" {
  source = "git::https://github.com/launchbynttdata/tf-aws-module_primitive-vpc_security_group_egress_rule.git?ref=1.0.0"

  security_group_id = aws_security_group.app.id
  ip_protocol       = "tcp"
  from_port         = 22
  to_port           = 22
  cidr_ipv4         = "10.0.1.0/24"
  description       = "Allow SSH to management subnet"
}
```

### IPv6 CIDR Example

```hcl
module "https_egress_ipv6" {
  source = "git::https://github.com/launchbynttdata/tf-aws-module_primitive-vpc_security_group_egress_rule.git?ref=1.0.0"

  security_group_id = aws_security_group.web.id
  ip_protocol       = "tcp"
  from_port         = 443
  to_port           = 443
  cidr_ipv6         = "::/0"
  description       = "Allow HTTPS to anywhere IPv6"
}
```

### Prefix List Example

```hcl
module "s3_egress" {
  source = "git::https://github.com/launchbynttdata/tf-aws-module_primitive-vpc_security_group_egress_rule.git?ref=1.0.0"

  security_group_id = aws_security_group.app.id
  ip_protocol       = "tcp"
  from_port         = 443
  to_port           = 443
  prefix_list_id    = data.aws_ec2_managed_prefix_list.s3.id
  description       = "Allow HTTPS to S3"
}
```

### Security Group to Security Group Example

```hcl
module "database_egress" {
  source = "git::https://github.com/launchbynttdata/tf-aws-module_primitive-vpc_security_group_egress_rule.git?ref=1.0.0"

  security_group_id            = aws_security_group.app.id
  ip_protocol                  = "tcp"
  from_port                    = 5432
  to_port                      = 5432
  referenced_security_group_id = aws_security_group.database.id
  description                  = "Allow PostgreSQL to database tier"
}
```

### ICMP Example

```hcl
module "icmp_egress" {
  source = "git::https://github.com/launchbynttdata/tf-aws-module_primitive-vpc_security_group_egress_rule.git?ref=1.0.0"

  security_group_id = aws_security_group.app.id
  ip_protocol       = "icmp"
  from_port         = -1
  to_port           = -1
  cidr_ipv4         = "10.0.0.0/8"
  description       = "Allow ICMP to private networks"
}
```

## Examples

The module includes comprehensive examples demonstrating various use cases:

- **[complete](examples/complete/)** - Multiple egress rules with different configurations
- **[minimal](examples/minimal/)** - Simplest possible configuration
- **[simple](examples/simple/)** - Basic example used for integration testing
- **[ipv6](examples/ipv6/)** - IPv6 CIDR blocks demonstration
- **[prefix_list](examples/prefix_list/)** - AWS managed prefix list usage
- **[sg_to_sg](examples/sg_to_sg/)** - Security group to security group references

Each example is runnable and includes:
- Complete Terraform configuration
- Sample `test.tfvars` file
- Documentation and usage instructions

## Validation Rules

The module enforces several validation rules to ensure proper configuration:

1. **Exactly one destination** - Must specify exactly one of: `cidr_ipv4`, `cidr_ipv6`, `prefix_list_id`, or `referenced_security_group_id`
2. **Protocol-port relationship** - TCP and UDP protocols require both `from_port` and `to_port`
3. **Port range validity** - `from_port` must be less than or equal to `to_port`
4. **Valid port numbers** - Ports must be between 1 and 65535 (or -1 for ICMP)

## Testing

### Prerequisites

```bash
# Install required tools
make configure
pre-commit install
```

### Local Testing

```bash
# Run all quality checks
make check

# Individual checks
make fmt       # Format code
make validate  # Validate syntax
make lint      # Static analysis
make test      # Unit tests
```

### Integration Testing

```bash
# Run integration tests on all examples
make integration-test

# Test specific example
cd examples/simple
terraform init
terraform plan -var-file=test.tfvars
terraform apply -var-file=test.tfvars
terraform destroy -var-file=test.tfvars
```

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 1.5.0 |
| aws | ~> 5.100.0 |

## Providers

| Name | Version |
|------|---------|
| aws | ~> 5.100.0 |

## Resources

| Name | Type |
|------|------|
| [aws_vpc_security_group_egress_rule.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_security_group_egress_rule) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| security_group_id | The ID of the security group to which this egress rule will be attached. | `string` | n/a | yes |
| ip_protocol | The IP protocol name or number. Use '-1' to specify all protocols. Protocol numbers: https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml | `string` | n/a | yes |
| from_port | The start of port range for the TCP and UDP protocols, or an ICMP type number. Required for tcp and udp protocols. Use -1 for ICMP type. | `number` | `null` | no |
| to_port | The end of port range for the TCP and UDP protocols, or an ICMP code. Required for tcp and udp protocols. Use -1 for ICMP code. | `number` | `null` | no |
| cidr_ipv4 | The destination IPv4 CIDR range for this egress rule. Mutually exclusive with cidr_ipv6, prefix_list_id, and referenced_security_group_id. | `string` | `null` | no |
| cidr_ipv6 | The destination IPv6 CIDR range for this egress rule. Mutually exclusive with cidr_ipv4, prefix_list_id, and referenced_security_group_id. | `string` | `null` | no |
| prefix_list_id | The ID of the prefix list for the destination of this egress rule. Mutually exclusive with cidr_ipv4, cidr_ipv6, and referenced_security_group_id. | `string` | `null` | no |
| referenced_security_group_id | The ID of the destination security group for this egress rule. Mutually exclusive with cidr_ipv4, cidr_ipv6, and prefix_list_id. | `string` | `null` | no |
| description | The description of this egress rule. | `string` | `null` | no |
| tags | A map of tags to assign to the egress rule. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| id | The Terraform resource ID of the security group egress rule. |
| security_group_rule_id | The AWS-assigned unique identifier for the security group rule. |
| security_group_id | The ID of the security group to which this egress rule is attached. |
| egress_rule_effective_source | A canonical string describing the effective destination for this egress rule (CIDR, prefix list, or security group). |
| arn | The ARN of the security group rule. |
| tags_all | A map of tags assigned to the resource, including those inherited from the provider default_tags. |

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for information on contributing to this module.

## License

Licensed under the Apache License, Version 2.0 (the "License").

You may obtain a copy of the License at [apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0).

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.100 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_vpc_security_group_egress_rule.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc_security_group_egress_rule) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cidr_ipv4"></a> [cidr\_ipv4](#input\_cidr\_ipv4) | The destination IPv4 CIDR range for this egress rule. Mutually exclusive with cidr\_ipv6, prefix\_list\_id, and referenced\_security\_group\_id. | `string` | `null` | no |
| <a name="input_cidr_ipv6"></a> [cidr\_ipv6](#input\_cidr\_ipv6) | The destination IPv6 CIDR range for this egress rule. Mutually exclusive with cidr\_ipv4, prefix\_list\_id, and referenced\_security\_group\_id. | `string` | `null` | no |
| <a name="input_description"></a> [description](#input\_description) | The description of this egress rule. | `string` | `null` | no |
| <a name="input_from_port"></a> [from\_port](#input\_from\_port) | The start of port range for the TCP and UDP protocols, or an ICMP type number. Required for tcp and udp protocols. Use -1 for ICMP type. | `number` | `null` | no |
| <a name="input_ip_protocol"></a> [ip\_protocol](#input\_ip\_protocol) | The IP protocol name or number. Use '-1' to specify all protocols. Protocol numbers: https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml | `string` | n/a | yes |
| <a name="input_prefix_list_id"></a> [prefix\_list\_id](#input\_prefix\_list\_id) | The ID of the prefix list for the destination of this egress rule. Mutually exclusive with cidr\_ipv4, cidr\_ipv6, and referenced\_security\_group\_id. | `string` | `null` | no |
| <a name="input_referenced_security_group_id"></a> [referenced\_security\_group\_id](#input\_referenced\_security\_group\_id) | The ID of the destination security group for this egress rule. Mutually exclusive with cidr\_ipv4, cidr\_ipv6, and prefix\_list\_id. | `string` | `null` | no |
| <a name="input_security_group_id"></a> [security\_group\_id](#input\_security\_group\_id) | The ID of the security group to which this egress rule will be attached. | `string` | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to assign to the egress rule. | `map(string)` | `{}` | no |
| <a name="input_to_port"></a> [to\_port](#input\_to\_port) | The end of port range for the TCP and UDP protocols, or an ICMP code. Required for tcp and udp protocols. Use -1 for ICMP code. | `number` | `null` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the security group rule. |
| <a name="output_egress_rule_effective_source"></a> [egress\_rule\_effective\_source](#output\_egress\_rule\_effective\_source) | A canonical string describing the effective destination for this egress rule (CIDR, prefix list, or security group). |
| <a name="output_id"></a> [id](#output\_id) | The Terraform resource ID of the security group egress rule. |
| <a name="output_security_group_id"></a> [security\_group\_id](#output\_security\_group\_id) | The ID of the security group to which this egress rule is attached. |
| <a name="output_security_group_rule_id"></a> [security\_group\_rule\_id](#output\_security\_group\_rule\_id) | The AWS-assigned unique identifier for the security group rule. |
| <a name="output_tags_all"></a> [tags\_all](#output\_tags\_all) | A map of tags assigned to the resource, including those inherited from the provider default\_tags. |
<!-- END_TF_DOCS -->
