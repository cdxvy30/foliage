package role_check

import rego.v1

default allow_bronze := false
default allow_silver := false
default allow_gold := false

# secret key "B41BD5F462719C6D6118E673A2389"
secret := opa.runtime().env["MY_SECRET"]

allow_bronze if {
	is_get
	some role in claims.roles
    role == "bronze" #可改
}

allow_silver if {
	is_get
	some role in claims.roles
    role == "silver" #可改
}

allow_gold if {
	is_get
	some role in claims.roles
    role == "gold" #可改
}

is_get if input.method == "GET"

is_reader if {
	some role in claims.roles
    role == input.service
}

claims := payload if {
	io.jwt.verify_hs256(bearer_token, "B41BD5F462719C6D6118E673A2389")
	[_, payload, _] := io.jwt.decode(bearer_token)
}

bearer_token := t if {
	v := input.token
	startswith(v, "Bearer ")
	t := substring(v, count("Bearer "), -1)
}

# input sample
# {
#     "method": "GET",
#     "token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQWxpY2lhIFNtaXRoc29uaWFuIiwicm9sZXMiOlsicmVhZGVyIiwid3JpdGVyIl0sInVzZXJuYW1lIjoiYWxpY2UifQ.md2KPJFH9OgBq-N0RonGdf5doGYRO_1miN8ugTSeTYc",
#     "service": "reader"
# }

