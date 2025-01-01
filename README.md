# Koyeb Task 

This project leverages Docker and Nomad to set up and run a web server using Nginx. Below are the requirements and setup instructions to get started.  

## Requirements  

Before starting, make sure you have the following tools installed on your system:  

- **Docker**: A containerization platform for running and managing containers.  
- **Homebrew**: A package manager for macOS (or Linux).  
- **Nomad**: A workload orchestrator by HashiCorp.  

## Installation and Setup  

### Step 1: Install Docker  

Ensure Docker is installed and running on your machine. For installation instructions, refer to the [official Docker documentation](https://docs.docker.com/get-docker/).  

### Step 2: Install Nomad  

Use Homebrew to install HashiCorp Nomad on Mac OS:  

```bash
brew tap hashicorp/tap  
brew install hashicorp/tap/nomad
```

### Step 3: Pull required docker image

Install the required Docker image for Nginx:  

```bash
make docker-nginx 
```

### Step 4: Golang dependences

Install the required Golang dependences:  

```bash
make setup
```

## Commands

### Start the web server

To start the web server, run the following command:  

### **Default Target**  
**`all`**  
- Runs the following targets in sequence: `nomad-start`, `wait-for-nomad`, and `run`.  
- This ensures Nomad is started, waits until it's ready, and then runs your Go application.  

```bash
make all
```

**`test`**
- Runs the tests for the Go application.  

```bash
make test
```


### **Nomad Targets**

**`nomad-start`**
- Starts the Nomad server in the background.  

```bash
make nomad-start
```

**`nomad-stop`**
- Stops the Nomad server.  

```bash
make nomad-stop
```

**`nomad-node-status`**
- Shows the status of the Nomad server.  

```bash
make nomad-node-status
```

### **Docker Targets**

**`docker-nginx`**
- Pulls the required Docker image for Nginx.  

```bash
make docker-nginx
```

## Provided Links

- Nomad UI - http://localhost:4646/ui
- Pastbin Script version -  https://pastebin.com/raw/vDzvUbJS
- Pastbin Text version - https://pastebin.com/raw/xp02ittK

## Files and Directories

- **`/utils`**: Contains utility functions for the Go application.
- **`/models`**: Contains the data models for the Go application.
- **`/server`**: Contains the server logic for the Go application.
- **`/main.go`**: The main entry point for the Go application.
- **`/Makefile`**: Contains the commands to build, run, and test the Go application.
- **`/service`**: Contains the endpoint for the Go application.
- **`/.github`**: Contains the GitHub Actions workflow for the Go application.

## Endpoints

- **`POST /services/:name`**: Creates a new service with the given name and return the URL to access it.

- **`PUT /services/:name`**: Creates a new service with the given name and return the URL to access it.

- **`GET /services`**: Returns a list of all services available.

## Conclusion

This project provides a simple way to set up and run a web server using Docker and Nomad. It includes instructions for installation and setup, as well as commands to start the web server and manage the Nomad server.