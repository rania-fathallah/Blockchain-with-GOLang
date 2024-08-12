package main

import (
	"crypto/sha256"
	"time"
)
// Block keeps block headers
type Block struct {
	Hash          []byte
	PrevBlockHash []byte
	Transactions  []*Transaction
	Timestamp     int64
	Nonce         int
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTXs() []byte {
	var txHashes [][]byte
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.Hash)
	}
	result:= sha256.Sum256(Serialize(txHashes))
	
	return result[:]
}

// NewBlock creates and returns Block, for testing purposes, mining
// can be activated or disabled using the boolean flag mine If mine is
// set to true, the nonce is computed so that the hash start with
// zeroBits
func NewBlock(txs []*Transaction, prevBHash []byte, mine bool, zeroBits int) *Block {
	block := &Block{[]byte{},  prevBHash, txs, time.Now().Unix(), 0}
	//If mine==true we mine the block
	if mine{
		block.Mine(zeroBits)
	}
	return block
}

// Creates and returns genesis Block, its hash must start with zeroBits
func NewGBlock(cbtx []*Transaction, zeroBits int) *Block {
	genesisBlock := NewBlock(cbtx, []byte{}, true, zeroBits);
	return genesisBlock;
}


// true if the block is correclty Hashed 
func (block *Block) IsCorrectlyHashed(zeroBits int) bool {	
	//To check if a block is correctly hashed we check if it it's hash start with the needed number of zeros
	return StartsWithXZeros(block.Hash, zeroBits)
}

// Hashes a block, private fnuction 
func (block *Block) computeHash() []byte {
	contents := Serialize([][]byte{
		block.PrevBlockHash,
		block.HashTXs(),
		IntToHex(block.Timestamp),
		IntToHex(int64(block.Nonce)),
	})
	result:=sha256.Sum256(contents)
	return result[:]
}		

// Computes and sets the hash of "block"
func (block *Block) SetHash(){
	//using the computerHash function to compute the hash of a Block block and set it's hash with the returned value
	block.Hash = block.computeHash()
}		

// Mines a block : iterates over nonces until the hash starts with the
// number of zeros defined by zeroBits
func (block *Block) Mine(zeroBits int) {
	block.Nonce=0;
	//A loop similar to while(true) loop to set the correct hash
	//we change the Nonce of the block in hope to get a correct hash
	for{
		block.SetHash()
		//The loop breaks if the block is correctly hashed
		if(block.IsCorrectlyHashed(zeroBits)){
			break;
		}
		block.Nonce++;
	}	
}		
