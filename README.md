# Kafka High-Availability Setup with Go Integration

A production-ready Kafka cluster setup with high availability, monitoring, and Go-based producer/consumer applications.

## 📋 Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Architecture](#architecture)
- [Quick Start](#quick-start)
- [Detailed Setup](#detailed-setup)
- [Monitoring](#monitoring)
- [Development](#development)
- [Troubleshooting](#troubleshooting)
- [License](#license)

## ✨ Features

- **High Availability**
  - 3-node Kafka cluster
  - Automatic leader election
  - Replication factor of 3
  - Multiple partitions per topic

- **Monitoring & Management**
  - Kafka UI web interface
  - Real-time metrics
  - Topic management
  - Consumer group tracking

- **Development Tools**
  - Go-based producer example
  - Go-based consumer example
  - Graceful shutdown handling
  - Error handling and retries

- **Security & Access**
  - External access configuration
  - Configurable security settings
  - Network isolation

## 🔧 Prerequisites

- Docker Engine 20.10+
- Docker Compose v2.0+
- Go 1.16 or later
- 4GB RAM minimum
- 10GB free disk space

## 🏗 Architecture

### Components
- **ZooKeeper**: Cluster coordination
- **Kafka Brokers**: 3-node cluster
- **Kafka UI**: Web-based monitoring
- **Go Applications**: Producer/Consumer examples

### Port Configuration
- Kafka Brokers:
  - Broker 1: 9092 (internal), 29092 (external)
  - Broker 2: 9093 (internal), 29093 (external)
  - Broker 3: 9094 (internal), 29094 (external)
- ZooKeeper: 2181
- Kafka UI: 8080

### Topics
- Default topics created:
  - test-topic (3 partitions, RF=3)
  - dev-topic (3 partitions, RF=3)

## 🚀 Quick Start

1. Clone the repository:
```bash
git clone https://github.com/yourusername/kafka-setup-go.git
cd kafka-setup-go
```

2. Start the Kafka cluster:
```bash
docker-compose up -d
```

3. Install Go dependencies:
```bash
go mod tidy
```

4. Run the consumer:
```bash
go run consumer/main.go
```

5. Run the producer:
```bash
go run producer/main.go
```

## 📖 Detailed Setup

### Kafka Cluster Configuration
The cluster is configured for high availability with:
- 3 Kafka brokers for redundancy
- ZooKeeper for cluster management
- Automatic topic creation
- Configurable retention policies

### Producer Application
- Sends timestamped messages
- Configurable message rate
- Handles broker failures
- Supports message batching

### Consumer Application
- Group-based consumption
- At-least-once delivery
- Offset management
- Graceful shutdown

## 📊 Monitoring

Access Kafka UI at http://localhost:8080 for:
- Topic list and configurations
- Consumer group status
- Message browser
- Cluster health metrics
- Broker status

## 💻 Development

### Adding New Topics
```bash
docker-compose exec kafka1 kafka-topics --create \
    --bootstrap-server kafka1:9092 \
    --topic your-topic \
    --partitions 3 \
    --replication-factor 3
```

### Modifying Applications
- Producer: `producer/main.go`
- Consumer: `consumer/main.go`
- Configuration: Environment variables in `docker-compose.yml`

## 🔍 Troubleshooting

### Common Issues
1. **Connection Refused**
   - Check if all containers are running
   - Verify port mappings
   - Ensure no port conflicts

2. **Producer/Consumer Errors**
   - Verify broker connectivity
   - Check topic existence
   - Review log levels

### Health Checks
```bash
# Check cluster status
docker-compose ps

# View logs
docker-compose logs -f kafka1

# List topics
docker-compose exec kafka1 kafka-topics --list --bootstrap-server kafka1:9092
```

## 📄 License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

### What this means:
- ✔️ Commercial use
- ✔️ Modification
- ✔️ Distribution
- ✔️ Patent use
- ✔️ Private use

### Requirements:
- 📝 License and copyright notice
- 📝 State changes
- 📝 Disclose source
- 📝 Same license

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## 📬 Contact

For questions and support, please open an issue in the GitHub repository.
