# CSV Data Ingestion
This project is designed to ingest data from a CSV file, send the data to a message queue (RabbitMQ), process the data in parallel using consumer services, and store the processed data in a PostgreSQL database and Redis cache. The project uses Docker Compose to manage multiple services including PostgreSQL, Redis, RabbitMQ, and two Go applications: csv-producer (which reads and sends CSV data) and csv-app (which processes and stores the data).

## Architecture Overview
CSV Producer: Reads a CSV file, processes its rows, and sends the data to a RabbitMQ message queue.

CSV App: Consumes messages from RabbitMQ, processes the data, and stores it in both PostgreSQL and Redis.

RabbitMQ: Message broker to facilitate communication between producer and consumer services.

PostgreSQL: Stores the processed data.

Redis: Caches the processed data for faster access.

## Features
CSV Ingestion: The csv-producer reads a CSV file and sends the data to a RabbitMQ queue.

Message Queue: RabbitMQ handles the message queue for asynchronous data processing.

Data Processing: The csv-app processes the messages, structures the data, and stores it in PostgreSQL and Redis.

API Server: The csv-app exposes a RESTful API for querying the processed data.

## Prerequisites

Make sure you have the following installed:

1. Docker

2. Docker Compose

3. Go (for building the csv-producer and csv-consumer services)

## Getting Started

* Clone the repository

```
git clone https://github.com/anudevSaraswat/csv-data-ingestion.git && cd csv-data-ingestion
```

* Environment Setup

Environment Variables for PostgreSQL:

Create a .env file in the root directory and define the environment variables required for PostgreSQL:

```
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_db_name
```

* Running the Project
  
The project uses Docker Compose to orchestrate all the services. You can build and run all the services with the following command:

```
docker compose up -d
```

This will:

Build all the services defined in the docker-compose.yml file. Start the containers for Redis, RabbitMQ, PostgreSQL, and the producer/consumer services.

## Accessing the Application

CSV Producer: This service doesn't expose any ports directly, but you can see the logs using docker logs <csv-producer-container-id> to verify it's reading and publishing messages to RabbitMQ.

CSV Consumer API: The consumer service exposes an API on port 8000. You can query the processed data using HTTP GET requests. Example:

```
curl http://localhost:8000/api/users
```

## Project Structure

```
├── csv-app/              # The consumer application (Go)
├── csv-producer/         # The producer application (Go)
├── database/             # Dockerfile for PostgreSQL setup
├── docker-compose.yml    # Docker Compose file to manage services
├── .env                  # Environment variables for docker-compose file
└── README.md             # This file
```

## Services Breakdown

1. **redis**: Uses `redis/redis-stack` image to run Redis with Redisearch module for indexing and searching.
2. **rabbitmq**: Uses the `rabbitmq` image with exposed ports for AMQP communication.
3. **postgres**: Builds from the ./database directory. The database stores user data processed from the CSV file.
4. **csv-producer**: Reads the people-10000.csv file and publishes data to RabbitMQ for processing.
5. **csv-consumer**: Listens to RabbitMQ and processes the data, storing it in both PostgreSQL and Redis.

The csv-producer and csv-consumer wait until RabbitMQ is available before attempting to connect.

## Troubleshooting

<ins>RabbitMQ Connection Issues</ins>: Ensure RabbitMQ is up and healthy before the csv-producer or csv-consumer attempts to connect. The wait-for-it.sh script is used to wait for RabbitMQ to become available.

<ins>PostgreSQL Issues</ins>: Ensure that the .env file is correctly set up with the required PostgreSQL credentials. Check the logs of the postgres container for any issues related to database initialization.
