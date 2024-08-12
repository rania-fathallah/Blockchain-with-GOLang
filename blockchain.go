package main

import (

)
//A structure created to help in the search for non spent transactions
type CustomStruct struct {
	Tid []byte
	Index     int
	TXOutputValue TXOutput
}

// Blockchain implements interactions with a DB
type Blockchain struct {
	GHash []byte // Hash of the Genesis Block
	Chain []Block// Slice of blocks
}

// AddBlock adds a new block with the provided transactions, the block
// is mined before addition 
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	//We create the new block with mine set to true so that it is mined
	b := NewBlock(transactions, prevBlock.Hash, true, 0)
	//We add the block to the blockchain
	bc.Chain = append(bc.Chain, *b)
}

// NewSendTransaction creates a set of transactions to transfert the
// amount from sender to receiver. If the sender has not the required
// amount, an error is returned 
func (bc *Blockchain) NewTransfertTX(from, to string, amount int) (*Transaction, error) {
	//an array of transaction inputs
	var TxIns []TXInput
	//an array of transaction output
	var TxOuts []TXOutput
	//an array of the unspent outputs
	var unspentOutputs []CustomStruct
	//the balance of the giver
	balance := bc.GetBalance(from)
	//if balance is less than the amount to transfer the 
	//transfer doesn't happen and an error is raised
	if balance < amount {
		return nil, ErrInsufficientFunds
	}
	//else the tranfert happen

	//Search for the unspent outputs the sender has
	for _, b := range bc.Chain {
		for _, t := range b.Transactions {
			//Add all transactions output in the transaction t to the unspentOutputs array
			for i, out := range t.TxOuts {
				if out.ScriptPubKey == from {
					unspentOutputs = append(unspentOutputs, CustomStruct{
						Tid:           t.Hash,
						Index:         i,
						TXOutputValue: out,
					})
				}
			}
			//By comparing the transactions inputs in the transaction t to the transaction outputs in
			//the array unspentOutputs 
			for _, tin := range t.TxIns {
				for j, unspent := range unspentOutputs {
					if EqualSlices(tin.Txid, unspent.Tid) && tin.Vout == unspent.Index {
						unspentOutputs = append(unspentOutputs[:j], unspentOutputs[j+1:]...)
					}
				}
			}
		}
	}
	for _, u := range unspentOutputs {
		TxIns = append(TxIns, TXInput{
			Txid:      u.Tid,
			Vout:      u.Index,
			ScriptSig: from,
		})
	}
	TxOuts = append(TxOuts, TXOutput{
		Value:        amount,
		ScriptPubKey: to,
	})
	if balance > amount {
		TxOuts = append(TxOuts, TXOutput{
			Value:        balance - amount,
			ScriptPubKey: from,
		})
	}
	tx := NewTransaction([]byte{}, TxIns, TxOuts)
	tx.ComputeHash()

	return tx, nil
}


// CreateBlockchain creates a new blockchain, evey adress in adresses
// is given the initial 
func NewBlockchain(addresses []string) *Blockchain {
	//the difficulty setting was chosen in random
	genesisBlock := NewGBlock([]*Transaction{}, 16)
    // Distribute initial rewards to addresses in the genesis block
    for _, address := range addresses {
        coinbaseTx := NewCoinbaseTX(address, "")
        genesisBlock.Transactions = append(genesisBlock.Transactions, coinbaseTx)
    }
	// Create the blockchain with the genesis block
	blockchain := &Blockchain{
		GHash: genesisBlock.Hash,
		Chain: []Block{*genesisBlock},
	}

    return blockchain
}

// creates a new blockchain given a valid genesis block
func NewBlockchainFromGB(genesis *Block) *Blockchain {
	blockchain := &Blockchain{
		GHash: genesis.Hash,
		Chain: []Block{*genesis},
	}
	return blockchain
}

func (bc *Blockchain) GetBalance(address string) int {
	var unspentOutputs []CustomStruct;
	balance := 0
	for _, b := range bc.Chain {
		for _, t := range b.Transactions{
			for i, out := range t.TxOuts{
				if out.ScriptPubKey == address {
					// Add unspent output to the list
					unspentOutputs = append(unspentOutputs, CustomStruct{
						Tid:           t.Hash,
						Index:      i,
						TXOutputValue: out,
					})
				}
			}
			for _, tin := range t.TxIns{
				for j, unspent := range unspentOutputs {
					if EqualSlices(tin.Txid, unspent.Tid) && tin.Vout == unspent.Index {
						// Remove spent output from the list
						unspentOutputs = append(unspentOutputs[:j], unspentOutputs[j+1:]...)
					}
				}

			}
		}
	}
	for _, u := range(unspentOutputs){
		balance += u.TXOutputValue.Value;
	}
	return balance
}
