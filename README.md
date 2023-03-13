# PubSub subscription example
...

## Installation
1. Install dependencies:

```bash
go install
```

2. Copy .env.example file

```bash
cp .env.example .env
```

3. Run

```bash
go run main.go
```

## Run with Docker
1. Build an image:
```bash
docker build -t golangpubsub:alpine . --no-cache
```

2. Run an image as a utility container:
```bash
docker run --rm -p 80:8080 -v $(pwd):/app golangpubsub:alpine run main.go
```
