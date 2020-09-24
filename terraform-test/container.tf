resource "docker_container" "wordpress" {
  name  = "wordpress"
  image = "wordpress:latest"
  restart = "always"

  ports {
    internal = "80"
    external = "9001"
  }
}