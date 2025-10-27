# Go HTTP Server with PostgreSQL and Docker Compose

This project provides a Go-based HTTP server that connects to a PostgreSQL database. The database is containerized using Docker Compose, and the Go server is run locally.

## Prerequisites

- Docker
- Docker Compose
- Go (version 1.21 or newer)

## Getting Started

1.  **Create a `.env` file**

    Create a `.env` file in the same directory as the `docker-compose.yaml` file and add the following content:

    ```env
    POSTGRES_USER=admin
    POSTGRES_PASSWORD=admin
    POSTGRES_DB=app
    SERVER_ADDR=0.0.0.0:8443
    ```

2.  **Start the Database**

    Start the PostgreSQL database using Docker Compose:

    ```bash
    docker-compose up -d
    ```

    This command runs the database container in detached mode.

3.  **Run the Go Server**

    In a separate terminal, run the Go server:

    ```bash
    go run ./cmd/api
    ```

    The server will connect to the PostgreSQL database running in Docker.

## Verifying the Server

You can check the logs from the `go run` command to ensure everything started correctly. You should see output indicating that the database connection was successful, migrations were run, and the server has started on port 8443.

To test the API, you can make a request to one of its endpoints with `curl`:

```bash
curl http://localhost:8443/api/blocked_sign/qry
```

## Stopping the Application

1.  **Stop the Go Server**

    Press `Ctrl+C` in the terminal where the server is running.

2.  **Stop the Database**

    To stop the database container, run:

    ```bash
    docker-compose down
    ```

To stop and remove the volume (deleting all data), run:

```bash
docker-compose down -v
```
