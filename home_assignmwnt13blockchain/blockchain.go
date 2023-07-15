package blockchain

import (
	"time"
)

//const DIFFICALTY = 2

type Blockchain struct {
	initialBock Block
	chain       []Block
	difficulty  int
}

func CreateBlockchain(difficulty int) Blockchain {
	initialBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}

	return Blockchain{
		initialBlock,
		[]Block{initialBlock},
		difficulty,
	}
}

func (b *Blockchain) AddBlock(from, to string, amount int) Block {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}

	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)

	return newBlock
}

func (b Blockchain) IsValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.generateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}

	return true
}
