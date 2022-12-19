resource "aws_security_group_rule" "rule" {
  security_group_id        = var.security_group_id
  type                     = var.type
  from_port                = local.from_port
  to_port                  = local.to_port
  protocol                 = var.protocol
  source_security_group_id = var.source_security_group_id
  cidr_blocks              = var.source_cidr_blocks
  ipv6_cidr_blocks         = var.source_ipv6_cidr_blocks
}
