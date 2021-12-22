# task-solver-server

## Run the Server Locally

## Configure the settings

In the **env** folder you can add multiple configurations and edit them to suit your needs.

By default the development environment (**dev**) is chosen.

```bash
# Install Swagger / OpenAPI dependencies
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Set the environment for the server
export TASK_SOLVER_ENV=dev

# To unset use
unset TASK_SOLVER_ENV

# Check the variable was set properly
env | grep TASK_SOLVER_ENV

# Generate swagger docs
swag init --dir ./src

# To compile the source files into a binary
go build -o ./bin/server ./src 

# To run the server
./bin/server
```

You should now be able to access the API using your custom settings, 
or by default at http://127.0.0.1:8080/swagger/index.html.

---

## Run the Server in Cloud

### 1. Build Image for Docker

When building you should provide a tag for the image to easily push it later to a repository:

```bash
# Build the image using the steps provided in the Dockerfile
docker build -t invokariman/task-solver-server:latest -f ./docker/Dockerfile . 
```

#### (Optional) Test the Server using Docker Locally

Run the container and visit http://127.0.0.1/api/v1/test to make sure it's working properly:

```bash
# Run the server in detached mode using dev config
docker run -d --name task-solver-server --env TASK_SOLVER_ENV=dev -p 127.0.0.1:8080:8080 invokariman/task-solver-server:latest

# List the locally available images
docker image ls

# Kill the container and remove it after you're done testing
docker rm $(docker stop $(docker ps -a -q --filter="name=task-solver-server"))
```

### 2. Publish the Image to a (Private) Repository

! **ATTENTION** ! The credentials are being copied in the image. ! **ATTENTION** ! 

When you publish to a public repository be sure to remove them!

- **Google**: https://cloud.google.com/container-registry/docs/advanced-authentication

- **Docker**: https://docs.docker.com/engine/reference/commandline/login/

Login to docker - If you are using 2FA, an [access token](https://docs.docker.com/docker-hub/access-tokens/)
will be necessary:

```bash
docker login
```

When you are ready to push the image to a [repository](https://docs.docker.com/docker-hub/repos/) run the following command:

```bash
# Publish the image to your repository
docker push invokariman/task-solver-server:latest
```

Visit your [repository page](https://hub.docker.com/) to make sure everything worked properly.

### 3. Run the Published Image as a Container in Azure

Login to [Azure](https://portal.azure.com/) using the command line:

```bash
docker login azure
```

Create a [resource group](https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/manage-resource-groups-portal)
and a [docker context](https://docs.docker.com/engine/context/working-with-contexts/) to be able to run the image in 
the cloud.

```bash
docker context create aci azure-cloud
```

Run the Server using the previous context:

```bash
docker --context azure-cloud run --name task-solver-server --env TASK_SOLVER_ENV=pro -p 8080:8080 invokariman/task-solver-server:latest
```

Go to the [Azure Portal](https://portal.azure.com/) and use the provided ip to test your application.

http://{AZURE_PROVIDED_IP}:{AZURE_PORT}/api/v1/test

Visit the [Docker Documentation](https://docs.docker.com/cloud/aci-integration/)
see more details regarding this process.

---

## API Documentation

## Swagger / OpenAPI

This project is using Swagger / OpenAPI to document its exported functionalities.

You can access the swagger page at http://127.0.0.1:8080/swagger/index.html for the **dev** environment.