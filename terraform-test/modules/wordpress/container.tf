resource "docker_container" "wordpress" {
  name  = "wordpress-${var.port_number}"
  image = "wordpress:latest"
  restart = "always"

  ports {
    internal = "80"
    external = var.port_number
  }
}
