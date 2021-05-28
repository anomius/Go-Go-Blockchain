package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	TimeStamp    int64 //stores the timestamp of when the block was created
	Data         []byte
	PreBlockHash []byte
	Hash         []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
	headers := bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}
