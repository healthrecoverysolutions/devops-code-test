module "wordpress" {
  count = var.wordpress_instance_count
  source = "./modules/wordpress"
  port_number = 9001+count.index
}
