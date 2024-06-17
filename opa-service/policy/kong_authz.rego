package kong.authz

import rego.v1

default allow := false

allow if {
  input.jwt.payload.role == "admin"
}