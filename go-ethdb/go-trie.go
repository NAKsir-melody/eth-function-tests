package main

import (
	"fmt"
	_ "encoding/hex"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/core/state"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/rlp"
	_ "github.com/ethereum/go-ethereum/accounts"
)

func main() {

	hash_list := []string{
		"0d9348243d7357c491e6a61f4b1305e77dc6acacdb8cc708e662f6a9bab6ca02",
		"f520abd5cf4fe1a16378bdf7d12fbabe6642a6f33996000e5763b39e15eca9bb",
		"a1a8122d87dcbe1634df20264274ed8f072e0eb3d7a608859689df9cb5f100d9",
	}

	addr_list := []string{
		"cea8f2236efa20c8fadeb9d66e398a6532cca6c8",
		"8e64566b5eb8f595f7eb2b8d302f2e5613cb8bae",
	}

	block_hash_list := []string{
		"4ab9d84e527f19e95ddc834c1765bec5ee3660c490f1ecc989a82db06432dc04",
		"705f309f574ba53ee2c1a00ac48dc417426ad0863e1106fc0c3cb49566c708b9",
		"bfca369fb07755806acf2a810a048ca6b2293c6b3ab05b6e1534f67120f7ee17",
		"e56fb7ef5f771834ad0dfb01081f830cdf34ea13f63a40c6cdc08fc9d9bb9f54",
	}
	// emptyState: c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470
	// emptyRoot:  56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421 

	ldb, err := ethdb.NewLDBDatabase("chaindata_tx", 0, 0);
	if err != nil {
		fmt.Printf("open error\n" + err.Error())
		return
	}

	trieDB := trie.NewDatabase(ldb)

	//trie
	state_trie, err := trie.NewSecure(common.HexToHash(hash_list[0]), trieDB, 0) //state
	if err != nil {
		fmt.Printf(err.Error() + "\n")
		return
	}

	it := trie.NewIterator(state_trie.NodeIterator(nil))
	for it.Next() {
		fmt.Printf("[%x, %x]\n", it.Key, it.Value)
	}

//state
	stateDB := state.NewDatabase(ldb)
	mystate, err := state.New(common.HexToHash(hash_list[0]), stateDB)
	if err != nil {
		fmt.Printf(err.Error() + "\n")
		return
	}

	addr := common.HexToAddress(addr_list[0])
	_ = common.HexToAddress(block_hash_list[3])
	if mystate.Exist(addr) {
		balance := mystate.GetBalance(addr)
		nonce := mystate.GetNonce(addr)
		fmt.Printf("%d\n",balance)
		fmt.Printf("%x\n",nonce)
	} else {
		fmt.Printf("no exist\n")
	}
}

