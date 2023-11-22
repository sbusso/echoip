# EchoIP Service

## Overview

EchoIP is a simple Go service using the Chi router framework. It returns the real IP address of the client making the request. The service can run as a standalone binary or inside a Docker container.

## Local Setup

### Prerequisites

- Go (version 1.18 or later)

### Running Locally

1. Clone the repository:

```bash
 git clone git@github.com:nuibits/echoip.git
 cd echoip
```

2. Build and run the service:

```bash
go build -o echoip
./echoip
```

By default, the service runs on port 3000. To use a different port, set the PORT environment variable:

```bash
PORT=8080 ./echoip
```

## Docker Setup

### Prerequisites

- Docker

### Building and Running with Docker

1. Build the Docker image:

Specify the port as a build argument (default is 3000):

```bash
docker build -t echoip .
```

2. Run the EchoIP container:

Replace `[host-port]` and `[container-port]` with the desired port numbers. For example:

```bash
docker run -p 8080:8080 -e PORT=8080 echoip
```
