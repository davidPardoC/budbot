# BudBot

## Setting up development environment

### Prerequisites:

- Go 1.22.x
- Docker (for local containers)

1. Create env file

```bash
cp .env.example .env
```

2. Raise local development containers with docker compose.

```bash
docker compose -f docker-compose.local.yml up
```

3. Install dependencies

```bash
go mod tidy
```

4. Run local server

```bash
go run main.go
```

Or you can use [air](https://github.com/air-verse/air) for Live reload

```bash
air
```

You can use [ngrok](https://ngrok.com/) for local development in order to receive webhooks.
