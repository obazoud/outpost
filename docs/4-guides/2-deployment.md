# Outpost Deployment

The service will be distributed as a Binary and Docker image. The executable will have 3 entry points `delivery`, `api`, and `log`. You can omit the entry point to start all 3 services within the same process. 

Since each service has a very different usage and performance profile, it’s recommended that each service be operated and scaled independently.

We will provide a helm chart deployable on Kubernetes; since the services need to pull from your message bus, you need to execute it on a running process.

Resource (CPU & RAM) allocation recommendations will also be provided for each service.