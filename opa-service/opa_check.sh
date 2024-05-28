#!/bin/bash

CURRENT_USER_EMAIL=$(gcloud config get-value account)

INPUT=$(cat <<EOF
{
  "request": {
    "user": "$CURRENT_USER_EMAIL"
  }
}
EOF
)

opa eval -i <(echo "$INPUT") -d ../../opa-service/deploy.rego "data.terraform.authz.allow"