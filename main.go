package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// Blockchain represents the blockchain as a slice of blocks
type Blockchain []Block

// calculateHash calculates the hash of a block
func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash)
	hash := sha256.New()
	hash.Write([]byte(record))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

// generateBlock creates a new block in the blockchain
func generateBlock(oldBlock Block, data string) Block {
	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// isBlockValid checks if a block is valid by verifying its hash
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func main() {
	// Create the genesis block
	genesisBlock := Block{0, time.Now().String(), "Genesis Block", "", ""}
	genesisBlock.Hash = calculateHash(genesisBlock)

	// Create a blockchain and add the genesis block
	var blockchain Blockchain
	blockchain = append(blockchain, genesisBlock)

	// Add a new block to the blockchain
	newBlock := generateBlock(blockchain[len(blockchain)-1], "Transaction Data")
	blockchain = append(blockchain, newBlock)

	// Validate the blockchain
	valid := isBlockValid(newBlock, blockchain[len(blockchain)-2])

	fmt.Printf("Genesis Block: %+v\n", genesisBlock)
	fmt.Printf("New Block: %+v\n", newBlock)
	fmt.Printf("Blockchain Validity: %t\n", valid)
}
