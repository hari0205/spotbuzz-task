# Spotbuzz Task

Steps to Run this project

- Clone the repo
- Install all the necessary modules
```
 go mod download
```

- Run the file
```
go run main.go
```

- **OPTIONAL**: You can use the existing Database credentials or use docker to run local instances for Databases 
```
 docker run -d --name postgres-container -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -p 5432:5432 postgres:latest
```
- The above command will start a POSTGRES DB in detached mode(Background)
- 
### Steps to Run this project with docker

- Clone the repo

- Run the following command


```
    docker-compose up --build -d
```
- `--build` : To build the docker compose . 
- `-d` : To run the app container in detached mode.
