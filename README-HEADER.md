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

Integrate this module like so:

```hcl
module "sg_rule" {
  source = "github.com/pbs/terraform-aws-sg-rule-module?ref=x.y.z"

  port = 6379

  security_group_id        = module.redis.sg_ids[0]
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
