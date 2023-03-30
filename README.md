# Simple Blockchain example using Go, gRPC and Badger
This project illustrates how a simple blockchain works.
It allows users to view and add nodes through a gRPC API.
Each node stores transaction data for the purchase of a video game.
If the game ID does not exist in the database, the transaction is blocked and the blockchain does not add the node.

<div align="center">

![](https://github.com/AleLoiac/Blockchain-go-grpc/blob/master/Blockchain_example.jpg)

</div>

Each new block in the blockchain is created by computing a hash value based on the data it contains, including the transaction data, and the hash of the previous block in the chain.
This hash value serves as a unique identifier for the block and ensures that any modification to the block or the chain would be immediately detectable.
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
* GetVideoGame
* ListVideoGames

### Postman

You can use Postman to try the APIs, click on "NEW" and then on "gRPC Request".

Set the server URL to "0.0.0.0:50051" and select an API from the list.

Remember to use the Example Message feature to check how the requests are structured.