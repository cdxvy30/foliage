# foliage
This repository consists of backend service and infrastructure configuration for Foliage which is a real-time ledger information service.

## System Architecture
![image](https://github.com/user-attachments/assets/c92b8c93-313d-4aee-94dd-a75d7a3bd002)


## OPA
`/policy` contains the Rego defined policies to decide whether a service can access to specific resource and whether a service can be deployed to specific environment.
`/data` defines the data used in evaluation stage in JSON.
`Dockerfile` can help deploy OPA as a containerized service.

## Pub/Sub
Real-time ledger information is achieved by Pub/Sub mechanism of Redis and WebSocket. Refer to `/emitter-service` and `/subscriber-service`.

## Infra
`/terraform` contains IaC configuration for the serverless service we use (Google Cloud Run).
