package main

import  (
  "bytes"
  "crypto/sha256"
)

type BlockChain struct {
  blocks []*Block
}

type Block struct {
  Hash      []byte
  Data      []byte
  PrevHash  []byte
}

func (b *Block) DeriveHash() {
  info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
  hash := sha256.Sum256(info)
  b.Hash = hash[:]
  // or b.Hash = hash
}

func CreateBlock(data string, prevHash []byte) *Block {
  block := &Block{[]byte{}, []byte(data), prevHash}
  block.DeriveHash()
  return block
}

func (chain *BlockChain) AddBlock(data string) {
  prevBlock := chain.blocks[len(chain.blocks)-1]
  new := CreateBlock(data, prevBlock.Hash)
  chain.blocks = append(chain.blocks, new)
}

func main()  {

}
