# PostgreSQL with Docker Compose

This setup provides a simple way to run a PostgreSQL database using Docker Compose.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

1.  **Create a `.env` file**

    Create a `.env` file in the same directory as the `docker-compose.yaml` file and add the following content:

    ```
    POSTGRES_USER=admin
    POSTGRES_PASSWORD=admin
    POSTGRES_DB=app
    ```

2.  **Start the database**

    Run the following command to start the PostgreSQL container in detached mode:

    ```bash
    docker-compose up -d
    ```

3.  **Connect to the database**

    You can connect to the database using any PostgreSQL client with the following credentials:

    -   **Host**: `localhost`
    -   **Port**: `5432`
    -   **User**: `admin` (or the value of `POSTGRES_USER` in your `.env` file)
    -   **Password**: `admin` (or the value of `POSTGRES_PASSWORD` in your `.env` file)
    -   **Database**: `app` (or the value of `POSTGRES_DB` in your `.env` file)

4.  **Stop the database**

    To stop the container, run:

    ```bash
    docker-compose down
    ```

    To stop and remove the volume (deleting all data), run:

    ```bash
    docker-compose down -v
    ```
