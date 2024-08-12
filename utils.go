package main

import (
	
)

// IntToHex converts an int64 to a byte array of length 8
// 
func IntToHex(num int64) []byte {
	//Create a byte array with 8 bytes
	b := make([]byte, 8)

    for i := 7; i >= 0; i-- {
		//assign the less important 8 bits to the current byte
        b[i] = byte(num & 0xFF) 
		// Right-shift by 8 bits to get the next set of bits
        num >>= 8
    }

	return b
}

// true if the hash starts with zeroBits zeros, note that the hash is
// a slice of *bytes* but we want zeroBits *bits* (a byte has 8 bits)

func StartsWithXZeros(hash []byte, zeroBits int) bool {
    j := 0
    var x byte
    for i := 0; i < len(hash); i++ {
        x = hash[i]
		//if it x is equal to zero that means all 8 bits are equal to zero
        if (x == 0) {
			//add 8 to the counter
			j=j+8;
        }else{
			for {
				if (x & 0x80) != 0 {
					//return true if the number of zero is more or equal to the demanded one
					if(j>=zeroBits){
						return true
					}
					//return false else
					return false
				}
				//add one to the counter j
				j++
				//left-shit the value x
				x <<= 1
			}
		}
    }

	if(j>=zeroBits){
		return true
	}
    return false
}


func EqualSlices(a, b []byte) bool {
	if len(a) != len(b){
		return false;
	}
	for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }
	return true
}

func EqualMaps(a, b map[string]int) bool {
	//check if have same length
	if len(a) != len(b) {
        return false
    }
	for keyA, valueA := range a {
        if valueB, ok := b[keyA]; !ok || valueA != valueB {
            return false
        }
    }
    return true
}

func EqualTransactions(a,b Transaction) bool{
	return EqualSlices(a.Hash,b.Hash)
}

func EqualBlocks(a,b Block) bool{
	return EqualSlices(a.Hash,b.Hash)
}

// Serializes a slice of byte slices by converting it to a byte slice so
// needed to easily hash data
func Serialize(input [][]byte )[]byte {
	var output []byte
	for _, o := range input {
		//append the unpacked elements of o to output
        output = append(output, o...)
    }
	return output
}

