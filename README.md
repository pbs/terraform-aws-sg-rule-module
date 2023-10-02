# PBS TF SG Rule Module

## Installation

### Using the Repo Source

Use this URL for the source of the module. See the usage examples below for more details.

```hcl
github.com/pbs/terraform-aws-sg-rule-module?ref=x.y.z
```

### Alternative Installation Methods

More information can be found on these install methods and more in [the documentation here](./docs/general/install).

## Usage

This module provisions a security group rule. Use in conjunction with other modules to modify ingress and egress rules on security groups provisioned by them.

Note that each security group rule requires an explicit description. Try to make these as descriptive as possible.

Integrate this module like so:

```hcl
module "sg_rule" {
  source = "github.com/pbs/terraform-aws-sg-rule-module?ref=x.y.z"

  security_group_id = module.redis.sg_ids[0]

  description = "Allow Lambda ${module.lambda.name} to access Redis"

  port                     = 6379
  source_security_group_id = module.lambda.sg
}
```

## Adding This Version of the Module

If this repo is added as a subtree, then the version of the module should be close to the version shown here:

`x.y.z`

Note, however that subtrees can be altered as desired within repositories.

Further documentation on usage can be found [here](./docs).

Below is automatically generated documentation on this Terraform module using [terraform-docs][terraform-docs]

---

[terraform-docs]: https://github.com/terraform-docs/terraform-docs

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3.2 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 4.5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.19.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_security_group_rule.rule](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group_rule) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_description"></a> [description](#input\_description) | Description of the rule. This must clearly describe the purpose of the rule. | `string` | n/a | yes |
| <a name="input_security_group_id"></a> [security\_group\_id](#input\_security\_group\_id) | The ID of the security group that contains the rule. | `string` | n/a | yes |
| <a name="input_from_port"></a> [from\_port](#input\_from\_port) | The start port | `number` | `null` | no |
| <a name="input_port"></a> [port](#input\_port) | The port to allow. | `number` | `null` | no |
| <a name="input_protocol"></a> [protocol](#input\_protocol) | The protocol to allow. Valid values are tcp, udp, and all. | `string` | `"tcp"` | no |
| <a name="input_source_cidr_blocks"></a> [source\_cidr\_blocks](#input\_source\_cidr\_blocks) | A list of CIDR blocks to allow access from. | `list(string)` | `null` | no |
| <a name="input_source_ipv6_cidr_blocks"></a> [source\_ipv6\_cidr\_blocks](#input\_source\_ipv6\_cidr\_blocks) | A list of IPv6 CIDR blocks to allow access from. | `list(string)` | `null` | no |
| <a name="input_source_security_group_id"></a> [source\_security\_group\_id](#input\_source\_security\_group\_id) | The ID of the security group that is allowed access. | `string` | `null` | no |
| <a name="input_to_port"></a> [to\_port](#input\_to\_port) | The end port | `number` | `null` | no |
| <a name="input_type"></a> [type](#input\_type) | The type of rule to create. Valid values are egress, ingress, and all. | `string` | `"ingress"` | no |

## Outputs

No outputs.
