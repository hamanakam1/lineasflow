# LineasFlows Edge-Mesh Daemon

[![CI](https://github.com/hamanakam1/lineasflow/actions/workflows/ci.yml/badge.svg)](https://github.com/hamanakam1/lineasflow/actions/workflows/ci.yml)
[![Release](https://github.com/hamanakam1/lineasflow/actions/workflows/release.yml/badge.svg)](https://github.com/hamanakam1/lineasflow/actions/workflows/release.yml)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go: 1.22+](https://img.shields.io/badge/Go-1.22%2B-00ADD8.svg)](go.mod)
[![Releases](https://img.shields.io/github/v/release/hamanakam1/lineasflow)](https://github.com/hamanakam1/lineasflow/releases)

LineasFlows is a high-performance edge-mesh synchronizer daemon (the "Ticket Dispenser") for the Bumperfort Fortran physics coprocessor. It reads calculated high-dimensional state outputs from Bumperfort via POSIX shared memory (zero-copy) and synchronizes states across a local mesh of Apple Silicon (M3 Max) and Raspberry Pi 5 nodes.

![Zero-allocation C-ABI encode/decode verification and throughput benchmark](docs/screenshots/test-bench.png)

## Architecture
- `cmd/lineasflowd/` daemon entrypoint
- `internal/shm/` cgo POSIX shared memory bridge (lock-free dirty bit reader)
- `internal/mesh/` raw UDP multicast broadcaster (no JSON/protobuf; raw binary)
- `pkg/types/` C-ABI mirrored structs

## Build
```
make build-mac
make build-pi
```

## Constraints
- No serialization overhead: raw binary payload only.
- Zero-allocation hot path for shared memory polling.
