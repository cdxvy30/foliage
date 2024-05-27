package terraform.authz

import rego.v1

default allow := false

allowed_deployers := {"cdxvy@cathayholdings.com.tw"}

allow if {
  input.request.user in allowed_deployers
}