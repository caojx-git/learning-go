package main

import (
	"strconv"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"fmt"
	"strings"
)

//区块结构
type Block struct {
	Hash          string
	PrevBlockHash string
	Timestamp     string
	Nonce         int
	Data          string
}

//区块链
type BlockChain struct {
	blocks []*Block
}

/*
	计算hash
 */
func (b *Block) calculateHash() string {
	record := strconv.Itoa(b.Nonce) + b.PrevBlockHash + b.Data + b.Timestamp;
	h := sha256.New()
	h.Write([]byte(record))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

/**
	创建区块
 */
func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{Timestamp: time.Now().String(), Data: data, PrevBlockHash: prevBlockHash}
	block.Hash = block.calculateHash()
	return block
}

/**
	添加区块
 */
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	fmt.Println("try add newBlock")
	newBlock.mineBlock(6)
	bc.blocks = append(bc.blocks, newBlock)
	fmt.Println("add newBlock success")

}

/*
	挖矿-工作量证明
 */
func (b *Block) mineBlock(difficulty int) {
	b.Nonce = 0
	for {
		hashstr := b.calculateHash()
		prefix := strings.Repeat("0", difficulty)
		if strings.HasPrefix(hashstr, prefix) {
			fmt.Println(hashstr+"--"+prefix)
			break
		}
		b.Nonce++
	}
}

func main() {
	//1.构建区块链
	bc := new(BlockChain)

	//2.创建创世区块
	block := NewBlock("Block创世区块", "")

	bc.blocks = append(bc.blocks, block)

	bc.AddBlock("区块2")
	bc.AddBlock("区块3")
	bc.AddBlock("区块4")
	bc.AddBlock("区块5")

	/*
		打印区块链
	 */
	for _, block := range bc.blocks {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Prev: %s\n", block.PrevBlockHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Println()
	}

}
