# ICMP Ping & HTTP GET with Prometheus Metrics

## Overview
This project is a simple Go application that performs:
- ICMP ping to a specified hostname/IP and records success/failure.
- HTTP GET requests to `google.com` and tracks response time.
- Exposes Prometheus metrics for monitoring.

## Features
- **ICMP Ping Monitoring:**
  - Sends an ICMP echo request to a specified host.
  - Records response time in milliseconds.
  - Tracks success/failure.
- **HTTP GET Monitoring:**
  - Sends an HTTP GET request to `https://www.google.com`.
  - Measures response time in milliseconds.
- **Prometheus Metrics Exposure:**
  - `icmp_ping_success`: 1 if the last ping was successful, 0 otherwise.
  - `icmp_ping_response_time`: Response time of the last ICMP ping.
  - `http_response_time`: Response time of the last HTTP GET request.
  - Metrics are exposed at `/metrics` (default port: 8080).

## Prerequisites
- Go installed (version 1.19 or higher recommended).
- Prometheus installed and configured to scrape the `/metrics` endpoint.

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/Thobani36/Split-Point-Assignment.git
   cd Split-Point-Assignment
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage
1. **Run the application:**
   ```sh
   go run ping.go --host=google.com --port=8080
   ```
   - `--host`: Target hostname/IP for ICMP ping (default: `google.com`).
   - `--port`: Port for Prometheus metrics server (default: `8080`).

2. **View logs in the console:**
   ```
   Ping to google.com: 10.23 ms
   HTTP GET to google.com took 482.45 ms
   ```

3. **Access Prometheus metrics:**
   Open a browser and go to:
   ```
   http://localhost:8080/metrics
   ```

## Example Output (Prometheus Metrics)
```
# HELP icmp_ping_success 1 if the last ping was successful, 0 if it failed.
# TYPE icmp_ping_success gauge
icmp_ping_success 1

# HELP icmp_ping_response_time Response time of the last ICMP ping in milliseconds.
# TYPE icmp_ping_response_time gauge
icmp_ping_response_time 10.23

# HELP http_response_time Response time of the last HTTP GET request in milliseconds.
# TYPE http_response_time gauge
http_response_time 482.45
```

## Project Structure
```
.
├── ping.go          # Main Go program
├── go.mod           # Go module file
├── README.md        # Documentation
```

## Challenges Faced & Solutions
- **Learning Golang**: Had no prior experience, so I studied Go documentation and practiced coding.
- **Networking (Ping, ICMP, Prometheus)**: Researched ICMP, used `go-ping`, and learned Prometheus monitoring.
- **Data Formatting**: Ensured response times were displayed with two decimal precision.
- **Efficient Metrics Exposure**: Registered only required metrics for Prometheus to optimize performance.


