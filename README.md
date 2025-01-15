Simple Stickers API
============================

A simple REST API that returns a random sticker that i've got from internet and saved to db  

This app is build with [Go Fiber](https://gofiber.io/) as the framework

### Prerequisite :

- Docker
- OS Linux Docker Image,
- Golang version 1.19 or latest

# Install

## Install localy
    go mod tidy

#### Config

    Make sure to change `.env` variable value with your local db

#### Run the app

    go run main.go

#### Run the tests

    cd internal/service
    go test -verbose --cover

## Build with docker

_**note:**_ The docker-compose file contain mysql

    docker-compose -f deployment/docker-compose.yaml up -d

Folder Structure
============================

### Top-level directory layout

    .
    ├── bootstrap                 # Connector with other depedency for example mysql, redis, kafka etc
    ├── config                    # Configuration folder that used by apps (.env like)
    ├── deployment                # Dockerfile and docker-compose.yaml
    ├── internal                  # Main directory for rest api
    ├── migrations                # Migrations for database
    ├── go.mod                  
    ├── main.go
    └── README.md

### Internal folder

The actual source files of this project are stored inside the
`internal` directory. this contains `api`,`entities`,`helper`, `middleware`, `repository`, `service` and `pkg`

    .
    ├── internal
    │   ├── api                 # Contains all api this
    │   │   ├── dto             # Used for processing the struct before send it as response
    │   │   ├── handlers        # Files that handling the requests and the response of endpoint
    │   │   ├── presenter       # Struct that used as response of endpoint
    │   │   ├── request_model   # Struct for request parameter/body
    │   │   └── router.go       # Routes of all endpoints
    │   ├── const               # Constant variable list
    │   ├── entities            # Struct of data that represent from database
    │   ├── middlewares         # List of middlewares runs before handler
    │   ├── pkg                 # List of package we can use  
    │   ├── repositories        # Database interactor
    │   └── service             # Business Logic of the endpoints
    └── ...