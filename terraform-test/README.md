**You will need Docker running locally to test this**

## Instruction

This creates a Docker container for Wordpress on the local Docker machine.

Update the configuration to launch N instances of the container based on the `wordpress_instance_count` variable using ports 9001 through 9xxx for the respective number of instances provided. The output should be changed to reflect the private IP of each container in a comma-separated list (e.g. `ip = 172.1.0.1,172.1.0.2,172.1.0.3,172.1.0.4,172.1.0.5`)

In your submitted PR, please put the output from both `terraform plan` and the output after `terraform apply` below with 3 containers provided as value on the `wordpress_instance_count` variable.

## Terrform Plan

terraform plan
Refreshing Terraform state in-memory prior to plan...
The refreshed state will be used to calculate this plan, but will not be
persisted to local or remote state storage.


------------------------------------------------------------------------

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # docker_container.wordpress[0] will be created
  + resource "docker_container" "wordpress" {
      + attach           = false
      + bridge           = (known after apply)
      + command          = (known after apply)
      + container_logs   = (known after apply)
      + dns              = (known after apply)
      + dns_opts         = (known after apply)
      + entrypoint       = (known after apply)
      + exit_code        = (known after apply)
      + gateway          = (known after apply)
      + hostname         = (known after apply)
      + id               = (known after apply)
      + image            = "wordpress:latest"
      + ip_address       = (known after apply)
      + ip_prefix_length = (known after apply)
      + ipc_mode         = (known after apply)
      + log_driver       = (known after apply)
      + log_opts         = (known after apply)
      + logs             = false
      + must_run         = true
      + name             = "wordpress-0"
      + network_data     = (known after apply)
      + read_only        = false
      + restart          = "always"
      + rm               = false
      + shm_size         = (known after apply)
      + start            = true
      + user             = (known after apply)
      + working_dir      = (known after apply)

      + ports {
          + external = 9001
          + internal = 80
          + ip       = "0.0.0.0"
          + protocol = "tcp"
        }
    }

  # docker_container.wordpress[1] will be created
  + resource "docker_container" "wordpress" {
      + attach           = false
      + bridge           = (known after apply)
      + command          = (known after apply)
      + container_logs   = (known after apply)
      + dns              = (known after apply)
      + dns_opts         = (known after apply)
      + entrypoint       = (known after apply)
      + exit_code        = (known after apply)
      + gateway          = (known after apply)
      + hostname         = (known after apply)
      + id               = (known after apply)
      + image            = "wordpress:latest"
      + ip_address       = (known after apply)
      + ip_prefix_length = (known after apply)
      + ipc_mode         = (known after apply)
      + log_driver       = (known after apply)
      + log_opts         = (known after apply)
      + logs             = false
      + must_run         = true
      + name             = "wordpress-1"
      + network_data     = (known after apply)
      + read_only        = false
      + restart          = "always"
      + rm               = false
      + shm_size         = (known after apply)
      + start            = true
      + user             = (known after apply)
      + working_dir      = (known after apply)

      + ports {
          + external = 9002
          + internal = 80
          + ip       = "0.0.0.0"
          + protocol = "tcp"
        }
    }

  # docker_container.wordpress[2] will be created
  + resource "docker_container" "wordpress" {
      + attach           = false
      + bridge           = (known after apply)
      + command          = (known after apply)
      + container_logs   = (known after apply)
      + dns              = (known after apply)
      + dns_opts         = (known after apply)
      + entrypoint       = (known after apply)
      + exit_code        = (known after apply)
      + gateway          = (known after apply)
      + hostname         = (known after apply)
      + id               = (known after apply)
      + image            = "wordpress:latest"
      + ip_address       = (known after apply)
      + ip_prefix_length = (known after apply)
      + ipc_mode         = (known after apply)
      + log_driver       = (known after apply)
      + log_opts         = (known after apply)
      + logs             = false
      + must_run         = true
      + name             = "wordpress-2"
      + network_data     = (known after apply)
      + read_only        = false
      + restart          = "always"
      + rm               = false
      + shm_size         = (known after apply)
      + start            = true
      + user             = (known after apply)
      + working_dir      = (known after apply)

      + ports {
          + external = 9003
          + internal = 80
          + ip       = "0.0.0.0"
          + protocol = "tcp"
        }
    }

Plan: 3 to add, 0 to change, 0 to destroy.

------------------------------------------------------------------------

Note: You didn't specify an "-out" parameter to save this plan, so Terraform
can't guarantee that exactly these actions will be performed if
"terraform apply" is subsequently run.

## Terrform Apply Output

terraform apply

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # docker_container.wordpress[0] will be created
  + resource "docker_container" "wordpress" {
      + attach           = false
      + bridge           = (known after apply)
      + command          = (known after apply)
      + container_logs   = (known after apply)
      + dns              = (known after apply)
      + dns_opts         = (known after apply)
      + entrypoint       = (known after apply)
      + exit_code        = (known after apply)
      + gateway          = (known after apply)
      + hostname         = (known after apply)
      + id               = (known after apply)
      + image            = "wordpress:latest"
      + ip_address       = (known after apply)
      + ip_prefix_length = (known after apply)
      + ipc_mode         = (known after apply)
      + log_driver       = (known after apply)
      + log_opts         = (known after apply)
      + logs             = false
      + must_run         = true
      + name             = "wordpress-0"
      + network_data     = (known after apply)
      + read_only        = false
      + restart          = "always"
      + rm               = false
      + shm_size         = (known after apply)
      + start            = true
      + user             = (known after apply)
      + working_dir      = (known after apply)

      + ports {
          + external = 9001
          + internal = 80
          + ip       = "0.0.0.0"
          + protocol = "tcp"
        }
    }

  # docker_container.wordpress[1] will be created
  + resource "docker_container" "wordpress" {
      + attach           = false
      + bridge           = (known after apply)
      + command          = (known after apply)
      + container_logs   = (known after apply)
      + dns              = (known after apply)
      + dns_opts         = (known after apply)
      + entrypoint       = (known after apply)
      + exit_code        = (known after apply)
      + gateway          = (known after apply)
      + hostname         = (known after apply)
      + id               = (known after apply)
      + image            = "wordpress:latest"
      + ip_address       = (known after apply)
      + ip_prefix_length = (known after apply)
      + ipc_mode         = (known after apply)
      + log_driver       = (known after apply)
      + log_opts         = (known after apply)
      + logs             = false
      + must_run         = true
      + name             = "wordpress-1"
      + network_data     = (known after apply)
      + read_only        = false
      + restart          = "always"
      + rm               = false
      + shm_size         = (known after apply)
      + start            = true
      + user             = (known after apply)
      + working_dir      = (known after apply)

      + ports {
          + external = 9002
          + internal = 80
          + ip       = "0.0.0.0"
          + protocol = "tcp"
        }
    }

  # docker_container.wordpress[2] will be created
  + resource "docker_container" "wordpress" {
      + attach           = false
      + bridge           = (known after apply)
      + command          = (known after apply)
      + container_logs   = (known after apply)
      + dns              = (known after apply)
      + dns_opts         = (known after apply)
      + entrypoint       = (known after apply)
      + exit_code        = (known after apply)
      + gateway          = (known after apply)
      + hostname         = (known after apply)
      + id               = (known after apply)
      + image            = "wordpress:latest"
      + ip_address       = (known after apply)
      + ip_prefix_length = (known after apply)
      + ipc_mode         = (known after apply)
      + log_driver       = (known after apply)
      + log_opts         = (known after apply)
      + logs             = false
      + must_run         = true
      + name             = "wordpress-2"
      + network_data     = (known after apply)
      + read_only        = false
      + restart          = "always"
      + rm               = false
      + shm_size         = (known after apply)
      + start            = true
      + user             = (known after apply)
      + working_dir      = (known after apply)

      + ports {
          + external = 9003
          + internal = 80
          + ip       = "0.0.0.0"
          + protocol = "tcp"
        }
    }

Plan: 3 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

docker_container.wordpress[0]: Creating...
docker_container.wordpress[1]: Creating...
docker_container.wordpress[2]: Creating...
docker_container.wordpress[0]: Creation complete after 0s [id=76793943c3bfc18dd676f1f5311968165a2811ade3288046e1766a9c747c6ed2]
docker_container.wordpress[1]: Creation complete after 0s [id=26843be631979674b981f73b4b9744b96b9c1818f63b5519fcb7bec6d87f3823]
docker_container.wordpress[2]: Creation complete after 0s [id=37cb0f543a4df5a749663c42e93bee5cd1f47478bed3dd2c0f90cc16ca6275ea]

Apply complete! Resources: 3 added, 0 changed, 0 destroyed.

Outputs:

ip = 172.17.0.2, 172.17.0.3, 172.17.0.4