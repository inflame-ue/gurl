# gurl

A RESTful URL shortening service built in Go.

Built following the [URL Shortening Service](https://roadmap.sh/projects/url-shortening-service) project spec from roadmap.sh.

## Endpoints

| Method | Path                           | Description              |
| ------ | ------------------------------ | ------------------------ |
| POST   | `/shorten`                     | Create a short URL       |
| GET    | `/shorten/{shortCode}`         | Retrieve original URL    |
| PUT    | `/shorten/{shortCode}`         | Update a short URL       |
| DELETE | `/shorten/{shortCode}`         | Delete a short URL       |
| GET    | `/shorten/{shortCode}/stats`   | Get URL access statistics|

## Tech Stack

- **Language:** Go (standard library `net/http` with Go 1.22+ mux)
- **Database:** SQLite via `modernc.org/sqlite`
- **Migrations:** goose

## Running

```bash
cp .env.example .env
go run ./...
```
