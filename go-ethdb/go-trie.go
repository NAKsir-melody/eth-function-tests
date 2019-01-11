package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
)

func main() {

	ldb, err := ethdb.NewLDBDatabase("chaindata", 0, 0);
	if err != nil {
		fmt.Printf("open error\n" + err.Error())
		return
	}
	hash := string("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")
	test_trie, _ := trie.New(common.HexToHash(hash),trie.NewDatabase(ldb))
	if test_trie != nil {
		fmt.Printf("trie error\n")
	}
	it := trie.NewIterator(test_trie.NodeIterator(nil))
	for it.Next() {
		fmt.Printf(string(it.Key) + ", " + string(it.Value))
	}
}

