# H3Net Protocol Client

This repository contains the official client libraries for the [H3Net Protocol](https://h3net.switchnomix.com). The H3Net Protocol is a modern, efficient protocol for network device management and operations, built on top of HTTP/3 with WebTransport support.

## Repository Structure

```console
h3netprotocol-client/
├── go/                   # Go client implementation
│   ├── pkg/              # Generated Go client code
│   │   ├── api/          # Core API types and interfaces
│   │   ├── models/       # Generated data models
│   │   └── transport/    # WebTransport implementation
│   ├── examples/         # Go examples
│   │   ├── hello/        # Hello message example
│   │   ├── health/       # Health check example
│   │   └── capabilities/ # Device capabilities example
│   └── go.mod            # Go module
├── rust/                 # Rust client implementation (coming soon)
├── schemas/             # OpenAPI schemas
│   ├── core/            # Core protocol definitions
│   ├── config/          # Configuration schemas
│   ├── discovery/       # Discovery schemas
│   ├── operations/      # Operations schemas
│   ├── pipeline/        # Pipeline schemas
│   ├── security/        # Security schemas
│   └── telemetry/       # Telemetry schemas
└── README.md            # This file
```

## Supported Languages

Currently supported client implementations:

- Go (v1.22+)
  - HTTP/3 with WebTransport support
  - Full protocol implementation
  - Type-safe API
  - Comprehensive examples

Future planned implementations:

- Rust
  - HTTP/3 with WebTransport support
  - Zero-copy operations
  - Async/await support
  - Memory safety guarantees

## Installation

### Go Client

1. Initialize a new Go module (if not already done):

```bash
go mod init your-module-name
```

2. Add the H3Net Protocol client to your project's `go.mod`:

```go
require github.com/switchnomix/h3netprotocol-client v0.1.1 // Use latest version

replace github.com/switchnomix/h3netprotocol-client => github.com/switchnomix/h3netprotocol-client/go/pkg v0.1.1
```

3. Download dependencies:

```bash
go mod tidy
go mod download github.com/switchnomix/h3netprotocol-client@go/pkg/v0.x.x
```

4. Import the package in your Go code:

```go
import (
    h3netprotocol "github.com/switchnomix/h3netprotocol-client"
)
```

Requirements:

- Go 1.22 or later
- HTTP/3 with WebTransport support (uses `github.com/quic-go/webtransport-go`)
- QUIC transport layer support

> **Note**: This client uses HTTP/3 with WebTransport for bidirectional streaming communication. It does not use the standard `net/http` package's HTTP/1.1 or HTTP/2 implementations.

> **Version Note**: The repository uses tags in the format `go/pkg/v0.x.x` for versioning. When specifying versions in your `go.mod`, use the version number without the `go/pkg/` prefix (e.g., `v0.1.1` instead of `go/pkg/v0.1.1`).

## Quick Start

### Go Client Example

```go
package main

import (
    "context"
    "crypto/tls"
    "encoding/json"
    "fmt"
    "log"
    "time"

    h3netprotocol "github.com/switchnomix/h3netprotocol-client"
    "github.com/quic-go/webtransport-go"
)

// Message represents a WebTransport message
type Message struct {
    Type    string      `json:"type"`
    Payload interface{} `json:"payload"`
}

const (
    webTransportEndpoint = "https://your-server.com:9000/webtransport"
)

func main() {
    // Create WebTransport client with TLS config
    client := &webtransport.Dialer{
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: true, // For development only
        },
    }

    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Connect to the server
    log.Printf("Connecting to endpoint: %s", webTransportEndpoint)
    _, session, err := client.Dial(ctx, webTransportEndpoint, nil)
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer session.CloseWithError(0, "client closing")

    log.Println("Connected successfully, waiting for health status...")

    // Accept stream from server
    stream, err := session.AcceptStream(ctx)
    if err != nil {
        log.Fatalf("Failed to accept stream: %v", err)
    }
    defer stream.Close()

    // Decode the HEALTH_STATUS response
    var response Message
    if err := json.NewDecoder(stream).Decode(&response); err != nil {
        log.Fatalf("Failed to decode response: %v", err)
    }

    if response.Type != "HEALTH_STATUS" {
        log.Fatalf("Expected HEALTH_STATUS response, got %s", response.Type)
    }

    // Log health status
    health := response.Payload.(map[string]interface{})
    status := health["status"].(map[string]interface{})
    components := health["components"].(map[string]interface{})
    metrics := health["metrics"].(map[string]interface{})

    log.Printf("Server health status: %s", status["state"])
    log.Printf("Uptime: %s", health["uptime"])
    
    // Log component statuses
    for component, status := range components {
        log.Printf("Component %s: %s", component, status)
    }
    
    // Log metrics
    log.Printf("Active connections: %.0f", metrics["activeConnections"])
    log.Printf("CPU usage: %.2f%%", metrics["cpuUsage"])
    log.Printf("Memory usage: %.2f%%", metrics["memoryUsage"])
}
```

## Protocol Features

The H3Net Protocol client supports:

1. **Core Operations**
   - Device discovery and registration
   - Health monitoring
   - Capability negotiation

2. **Configuration Management**
   - Device configuration
   - VLAN management
   - ACL configuration
   - Routing setup

3. **Telemetry**
   - Interface statistics
   - Resource metrics
   - Environmental data

4. **Security**
   - Role-based access control
   - API key authentication
   - TLS encryption

## Documentation

- [Go Client Documentation](go/docs/README.md)
- [API Reference](schemas/docs/api_endpoints.md)
- [Examples](go/examples/)
- [Protocol Specification](schemas/docs/protocol.md)

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Support

For support, please:

1. Check the documentation
    - [swagger-ui](https://h3net.switchnomix.com/api/)
    - [Specification](https://h3net.switchnomix.com/specification/H3Net_Protocol_Specification.html)
    - [local](go/pkg/docs/)
2. Open an issue in this repository
3. Join our user group: <h3netprotocol@googlegroups.com>

## Security

Please report security issues to <security@switchnomix.com>

## Sponsorship and Collaboration

We welcome sponsorships and collaborations to help advance the H3Net Protocol ecosystem. Whether you're interested in:

- Supporting the development of new client implementations
- Contributing to protocol enhancements
- Collaborating on integration projects
- Sponsoring documentation and examples
- Supporting community events and workshops

Please reach out to us at [info@switchnomix.com](mailto:info@switchnomix.com) to discuss how we can work together to make the H3Net Protocol even better.
