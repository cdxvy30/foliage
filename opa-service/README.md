# OPA Service
We utilize OPA service in both application side authorization and Terraform deployment workflow.

## Current Workflow
1. Define policies for both App side and DevOps side under `/policy` directory.
2. Push commit to trigger GitHub Actions, which will run:
   1. Bundle policies into `tar`.
   2. Build image for this bundle.
   3. Push image to Google Container Registry.
3. Execute on Cloud Run.

## Query API
- `/v1/data/terraform/authz/allow`: ask for deployment permission.
  - Request:
    ```
    {
      "input": {
        "request": {
            "user": "<email address>"
        }
      }
    }
    ```
- `/v1/data/hello/msg`