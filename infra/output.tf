output "receiver_ip" {
  description = "External IP address for receiver. You should register this address to your Discord bot."
  value       = google_compute_instance.instance_receiver.network_interface.0.access_config.0.nat_ip
}

output "start_endpoint" {
  description = "HTTP endpoint to start the target directly"
  value       = module.cf_start.https_endpoint
}

output "stop_endpoint" {
  description = "HTTP endpoint to stop the target directly"
  value       = module.cf_stop.https_endpoint
}
