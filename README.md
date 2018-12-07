# NI Les cools - Back-end

Backend written in Go (v1.11)

Technologies
- Gin/JWT/Postgres

Deployments:
- Docker
- Docker-compose

Executable: ./cmd/main.go + ./cmd/config.go
Flags: -config CONFIG_PATH // Default config.yaml

NOTE: THE API WILL CREATE FAKE DATA AT STARTUP IF TABLES ARE EMPTY OR DON'T EXIST

# API

###### User

- /user/auth **POST**
- /user/signup **POST**
- /user/pos **PUT** **GET**
- /user/refresh_token **POST**
- /food **GET**
- /food/create **POST**
- /alert **GET**
- /alert/create **POST**
- /mission **GET**
- /mission/create **POST**
- /equipment **GET**
- /equipment **POST**
- /weather **GET**
- /weather/create **POST**