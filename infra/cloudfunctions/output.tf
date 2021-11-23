output "https_endpoint" {
  description = "HTTP endpoint to start the target directly"
  value       = google_cloudfunctions_function.function_cf.https_trigger_url
}
