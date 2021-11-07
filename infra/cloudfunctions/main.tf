# Public cloud function template

data "archive_file" "archive_cf" {
  type        = "zip"
  source_dir  = "../functions/${var.function_name}"
  output_path = "../functions/${var.function_name}.zip"
}

resource "google_storage_bucket_object" "object_source_cf" {
  name   = "${var.function_name}.${data.archive_file.archive_cf.output_md5}.zip"
  bucket = var.source_bucket_name
  source = data.archive_file.archive_cf.output_path
}

resource "google_cloudfunctions_function" "function_cf" {
  name                  = "server-kun-${var.function_name}"
  description           = var.description
  region                = var.function_region
  runtime               = "go113"
  available_memory_mb   = 128
  timeout               = 30
  source_archive_bucket = google_storage_bucket_object.object_source_cf.bucket
  source_archive_object = google_storage_bucket_object.object_source_cf.name
  trigger_http          = true
  entry_point           = var.entry_point

  environment_variables = {
    SERVER_PROJECT  = var.project
    SERVER_ZONE     = var.target_zone
    SERVER_INSTANCE = var.target_instance_name
  }
}

resource "google_cloudfunctions_function_iam_member" "invoker_cf" {
  region         = google_cloudfunctions_function.function_cf.region
  cloud_function = google_cloudfunctions_function.function_cf.name

  role   = "roles/cloudfunctions.invoker"
  member = "allUsers" # TODO: Limit invoker to the receiver instance
}
