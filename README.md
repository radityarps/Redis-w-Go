# Redis with Go

A lightweight, modular in-memory key-value store built in Go to explore advanced Go design patterns, interfaces, and polymorphism.

## Overview

This project is a technical learning and portfolio project aimed at mastering Go fundamentals and advanced design principles. It demonstrates clean architecture and interface segregation by decoupling storage implementations behind a unified contract (`store.Storer`), allowing transparent swapping between standard in-memory storage and TTL-based expiration storage.

## Features

- **Interface-Driven Design (Polymorphic Storage)**: Abstracted via `store.Storer` interface, enabling any compliant storage engine to plug into existing application logic.
- **In-Memory KV Store (`store/kv`)**:
  - Basic key-value operations (`Get`, `Set`, `Delete`, `Len`).
  - Configurable capacity limits (`maxKeys`) with `ErrMaxKeysExceeded` validation.
  - Sorted key retrieval (`Keys()`).
- **TTL Store (`store/ttl`)**:
  - Key expiration with configurable durations.
  - Passive eviction upon expired access.
- **Unit Tested**: Comprehensive test suite covering roundtrips, empty key validation, table-driven deletion tests, and TTL expirations.

## Project Structure

```text
.
├── go.mod
├── main.go             # Entrypoint demonstrating polymorphic store usage
└── store/
    ├── store.go        # Storer interface and shared error definitions
    ├── kv/             # Standard in-memory store with capacity bounds
    │   ├── kv.go
    │   └── kv_test.go
    └── ttl/            # Time-To-Live (TTL) expiration store
        ├── ttl_store.go
        └── ttl_store_test.go
```

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.22+ installed.

### Running the Example

Run the entrypoint demonstration:

```bash
go run main.go
```

### Running Tests

Execute all unit tests across packages:

```bash
go test -v ./...
```

## Roadmap

- [ ] **Concurrency & Thread-Safety**: Protect storage engines from race conditions using `sync.RWMutex`.
- [ ] **RESP (Redis Serialization Protocol)**: Implement a custom RESP parser and serializer.
- [ ] **TCP Server**: Expose the store via TCP so standard `redis-cli` can connect and issue commands.
- [ ] **Advanced Eviction Policies**: Support LRU (Least Recently Used) and LFU (Least Frequently Used) strategies when capacity is reached.
- [ ] **Data Persistence**: Implement Append-Only File (AOF) logging and point-in-time snapshotting (RDB).
