# Golang Microservice
This repository contains a microservice architecture built with Golang, Docker, and gRPC. The services communicate with each other through gRPC, and Nginx is used as a reverse proxy. PostgreSQL is used as the database for the services.

# Overview
This project aims to provide a scalable microservice architecture with Go, Docker, and gRPC. The current setup includes:

1. User Service: Handles user management.
2. Auth Service: Handles authentication and user registration.
3. Chat Service: Handlers chat management.

# Setup
To set up and run this project locally, follow these steps:

1. Clone the repository
``` 
git clone https://github.com/insanXYZ/golang-microservice.git
```
2. Go to directory
```
cd golang-microservice
```
3. Build the Docker containers:
```
docker-compose build

docker-compose start
```