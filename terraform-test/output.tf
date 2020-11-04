output "ip-csl" {
    value = join(",", module.wordpress.*.ip )
}
