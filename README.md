# Distributed-Arithmetic-Expression-Calculator
D-A-E-C

Project from Yandex Lyceum Course - Programming in Go

Start point: http://localhost:8099/
## Run project: 
0) install [docker engine](https://docs.docker.com/engine/install/) and [docker compose](https://docs.docker.com/compose/install/)
1) `cd <"path/to/project">`
2) `docker compose -f docker-compose.yml -p distributedarithmeticexpressioncalculator up -d`

## You can get error like: 

`Error response from daemon: Ports are not available: exposing port TCP 0.0.0.0:5673 -> 0.0.0.0:0: listen tcp 0.0.0.0:5673: bind: address already in use`

In this case you should change to free port in [docker-compose](docker-compose.yml)

Change only exposed ports

Example:
- From:
```yaml
  postgres:
    container_name: daec_postgresql
    image: postgres
    ports:
      - "5432:5432"
```
- To:
```yaml
  postgres:
    container_name: daec_postgresql
    image: postgres
    ports:
      - "8100:5432"
```

## Tech specification:
- [specifications](docs/technical%20specification-RU.md)
- [criteria](docs/evaluation%20criteria.md)


## No Front-end

Use OpenAPI in [file](api/api.yaml)

Later will be Postman file

# Rules for expression

1) expression should start with number
2) expression should not contain extra symbols
3) digit should not start with zero
4) expression should contain operations
5) it is forbidden to divide by zero
6) brackets should be correct
7) it is forbidden to use allowed chars nearby

# Token idempotency
- Header: `X-Idempotency-Token`
- It is used in `/expression` POST HTTP method
- It consists of token from user, separator and expression: "<"user token idempotency">__<"expression">". Example: `dkskdhen392h__2+3*4`
- It is lived 360s. It is described in [config.yaml](config.yaml)
- **If header is null, token will not be used**
- No caching in `500` http code
