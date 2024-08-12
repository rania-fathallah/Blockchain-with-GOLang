# Blockchain-with-GOLang

I independently developed a blockchain from scratch using GoLang. This hands-on experience provided an in-depth understanding of blockchain fundamentals. By implementing the technology myself, I gained a profound appreciation for the intricate mechanisms and complexities involved in creating a decentralized, immutable ledger. This project significantly enhanced my programming skills and solidified my knowledge of blockchain architecture.

## How This Project Work :
This section provides detailed explanations for each component of the code.

### block.go :


#### func NewBlock :
        Creating a new block with the choice to mine the block or not(meaning to set the Hash or not).The zeroBits parameter is
        the number of zeros the block Hash need to start with if mine is set to true.

#### func NewGBlock:
        Creating a block that get mined(mine is set to true) but that doesn't contain the hash of a previous block(it is set to
        []byte{}).The genesis block is the first block in the blockchain.

#### func IsCorrectlyHashed :
        Check if a block Hash start with a specific number of zeros at least.

#### func computeHash :
        Compute the hash of a block

#### func mine :
        Mine a block. The process is that for a set difficulty (number of zeros) we compute the hash of a block till we get a hash
        that start with the specific number of zeros or more. Each time we compute the hash but don't get a correct hash we increment
        the Nonce value in the block to get a different hash.


### utils.go :


#### func IntToHex :
        convert an int64(an int that has a max value of 8 bytes = 64 bits) to a []byte. For num for 8 times(the needed number of byte to reprensent an int64) which is the length of the array of bytes we get the first 8 bits = 1bit of the number(the ones with the less value) and we assign it
        to the current case of the bits array the we Rigth-Shift the number with 8 bits to have the next 8 bits as the less valuable one in the
        next iteration.

#### func StartsWithXZeros :
        This function first count the number of zeros a hash(a []byte) start with then return false if it is less than zeroBits and true else.
        For the length of the hash each time we select a byte "x" if it is equal 0 we just add 8 to the counter( since that mean that all
        8bits are equal to zero) else we loop inside the byte, each time we get the most value bit if it isn't equal to zero we just return
        true or false depending on the case else wa add 1 to the counter and we Left-Shift "x".

#### func EqualSlices :
        This function compare two []byte to see if they are equal or not. Wee first check if the have the same length if not return false.
        then for each byte we check if equal.

#### func EqualMaps :
            Check if two maps are equal. Check if they have the same length then check if if for each equal key the value is equal and each element exist in both maps.

#### func EqualTransactions :
        compare the hash of the two transactions. if equal hash then both equal.

#### func EqualBlocks :
        compare the hash of the two blocks.

#### func Serialize :
        In order for we append the elements of the arrays inside the [][]byte array to a []byte "output".


### transaction.go :


#### func NewCoinbaseTX :
        a coinbase transaction is a transaction you get yout first time in the blockchain in out code it gives you a gift of 10 bitcoins. it doesn't have the usual defined transaction inputs (it only got one transaction input that doesn't contain the hash of a transaction and an index for a transaction output : it is set to -1)  while for the transaction output it just states that person got 10 new bitcoins.


### blockchain.go :


#### func AddBlock :
        this function add a block to the end of the blockchain. it has the transactions as parametre and create a block that with mine set to true
        then append the created block to the block chain.

#### func NewTransfertTX:
        this function create a transaction(that is an amount give from one person to another) so we need to set the transation inputs and genrate the transaction outputs. We check the balance of the person sending the bitcoins if less than the amount to send an error arise. if it is more or equal we try to t=finc the array of the transaction input by finding the unspent outputs of the sender. We search for all transactions output
        of the sender then remove the ones used before as inputs. Create the transaction inputs from that list. If amount less that balance then
        another transaction output need to be added that return the amount exceeding the amount needed to the sender.

#### func NewBlockchain :
        Create a new blockchain. Create the genesis block that contain a set of coin base transactions that are created one for each address specified in the string addresses.

#### func NewBlockchainFromGB :
        Create a blockchain using an already defined genesis block.

#### func GetBalance :
        First we get the unspentOutputs list like we did in the NewTransfertTX function. Then we use it to add each of the values it has to the balance.


### CustomStruct struct :
    it is a struct used in the GetBalance function and the NewTransfertTX function. to help have a list containing the unspent
    outputs and the transaction hash that contain it and the index of this output in the TxOuts array.
