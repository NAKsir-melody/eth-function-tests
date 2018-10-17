package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"fmt"
	"time"
)

func main() {
	ScanBlock()
}
// major
func ScanBlock() {
	ctx := context.Background();
	// main net
	client, err := ethclient.DialContext(ctx, "https://mainnet.infura.io/v3/d023760061574ff7b403904d70ee0d55")

	// test net
	//client, err := ethclient.DialContext(ctx, "https://ropsten.infura.io/v3/d023760061574ff7b403904d70ee0d55")

	// private net
	//client, err := ethclient.DialContext(ctx, "http://127.0.0.1:8080")
	if err != nil {
		fmt.Println("rpc conn error")
		return
	}
	defer client.Close()

	var nBlock uint64
	nBlock = 0
	for ; ; {
		//b_num := big.NewInt(1)
		//block, err := client.BlockByNumber(ctx,b_num)
		block, err := client.BlockByNumber(ctx,nil)
		if err != nil {
			fmt.Println("block get failed")
			return
		}
		bn := block.NumberU64();
		if(nBlock != bn) {
			nBlock = bn;
			fmt.Println(nBlock)
			//fmt.Println(block.Time())
			//fmt.Println(block.Difficulty())
			//block_hash := block.Hash()
			block_body := block.Body()

			for no, tx := range block_body.Transactions {
				fmt.Println("==============")
				fmt.Println(no ,tx.Hash().Hex())
				from, err := client.TransactionSender(ctx, tx, block.Hash(),uint(no))
				if(err == nil){
					fmt.Println("from: ",from.Hex())
				}
				if tx.To() != nil {
					fmt.Println("to: ",tx.To().Hex())
				} else {
					fmt.Println("to: 0 - contract creation")
				}
				result := tx.Value()
				fmt.Println("value: ",result)
				fmt.Println("data: ",tx.Data())
				time.Sleep(50 * time.Millisecond)
			}
		}
		time.Sleep(1000 * time.Millisecond)
	}

}

