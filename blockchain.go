package main

type BlockChain struct {
	blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func CumGoBrr() *Block {
	return NewBlock("This is Power", []byte{})
}

func NewBlockChain() *BlockChain { //JesusWeptAsThereWereNoWorldsToConqure
	return &BlockChain{[]*Block{CumGoBrr()}}
}
