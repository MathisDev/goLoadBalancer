**Project Name: Lightweight Load Balancer**

---

### Introduction
This project introduces a lightweight load balancer implemented in Go. Unlike traditional load balancers that act as proxies, this load balancer functions by intelligently distributing incoming traffic across multiple backend servers without proxying the data. It aims to provide high availability, scalability, and fault tolerance for distributed systems by efficiently managing incoming requests.

### Key Features
- **Non-Proxy Based**: Does not proxy traffic, allowing for faster request processing and reduced overhead.
- **Dynamic Load Distribution**: Utilizes various load balancing algorithms (e.g., Round Robin, Least Connections) to distribute traffic evenly across backend servers.
- **Health Checking**: Regularly checks the health status of backend servers to ensure that only healthy servers receive traffic.
- **Scalability**: Easily scalable architecture to handle increasing loads by adding or removing backend servers dynamically.
- **Fault Tolerance**: Automatically detects and redirects traffic away from failed servers, ensuring uninterrupted service.
  
### Installation
1. Install Go (if not already installed).
2. Clone the repository:
   ```bash
   git clone https://github.com/your/repository.git
