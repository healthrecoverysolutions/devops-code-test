output "ip" {
  value = join(", ", docker_container.wordpress.*.ip_address)
}