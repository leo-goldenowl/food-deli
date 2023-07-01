# Set up

### Prerequisites

1. Install [Go](https://golang.org/doc/install)
2. Install [Docker](https://docs.docker.com/get-docker/)
3. Install [Docker Compose](https://docs.docker.com/compose/install/)
4. Make a copy of `.env.example` and rename it to `.env` and fill in the values

```sh
cp .env.example .env
```

5. Make a copy of `.docker.env.example` and rename it to `.docker.env` and fill in the values

```sh
cp .docker.env.example .docker.env
```

6. Run the docker compose

```sh
docker-compose up
```

7. Mirgrate the database

```sh
go run mirgration/mirgration.go
```

8. Install [Air - live reload fo Go apps](https://github.com/cosmtrek/air)

### Run the app

1. Install the dependencies

```sh
go mod download && go mod tidy
```

2. Run the app

- if need to run the app with live reload

```sh
air
```

- if not need to run the app without live reload

```sh
go run main.go
```

### Deployment
1. Request access to join [render.com](https://dashboard.render.com/)
2. Follow the doucmentation [here](https://render.com/docs/deploy-go-gin) to deploy
3. Follow this documentation [here](https://render.com/docs/configure-environment-variables) to set all value of .env into the app

### Database schema
[![](https://mermaid.ink/img/pako:eNqdU2FrwjAQ_Sshn_UP9JuwDoZjDKvfhHJrTpetTeRyUUbtf99pZVrjHCwpNH3v5fLueml15Q3qTCM9WFgTNEunZCyKfFaoy3XbfxxGjNYoeV6nZywwWbdWDhpMQGzA1gm6gRB2nkxCBKg5AWsIXN6MvrKUUNax0MAxnLGP4N2bgi0w0Bll26CqCIHRlMBXRNyYAdH1r1lezCeL2eRlXiTAX4W6Ze1e8cAYwpDkUfu1v8Yqv0UanlTbT7Ff-ej4vzn_JOJ3Dqm05pdSlM9P0zxNX8wzRALHslU9Tq_oGI4xB8RNf_15fS_u9-OxbwdVz9RSy7yjOfm7EF7uvy_XI90gSRsbuSrHDJea31F-lz4IDK4gStNK2E6kENkXX67SGVPEke4rerpfOltBHbD7BlgrCZ4?type=png)](https://mermaid.live/edit#pako:eNqdU2FrwjAQ_Sshn_UP9JuwDoZjDKvfhHJrTpetTeRyUUbtf99pZVrjHCwpNH3v5fLueml15Q3qTCM9WFgTNEunZCyKfFaoy3XbfxxGjNYoeV6nZywwWbdWDhpMQGzA1gm6gRB2nkxCBKg5AWsIXN6MvrKUUNax0MAxnLGP4N2bgi0w0Bll26CqCIHRlMBXRNyYAdH1r1lezCeL2eRlXiTAX4W6Ze1e8cAYwpDkUfu1v8Yqv0UanlTbT7Ff-ej4vzn_JOJ3Dqm05pdSlM9P0zxNX8wzRALHslU9Tq_oGI4xB8RNf_15fS_u9-OxbwdVz9RSy7yjOfm7EF7uvy_XI90gSRsbuSrHDJea31F-lz4IDK4gStNK2E6kENkXX67SGVPEke4rerpfOltBHbD7BlgrCZ4)