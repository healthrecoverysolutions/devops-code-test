resource "docker_container" "wordpress" {
  name  = format("wordpress-%s", count.index)
  image = "wordpress:latest"
  restart = "always"
  count = var.wordpress_instance_count

  ports {
    internal = "80"
    external = format("9%03s", 1+count.index)
  }
}