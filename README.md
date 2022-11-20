docker compose up -d --build

docker compose build

docker compose up

docker compose down


docker compose run --rm app go mod tidy

docker compose exec server bash


go get github.com/joekesov/go-rss-reader-package

docker compose run --rm go_rss_service go mod init github.com/joekesov/go-rss-reader-service

docker compose run --rm go_rss_service air init



docker compose run --rm go_rss_service go mod tidy

docker compose run --rm app go mod tidy

go get github.com/joekesov/go-rss-reader-package@v0.1.0