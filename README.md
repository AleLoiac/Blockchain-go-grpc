# Simple Blockchain example using Go, gRPC and Badger
This project illustrates how a simple Blockchain works. It allows users to view and add nodes through a gRPC API.
Each node stores the transaction data of the purchase of a videogame.
If the game ID doesn't exist in the DB the transaction is blocked and the blockchain doesn't add the node.

## Getting Started

To get started, follow these steps:

Clone the repository:

`git clone https://github.com/AleLoiac/Blockchain-go-grpc.git`

Install the dependencies:

`go mod tidy`

Create a directory to save data and direct Badger to its path, for example:

`db, _ := badger.Open(badger.DefaultOptions("/Users/aless/Desktop/Go/Blockchain/DB"))`

Run the server:

`go run server.go`

## Usage

The blockchain supports the following API methods:

* AddBlock
* GetBlock
* GetBlockchain
* AddVideoGame

### Postman

You can use Postman to try the APIs, click on "NEW" and then on "gRPC Request".

Set the server URL to "0.0.0.0:50051" and select an API from the list.

Remember to use the Example Message feature to check how the requests are structured.
