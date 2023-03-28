# Simple Blockchain example using Go, gRPC and Badger
This is a blockchain built with Go and gRPC. It allows users to view and add nodes through a gRPC API.

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

### Postman

You can use Postman to try the APIs, click on "NEW" and then on "gRPC Request".

Set the server URL to "0.0.0.0:50051" and select an API from the list.

Remember to use the Example Message feature to check how the requests are structured.
