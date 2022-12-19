locals {
  from_port = var.from_port != null ? var.from_port : var.port
  to_port   = var.to_port != null ? var.to_port : var.port
}
