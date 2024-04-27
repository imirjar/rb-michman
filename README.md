# Michman App

Hi! My name is Michman and I can manage all of your databases in a single API. I consist of three parts:

>**Diver** is a database agent that must be placed in a database local network or in a safe access point.

>**Michman** is a single API that can manage all of existing divers.

>**AUTH** is an auth app which can generate JWT

# INSTRUCTION

First of all you must input ./.env file in the root directory of the project.

## Michman
1) Input those variables from ./.env file: \
```
MICHMAN_ADDR --> address where app will start\
COMPRESSION  --> if exists we can compress our responses
SECRET       --> secret for JWT token creating\
```
2) **Run app** with command
```
go run cmd/michman/main.go
```
4) Run



## Diver

If you need you could run test db docker container
```
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=praktikum -d postgres:latest
```

1) Input those variables from ./.env file: \
```
DIVER_ADDR     --> address where app will start\
COMPRESSION    --> if exists we can compress our responses
SECRET         --> secret for JWT token creating\
TARGET_DB_CONN --> connection to PSQL db where we will sent requests
```
2) **Run app** with command
```
go run cmd/diver/main.go
```
4) Run

## Auth
1) Input those variables from ./.env file: \
```
AUTH_ADDR --> address where app will start\
TOKEN_EXP --> live time of JWT token\
SECRET    --> secret for JWT token creating\
```

2) **Run app** with command
```
go run cmd/auth/main.go
```
