up:
	docker compose --env-file .env.compose up -d

down:
	docker compose --env-file .env.compose down

fresh:
	docker compose --env-file .env.compose down --remove-orphans
	docker compose --env-file .env.compose build --no-cache
	docker compose --env-file .env.compose up -d --build -V

logs:
	docker compose logs -f

build:
	docker build --platform linux/amd64 -t $(name) -f Dockerfile .
