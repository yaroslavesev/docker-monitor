# Docker Monitor (Go + Angular + PostgreSQL + Nginx)

This repository contains a **container-monitoring application** built with:
- **Go** (Golang) for the Backend and the Pinger service
- **Angular** (TypeScript) for the Frontend
- **PostgreSQL** as the database
- **Nginx** acting as a reverse proxy (optional, but recommended)

The application’s purpose is to:
1. **Periodically ping Docker containers** (or any IP addresses)
2. Store the ping results in a **PostgreSQL** database
3. Provide a **REST API** (Go + Gin) to manage container records
4. Allow users to see the ping statuses on a **dynamic web page** (Angular)

Below you’ll find:

- [Project Architecture](#project-architecture)
- [Services Overview](#services-overview)
- [Folder Structures](#folder-structures)
- [Endpoints (Postman Collection)](#endpoints-postman-collection)
- [How to Run (Docker Compose)](#how-to-run-docker-compose)
- [Nginx Setup](#nginx-setup)
- [Technologies Used](#technologies-used)
- [Future Plans](#future-plans)

---

## Project Architecture

We have **three main services** plus one optional proxy:

1. **Backend** (Go + Gin):
    - Exposes REST endpoints: create/read/update/delete containers (`/api/containers`).
    - Handles database connections (PostgreSQL).
    - Provides a repository layer (`/internal/repositories`) for CRUD operations.

2. **Pinger** (Go):
    - Periodically fetches the list of IP addresses (Docker containers) from the Backend.
    - Executes `ping` commands or uses a library to do ICMP checks.
    - Sends back the ping results (timestamps) to the Backend.

3. **Frontend** (Angular):
    - Displays a table of containers/IP addresses, last ping time, last success time.
    - Allows you to add or remove IP addresses through the Backend API.
    - Automatically refreshes every X seconds so you can see updates without manual reload.

4. **Nginx** (optional reverse proxy):
    - Listens on port 80.
    - Routes `/api/` requests to the Backend.
    - Routes `/` to the Frontend (or the Angular dev server).
    - Resolves any CORS concerns by serving everything under one domain (if needed).

> The **PostgreSQL** service is also defined in `docker-compose.yml`, storing all container and ping data.

---

## Services Overview

### 1. Backend

- Uses [Gin](https://github.com/gin-gonic/gin) as the HTTP framework.
- Talks to PostgreSQL via GORM or standard database library (depending on your preference).
- Defines REST routes for container management:
    - `POST /api/containers`
    - `GET /api/containers`
    - `PUT /api/containers/:id`
    - `DELETE /api/containers/:id`

### 2. Pinger

- Written in Go.
- On a schedule (e.g. every 10 seconds), it:
    1. Requests all containers from the Backend (GET `/api/containers`).
    2. Pings each container’s IP.
    3. Sends updates back (PUT `/api/containers/:id`) with last ping time and last success time if ping was successful.

### 3. Frontend

- Built with **Angular** and TypeScript.
- Provides a simple UI (table) by **Bootstrap** to list all containers (IP addresses) and show their ping status.
- You can add a new IP address, remove an existing one, or see updates in near real-time.
- **Refresh** or **auto-refresh** the list every X seconds.

### 4. Nginx (optional)

- If you choose to run Nginx:
    - You can hide the ports of the backend and the frontend behind `http://localhost:80/`.
    - Typically routes `/api/` → `backend:8080/api/` and `/` → `frontend:4200/` (or the built Angular app on port 80 inside a container).

---

## Folder Structures

Below are **example** structures for each sub-project:

### Backend

```
backend/
├── Dockerfile
├── go.mod
├── go.sum
├── internal/
│   ├── app/
│   │   └── app.go               // start the backend-engine!
│   ├── config/
│   │   └── config.go            // environment/config loading
│   ├── controllers/            // CRUD controllers
│   │   ├── createContainer.go
│   │   ├── deleteContainer.go
│   │   ├── readAllContainers.go
│   │   └── updateContainer.go
│   ├── db/
│   │   ├── db.go                // DB connection
│   │   └── migrations.go        // AutoMigrate or raw SQL migrations
│   ├── models/
│   │   └── container.go
│   ├── repositories/
│   │   └── containerRepository.go // CRUD operations in DB
│   └── server/
│       └── routes.go            // Register Gin routes (/api/...)
└── main.go                      // Entrypoint 
```

### Pinger

```
pinger/
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── internal/
│   ├── app/
│   │   └── app.go         // starts the pinger-engine!
│   ├── config/
│   │   └── config.go      // env loading: BACKEND_URL, PING_INTERVAL_SECONDS, etc.
│   ├── models/
│   │   └── container.go   // mirror the Container struct from the backend
│   │   └── config.go      // Config struct
│   └── service/ //main logic to ping IPs, call backend
│       └── containersAPI.go
│       └── pingerLoop.go
│       └── pingExecutor.go
│       └── pingProcess.go  
└── .env                    // optional environment variables
```

### Frontend (Angular)

```
frontend/
├── Dockerfile
├── package.json
├── angular.json
├── tsconfig.json
├── src/
│   ├── index.html
│   ├── main.ts
│   ├── styles.css
│   ├── app/
│   │   ├── models/ 
│   │   │   └── container.ts // Container struct
│   │   ├── pages/
│   │   │   └── containers/ // Container component
│   │   │       ├── containers.component.html
│   │   │       ├── containers.component.css
│   │   │       └── containers.component.ts
│   │   ├── services/ // HTTP-handler service
│   │   │   └── api.service.ts
│   │   ├── app.component.ts
│   │   ├── app.routes.ts
│   │   └── app.config.ts
│   └── environments/ // env info
│       ├── environment.development.ts
│       └── environment.ts
```

> In a production build, Angular can be served as static files (e.g. via Nginx). In dev mode, you can use the Angular dev server.

---

## Endpoints (Postman Collection)

A sample Postman collection is available here:

> [**Postman Docs**](https://documenter.getpostman.com/view/37431052/2sAYX6oMJE)

It includes endpoints:
- `GET /api/containers`
- `POST /api/containers`
- `PUT /api/containers/:id`
- `DELETE /api/containers/:id`

Use it to quickly test your backend.

---

## How to Run (Docker Compose)

The quickest way is:

1. **Install Docker** and **Docker Compose**.
2. **Clone** this repository.
3. **Adjust** environment variables if needed:
    - In the `backend/.env` or `docker-compose.yml`.
    - For pinger, set `BACKEND_URL` to `http://backend:8080`.
4. From the **root directory** (where `docker-compose.yml` is located), run:
   ```bash
   docker-compose up --build
   ```
5. Wait until all containers spin up:
    - `db` (PostgreSQL)
    - `backend`
    - `pinger`
    - `frontend`
    - `nginx` (if included)

Depending on your setup:

- If you are **using Nginx** as a reverse proxy, open [http://localhost](http://localhost) and see your Angular app. All requests to `/api` go to the backend.
- If you **exposed** the frontend on port 4200, open [http://localhost:4200](http://localhost:4200). Then your Angular dev server is accessible. You may need to handle CORS or reference the backend differently.

### Note on local development

- If you want **live code reloading** for the Angular front or the Go backend, you may run them outside of Docker (on your host machine) and still rely on the DB in Docker.
- Or you can keep everything inside Docker for a consistent environment.

---

## Nginx Setup

There is an **optional** Nginx service defined in `docker-compose.yml` (and a config like `nginx/default.conf`) that:

```nginx
server {
    listen 80;

    location /api/ {
        proxy_pass http://backend:8080/api/;
    }

    location / {
        proxy_pass http://frontend:80/; 
        # or if using dev server: :4200
    }
}
```

So everything becomes accessible under **`http://localhost`**:

- The **frontend** is served from `/`.
- The **backend** endpoints are proxied at `/api/`.

> If you do this, you typically **do not** need special CORS headers in the Go backend. Everything is from the same domain (`localhost`).

---

## Technologies Used

1. **Go (Golang)**
    - [Gin Gonic](https://github.com/gin-gonic/gin) for HTTP server
    - Optionally [GORM](https://gorm.io/) for ORM or direct DB queries
    - For the **pinger** service, standard library or an external ping library.

2. **Angular** (TypeScript, CLI, etc.)
    - Provides a reactive front-end with a table of containers.

3. **PostgreSQL**
    - Official Docker image `postgres:14`.
    - Stores container IP addresses and ping results.

4. **Docker & Docker Compose**
    - All services are containerized via Dockerfiles.
    - A single `docker-compose.yml` orchestrates them.

5. **Nginx** (optional)
    - Reverse proxy for a single “point of entry” and simpler domain structure.

---

## Future Plans

Possible **advanced** or “next step” features:

- **netns** (Network Namespaces) to isolate ping behavior or test pings from different subnets/contexts.
- **Message Queue** (e.g., RabbitMQ, Kafka) to handle pinger results asynchronously and scale out more easily.
- **Authentication** / token-based security on the backend to control who can add or remove containers.
- **Extended monitoring**: store historical ping times, produce charts or logs of success/failure over time.

---

## Conclusion

With this setup, you have:

- **Backend**: RESTful API in Go, storing data in PostgreSQL.
- **Pinger**: Automated ping service (Go).
- **Frontend**: Angular application to display/edit container data.
- (Optional) **Nginx**: Reverse proxy for a unified domain.

Run `docker-compose up --build`, open your browser to [http://localhost](http://localhost), and watch the containers be pinged in real time!
