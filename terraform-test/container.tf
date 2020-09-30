resource "docker_container" "wordpress" {
  name  = "wordpress-${count.index}"
  image = "wordpress:latest"
  restart = "always"
  count = var.wordpress_instance_count

  ports {
    internal = "80"
    external = format("%d", 9001 + count.index)
  }
}