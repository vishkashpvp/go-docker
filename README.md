# Go-Docker

Go-Docker is a simple and minimal template designed to help you get started with Dockerizing your [Go](https://go.dev/) projects.

## 1. Overview

The Go-Docker template provides a basic structure and guidance to assist you in Dockerizing your Go projects, making it easier to package and distribute your applications as containers.

- Provides a starting point for Dockerizing your Go projects.
- Simplifies Docker image creation steps for Go applications.

Dockerfile gives you a good understanding of how docker images and containers are built.

## 2. Getting Started

Follow these steps to start using the Go-Docker template for your Go applications.

### 2.1 Prerequisites

Before using the Go-Docker template, ensure you have the following prerequisites installed:

- [Docker](http://docker.com) running locally
- [Go](https://go.dev) 1.19 or later
- [Git](https://git-scm.com)

### 2.2 Run Go-Docker

To run and test the Go-Docker template locally, you follow these steps:

#### 2.2.1 As Docker Container

Dockerfile has extra info regarding each step/layer that is happening like how it is being built or run

1. Clone this repository to your local machine:

   ```shell
   $ git clone https://github.com/viskashpvp/go-docker.git
   ```

2. Change into the project directory:

   ```shell
   $ cd go-docker
   ```

3. For dev environment, duplicate `.env.example` file and rename copy file to `.env` and add values/secrets accordingly

4. Build Docker image:

   ```shell
   $ docker build -t go-docker:v1.0-test .
   ```

   `-t` is short for `--tag` flag, is used to label the image, if omitted `latest` is used as default value by Docker

   check if image is built, using `docker images` command, you should be seeing something similar to this

   ```shell
   $ docker images
   REPOSITORY           TAG         IMAGE ID       CREATED         SIZE
   go-docker            v1.0-test   da8c488c48e1   7 seconds ago   1.16GB
   ```

5. Run Go-Docker as a Docker container:

   ```
   $ docker run -d -p 8080:8080 -e APP_NAME=go-docker-server -e APP_VERSION=v1.0 --name go-docker-container-1 go-docker:v1.0-test
   ```

   `-d` is short for `--detach`, will “detach” from the container and return you to the terminal prompt <br>
   `-p` is short for `--publish`, is used to publish a port pf container, follows `[host_port]:[container_port]` format <br>
   `-e` is short for `--env`, is used to set environment variables <br>
   `--name` flag is used to create a name for the container

   verify if container has started, using `docker ps` command, you should be seeing something similar to this

   ```shell
   $ docker ps
   CONTAINER ID     IMAGE                   STATUS          NAMES
   37cdb6273493     go-docker:v1.0-test     Up 7 seconds    go-docker-container-1
   ```

   You can also able to view other fields like `COMMAND`, `CREATED`, `PORTS`

6. Test if you are able to access container via exposed ports on localhost:

   ```shell
   $ curl localhost:8080/
   Welcome to the server
   ```

It's running successfully as a docker container.

#### 2.2.2 Without Docker

Follow first three steps as mentioned in [2.2.1](#221-as-docker-container)

4.  a. Build and Run Go application:

    ```shell
    $ go build -o main .
    $ ./main
    ```

    b. Run Go application without building:

    ```shell
    $ go run main.go
    ```

    NOTE: For production purposes, you can use Docker secret management or Github actions secrets to provide environment variables that are required by the application.

    ```shell
    $ GIN_MODE=release go run main.go
    ```

5.  Test if you are able to access server on localhost:

    ```shell
    $ curl localhost:8080/
    Welcome to the server
    ```

It's running successfully without docker as well.

## 3. Contributing

We welcome contributions to the Go-Docker template from the community. To contribute, please follow these guidelines:

- Fork the Go-Docker repository.
- Create a new branch for your feature or improvement: `git checkout -b feature-name`.
- Make your changes and commit them: `git commit -m "description of changes"`.
- Push your changes to your fork: `git push origin feature-name`.
- Create a pull request to the main Go-Docker repository.

Please ensure that your contributions include relevant documentation.

## 4. Feedback

We value your feedback! <br>
If you find any issues with the code, language, grammar, or have suggestions for improvements, please let us know. <br>
Your contributions and feedback help make this template better.
