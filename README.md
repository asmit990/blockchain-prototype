# Blockchain (Go) — Educational Implementation

> Built end‑to‑end by **Asmit**. MIT‑licensed, production‑unsafe by design. Use it to learn, not to custody funds.

A compact, readable blockchain written in Go. It covers wallets, signed transactions, mining with proof‑of‑work, a simple HTTP API, and basic conflict resolution between peers. The focus is clarity over cleverness: short files, direct logic, plenty of comments.

---

## Features

* **Blocks** with timestamp, transactions, previous hash, nonce (PoW)
* **Wallets** with ECDSA keypairs and transaction signing
* **Mempool** for pending transactions
* **Mining** endpoint that collects txs, finds a valid nonce, mints a reward, and appends a block
* **HTTP servers** for the blockchain and the wallet
* **Naïve consensus** (longest‑chain wins) stubbed via neighbor discovery utilities

> **Note**: This is an instructional codebase. It omits many things you’d need in the wild: chain reorgs, DoS hardening, fee markets, UTXO/Account pruning, full P2P, robust validation, etc.

---

## Project Structure

```
├── LICENSE
├── README.md
├── block
│   └── blockchain.go        # Core blockchain logic
├── blockchain_server
│   ├── blockchain_server.go # Server for interacting with blockchain
│   └── main.go              # Entry point for the blockchain server
├── cmd
│   └── main.go              # CLI interface (optional)
├── error.log                # Log file
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
├── utils
│   ├── ecdsa.go             # Cryptographic signing (ECDSA)
│   ├── json.go              # JSON helpers
│   └── neighbor.go          # Peer discovery / neighbor handling (naïve)
├── wallet
│   └── wallet.go            # Wallet implementation
└── wallet_server
    ├── main.go              # Wallet API server
    ├── template             # Optional HTML templates
    ├── wallet_server        # Additional wallet server logic
    └── wallet_server.go     # Wallet HTTP handlers
```

---

## How it Works (High‑Level)

### 1) Blockchain Basics

Each block holds:

* **Timestamp**
* **Transactions** (array)
* **PrevHash** (link to previous block)
* **Nonce** (proof‑of‑work)
* **Hash** (block header hash)

### 2) Transactions

Wallets create transactions and sign them with the sender’s private key. A transaction contains:

* **From** (public address)
* **To** (recipient address)
* **Amount**
* **Signature** (ECDSA over the payload)

### 3) Mining (Proof‑of‑Work)

* Collect transactions from the mempool
* Build a candidate block
* Increment nonce until `hash(header) < target`
* On success: append block, clear included txs, credit miner reward

### 4) Wallets

* Generate ECDSA keypair (private/public)
* Derive a public address
* Sign transactions locally and submit to the node

### 5) Server & Consensus

* **Blockchain Server** exposes REST endpoints to query chain, submit transactions, and mine
* **Consensus**: longest‑chain wins (placeholder). `utils/neighbor.go` contains simple neighbor mechanics you can extend to real P2P gossip/sync

---

## Requirements

* **Go 1.23.4+** (or newer)
* A POSIX‑like shell (macOS/Linux) or PowerShell on Windows

---

## Setup

```bash
# Clone
git clone https://github.com/asmit990/blockchain.git
cd blockchain

# Reset/initialize module files (optional)
rm -f go.mod go.sum

go mod init blockchain
go mod tidy
```

> If you already have a `go.mod`, you can skip the removal. `go mod tidy` will pull exact dependencies.

---

## Run

### 1) Start the Blockchain Server

```bash
go run blockchain_server/main.go
```

Default port: **6000**.

### 2) Start the Wallet Server (optional)

```bash
go run wallet_server/main.go
```

Default port: **7000** (or as defined in the code).

---

## API Quickstart (HTTP)

Use `curl` or your favorite REST client.

**Get the blockchain**

```bash
curl http://localhost:6000/
```

**Mine a block**

```bash
curl http://localhost:6000/mine
```

**View pending transactions**

```bash
curl http://localhost:6000/transactions
```

**Submit a transaction** *(via wallet server submitting to blockchain)*

```bash
# Example body may differ based on your wallet_server handlers
curl -X POST http://localhost:7000/transaction \
  -H 'Content-Type: application/json' \
  -d '{
        "from":   "<sender-address>",
        "to":     "<recipient-address>",
        "amount": 10
      }'
```

---

## Configuration

* **Ports**: See `main.go` files in `blockchain_server/` and `wallet_server/`
* **Mining reward / difficulty**: Defined in `block/blockchain.go` (search for constants). Tune for faster local testing.
* **Logging**: Output to `error.log` where applicable.

---

## Notes on Security & Limitations

* ECDSA keys are not stored securely. Use a proper keystore for anything serious.
* No persistent database: the chain lives in memory unless you add storage.
* Consensus is naïve. There is no Sybil resistance, no finality, no network partition handling.
* No fee market or mempool policy (RBF, min‑fee, eviction). Transactions are accepted on sight.
* Mining is intentionally easy; difficulty targets are for demos.

This is a learning scaffold. Treat it that way.

---

## Roadmap / Future Work

* [ ] Real P2P networking (gossip, inventory, block/tx relay)
* [ ] Robust chain sync & reorg handling
* [ ] Proof‑of‑Stake (PoS) experimental path
* [ ] Wallet UI (basic React) and mnemonic/keystore support
* [ ] Persistent storage (BoltDB/Badger/SQLite)
* [ ] Fee market + mempool rules
* [ ] Unit tests & property tests

---

## Development

```bash
# Lint/format
go fmt ./...

# Build binaries (optional)
go build -o bin/blockchain ./blockchain_server
go build -o bin/wallet     ./wallet_server

# Run tests (add as you go)
go test ./...
```

Coding style: keep functions small, favor clarity over abstraction, document invariants near the code that enforces them.

---

## License

MIT. See `LICENSE`.

---

## Attribution

Made entirely by **Asmit**. If you fork or extend, please keep author credit in the README and commit history.

