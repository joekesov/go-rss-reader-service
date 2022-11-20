# github.com/joekesov/go-rss-reader-service

## Install and run the RSS reader Service

1. From .env.example file create a new .env file and choose the port number for the REST API
```
// .env file
APP_PORT=3001
```

2. To build and run the REST Api service application

```bash
docker compose up -d --build
```

or if you want to run it in debug mode

- Build the container

```bash
docker compose build
```

- and run the container 

```bash
docker compose up
```

```
...
go-rss-reader-service-go_rss_service-1  | running...
go-rss-reader-service-go_rss_service-1  | [go-lib] initializing ...
go-rss-reader-service-go_rss_service-1  | INFO[0000] starting server                               addr=":3000"
```

## Use

```bash
curl -H 'Content-Type: application/json' -X POST \
    -d '{"urls":["https://blog.centos.org/feed","https://www.theregister.com/software/devops/headlines.atom"]}' \
    http://localhost:3001/rss
```

```json
[
    ...
  {
    "title": "GitHub puts prebuilt Codespaces into public beta",
    "source": "The Register - Software: Devops",
    "source_url": "https://www.theregister.com/software/devops/",
    "link": "https://go.theregister.com/feed/www.theregister.com/2022/02/25/github_codespaces_beta/",
    "publish-date": "2022-02-25T03:04:13Z",
    "DateValid": true,
    "description": "<h4>Say goodbye to your coffee break</h4> <p>Microsoft's GitHub on Wednesday said customers using its Codespaces hosted development environments can now try out prebuilt systems in a public beta test.…</p> <p><!--#include virtual='/data_centre/_whitepaper_textlinks_top.html' --></p>"
  },
    ...
]
```