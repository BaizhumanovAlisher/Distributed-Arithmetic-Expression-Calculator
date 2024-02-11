# Distributed-Arithmetic-Expression-Calculator

Project from Yandex Lyceum Course - Programming in Go

## Tech specification:
- [Russian](docs/technical%20specification-RU.md)

## Front-end -- SPA

Recommendation: 

Do not use GUI, (in the future), but you are able

Use API in [file](api/api.yaml) 

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
- It is lived 60s. It is described in [config.yaml](orchestrator/config.yaml)
- If header is null, token will not be used
- No cache in `500` http code