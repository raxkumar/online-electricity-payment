# backend prototype

This is a go-micro prototype generated using WeDAA, you can find documentation and help at [WeDAA Docs](https://www.wedaa.tech/docs/introduction/what-is-wedaa/)

## Prerequisites

- go version >= 1.20

## Project Structure

```
├── auth/ (IdP configuration for keycloak)
├── config/ (configuration properties loader)
├── controllers/ (api controllers)
├── db/ (DB connection configuration)
├── docker/ (contains docker compose files for external components based on architecture design)
├── handler/ (DB handler methods)
├── migrate/ (database schema change management)
├── proto/ (proto files supporting DB models)
├── resources/ (configuration properties)
├── Dockerfile (for packaging the application as docker image)
├── README.md (Project documentation)
├── comm.yo-rc.json (generator configuration file for communications)
├── go.mod
└── main.go
```

## Dependencies

This application is configured to work with external component(s).

Docker compose files are provided for the same to get started quickly.

Component details:

- Keycloak as Identity Management: `docker compose -f docker/keycloak.yml up -d`
- Postgresql DB: `docker compose -f docker/postgresql.yml up -d`

On launch, backend will refuse to start if it is not able to connect to any of the above component(s).

## Get Started

Install required dependencies: `go mod tidy`

Run the prototype locally: `go run .`

Open [http://localhost:9020/hello](http://localhost:9020/hello) to view it in your browser.

The page will reload when you make changes.

## Containerization

Build the docker image: `docker build -t backend:latest .`

Start the container: `docker run -d -p 9020:9020 backend:latest`


## protoc cmd:-

protoc --proto_path=proto \
  --go_out=pb --go_opt=paths=source_relative \
  --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
  --micro_out=pb --micro_opt=paths=source_relative \
  proto/*.proto


## docker container

docker stop f61ad6ee319a dc23f2c2cee3 8d5aa9239eac

docker start f61ad6ee319a dc23f2c2cee3 8d5aa9239eac