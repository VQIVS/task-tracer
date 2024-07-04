# task tracker

task tracker is a simple application designed to help users manage their tasks. it provides both a web api and a cli for creating, listing, and managing tasks.

## table of contents

- [overview](#overview)
- [features](#features)
- [project structure](#project-structure)
- [setup and installation](#setup-and-installation)
- [configuration](#configuration)
- [running the server](#running-the-server)
- [using the cli](#using-the-cli)
- [api endpoints](#api-endpoints)
- [contributing](#contributing)
- [license](#license)

## overview

the task tracker application allows users to create, list, and manage tasks. the application is built with golang, using gorm for database interactions, cobra for cli, and gorilla mux for routing.

## features

- create tasks
- list all tasks
- mark tasks as completed
- command line interface (cli) for managing tasks
- restful api for interacting with tasks

## project structure

\```
.
├── api
│   ├── handlers
│   ├── middlewares
│   └── setup.go
├── cmd
│   ├── cli
│   │   └── main.go
│   └── server
│       └── main.go
├── config
│   ├── config.go
│   ├── config.yaml
│   └── read.go
├── go.mod
├── go.sum
├── internal
│   ├── task
│   │   ├── model.go
│   │   ├── repo.go
│   │   └── service.go
├── pkg
│   ├── adapters
│   │   └── storage
│   ├── database
│   │   └── db.go
│   └── jwt
└── service
\```

## setup and installation

### prerequisites

- golang 1.16 or later
- sqlite (or another supported database)

### installation

1. clone the repository:

\```bash
git clone https://github.com/your-username/task-tracker.git
cd task-tracker
\```

2. install dependencies:

\```bash
go mod tidy
\```

## configuration

1. create a `config.yaml` file in the `config` directory with the following content:

\```yaml
database_url: "task_tracker.db"
server_port: "8080"
jwt_secret: "your_secret_key"
\```

2. make sure the configuration file matches the structure in `config/config.go`.

## running the server

to start the server, run the following command:

\```bash
go build -o bin/server cmd/server/main.go
./bin/server
\```

the server should start on the port specified in the configuration file (e.g., `http://localhost:8080`).

## using the cli

### building the cli

to build the cli tool, run the following command:

\```bash
go build -o bin/cli cmd/cli/main.go
\```

### using the cli

- list all tasks:

\```bash
./bin/cli list
\```

- create a new task:

\```bash
./bin/cli create --title "new task" --description "details of the new task"
\```

## api endpoints

### list all tasks

**get** `/tasks`

_response_:

\```json
[
    {
        "id": 1,
        "title": "sample task",
        "description": "this is a sample task",
        "completed": false
    }
]
\```

### create a new task

**post** `/task`

_request_:

\```json
{
    "title": "new task",
    "description": "details of the new task"
}
\```

_response_:

\```json
{
    "id": 2,
    "title": "new task",
    "description": "details of the new task",
    "completed": false
}
\```

## contributing

contributions are welcome! please fork the repository and submit a pull request.

1. fork the repository.
2. create a new branch (`git checkout -b feature-branch`).
3. make your changes.
4. commit your changes (`git commit -m 'add some feature'`).
5. push to the branch (`git push origin feature-branch`).
6. open a pull request.

## license

this project is licensed under the mit license. see the [license](license) file for details.
