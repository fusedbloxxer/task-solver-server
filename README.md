# task-solver-server

## Run the Server Locally

```bash
# Set the environment for the server
export TASK_SOLVER_ENV=dev

# To unset use
unset TASK_SOLVER_ENV

# Check the variable was set properly
env | grep TASK_SOLVER_ENV

# To compile the source files into a binary
go build -o ./bin/server.exe ./src 

# To run the server
go run ./src
```

---

## Run the Server in Cloud

### Build Image for Docker

```bash
# Build and save the image to a private Docker Repository
docker build --tag initial-version -f ./docker/Dockerfile . -t invokariman/task-solver-server

# View the local images
docker image ls

# TODO:
# Find a way to pass the Google Service Account Credentials to the running container
# Without compromising security

# Run the image as a container, in background, to test it
docker run -d --name task-solver-server --env TASK_SOLVER_ENV=dev -p 127.0.0.1:8080:8080 invokariman/task-solver-server

# Kill the container and remove it after you're done testing
docker rm $(docker stop $(docker ps -a -q --filter="name=task-solver-server"))
```