# resource "google_service_account" "cloud_run_service_account" {
#   account_id   = "cloud-run-sa"
#   display_name = "Cloud Run Service Account"
# }

# resource "google_project_iam_member" "firestore_access" {
#   project = var.GCP_PROJECT
#   role    = "roles/datastore.user"
#   member  = "serviceAccount:${google_service_account.cloud_run_service_account.email}"
# }

resource "google_cloud_run_v2_job" "job" {
  name     = var.container_name
  location = var.GCP_REGION
  template {
    template {
      containers {
        image = "docker.io/cdxvy30/twse-service-amd-rate1:latest"
        resources {
          limits = {
            cpu    = "1"
            memory = "1024Mi"
          }
        }
      }
      service_account = var.service_account
    }
  }
}

resource "google_cloud_run_v2_job_iam_member" "invoker" {
  name     = google_cloud_run_v2_job.job.name
  location = google_cloud_run_v2_job.job.location
  role     = "roles/run.invoker"
  member   = "serviceAccount:${var.service_account}"
}

resource "google_cloud_run_v2_job_iam_member" "admin" {
  name     = google_cloud_run_v2_job.job.name
  location = google_cloud_run_v2_job.job.location
  role     = "roles/run.admin"
  member   = "serviceAccount:${var.service_account}"
}

# Allow all users to invoke the service (for testing purposes)
# resource "google_cloud_run_service_iam_member" "public_access" {
#   location = google_cloud_run_service.cloud_run_service.location
#   project  = google_cloud_run_service.cloud_run_service.project
#   service  = google_cloud_run_service.cloud_run_service.name

#   role   = "roles/run.invoker"
#   member = "allUsers"
# }

# output "cloud_run_url" {
#   value = google_cloud_run_service.cloud_run_service.status[0].url
# }