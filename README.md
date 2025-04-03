# Blockchain Project

Welcome to this simple blockchain implementation! 🚀 This project is designed to help you understand the basics of how a blockchain works, including transactions, mining, and network consensus. Below, you'll find a breakdown of how everything fits together and how you can run it yourself.

---

## 📂 Project Structure

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
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksum file
├── utils
│   ├── ecdsa.go             # Handles cryptographic signing
│   ├── json.go              # JSON utilities
│   └── neighbor.go          # Handles peer-to-peer connections
├── wallet
│   └── wallet.go            # Wallet implementation
└── wallet_server
    ├── main.go              # Wallet API server
    ├── template             # HTML templates (if any UI is included)
    ├── wallet_server        # Additional wallet server logic
    └── wallet_server.go     # Handles wallet interactions
```

---

## ⚡ How It Works

### 1️⃣ Blockchain Basics
At its core, this project implements a basic **blockchain**, which is a distributed ledger that stores transactions securely. It consists of **blocks**, each containing:
- A timestamp ⏳
- A list of transactions 📜
- A reference (hash) to the previous block 🔗
- A unique identifier (nonce) to verify proof-of-work ✅

### 2️⃣ Transactions
Transactions are the foundation of the blockchain. A **wallet** creates transactions, signs them with a private key 🔑, and broadcasts them to the blockchain network. Transactions include:
- Sender address 📤
- Recipient address 📥
- Amount 💰

### 3️⃣ Mining ⛏️
Mining is the process of validating transactions and adding them to the blockchain. The blockchain server handles mining by:
1. Collecting pending transactions 📌
2. Finding a valid **nonce** (proof-of-work) 🧩
3. Creating a new block and adding it to the chain 🔄
4. Rewarding the miner with cryptocurrency 🏆

### 4️⃣ Wallets 🔐
Each user has a **wallet** that generates a public/private key pair for secure transactions. The wallet server handles:
- Creating a new wallet 🆕
- Signing transactions ✍️
- Sending transactions to the blockchain 📡

### 5️⃣ Blockchain Server 🌍
The **blockchain server** is the core of the network. It:
- Maintains the blockchain ledger 📖
- Processes transactions and blocks 🏗️
- Supports API endpoints to interact with the blockchain 🔌
- Ensures network consensus by resolving conflicts ⚖️

---

## 🚀 Running the Project

### Step 1: Install Dependencies
Make sure you have Go installed (v1.23.4+). If not, install it from [Go's official site](https://go.dev/).

Clone this repository:
```bash
git clone https://github.com/your-repo/blockchain.git
cd blockchain
```

Initialize the Go module:
```bash
rm go.mod go.sum # Remove old module files if needed
go mod init blockchain
go mod tidy # Install dependencies
```

### Step 2: Start the Blockchain Server
```bash
go run blockchain_server/main.go
```
This starts the blockchain server on the default port (6000).

### Step 3: Start the Wallet Server (Optional)
```bash
go run wallet_server/main.go
```
This allows wallets to create and send transactions.

### Step 4: Interact with the Blockchain
You can use HTTP requests to interact with the blockchain. Some example endpoints:

- **Get the blockchain:**
  ```bash
  curl http://localhost:6000/
  ```
- **Mine a block:**
  ```bash
  curl http://localhost:6000/mine
  ```
- **View pending transactions:**
  ```bash
  curl http://localhost:6000/transactions
  ```

---

## 🛠️ Future Improvements
While this project is a great starting point, there are several enhancements that can be made:
- ✅ Add peer-to-peer networking for a fully decentralized blockchain
- ✅ Implement proof-of-stake (PoS) as an alternative consensus mechanism
- ✅ Improve the wallet UI for easier transaction management

