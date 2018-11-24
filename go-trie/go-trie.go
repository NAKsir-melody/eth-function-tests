package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
)

func main() {

	triedb := ethdb.NewMemDatabase()
	test_trie, _ := trie.New(common.Hash{},trie.NewDatabase(triedb))

	key1 := []byte{0xa,7,1,1,3,5,5}
	key2 := []byte{0xa,7,7,0xd,3,3,7}
	key3 := []byte{0xa,7,0xf,9,3,6,5}
	key4 := []byte{0xa,7,7,0xd,3,9,7}
	key5 := []byte{0xa,7,7,0xd,3}

	test_trie.Update(key1, []byte{1})
	test_trie.Update(key2, []byte{2})
	test_trie.Update(key3, []byte{3})
	test_trie.Update(key4, []byte{4})
	test_trie.Update(key5, []byte{5})

	result, _ := test_trie.TryGet(key1)
	fmt.Printf("%v\n", result)
	result, _ = test_trie.TryGet(key2)
	fmt.Printf("%v\n", result)
	result, _ = test_trie.TryGet(key3)
	fmt.Printf("%v\n", result)
	result, _ = test_trie.TryGet(key4)
	fmt.Printf("%v\n", result)
}

