package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         map[string]interface{} `json:"data"`
	hash         string                 `json:"hash"`
	previousHash string                 `json:"previous_hash"`
	timestamp    time.Time              `json:"timestamp"`
	pow          int                    `json:"pow"`
}

func (b *Block) GetHash() string {
	return b.hash
}

func (b *Block) GetPow() int {
	return b.pow
}

func (b Block) generateHash() string {
	data, _ := json.Marshal(b.data)
	dataForHash := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	hash := sha256.Sum256([]byte(dataForHash))

	return fmt.Sprintf("%x", hash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.generateHash()
	}
}
