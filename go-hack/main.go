package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/rand"
	"sync"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"os"
	"time"
//	"encoding/hex"
)

var partitions = int(7)
var goRequest = gorequest.New()


func main() {
	var wg sync.WaitGroup
	for i := 0; i < partitions; i++ {
		wg.Add(1)
		addr := generateSeedAddress()
		log.Printf("Seed addr: %v\n", addr)
		go generateAddresses(addr)
		//generateAddresses(addr)
	}
	wg.Wait()
}
 
func generateSeedAddress() []byte {
	privKey := make([]byte, 32)
	for i := 0; i < 32; i++ {
		privKey[i] = byte(rand.Intn(256))
	}
	return privKey
}

func generateAddresses(seedPrivKey []byte) {
	for ; ; {
		incrementPrivKey(seedPrivKey)
		priv := convertToPrivateKey(seedPrivKey)
		myethscan := "https://api.etherscan.io/api?module=account&action=balance&address="
		address := crypto.PubkeyToAddress(priv.PublicKey)
		tail :="&tag=latest&apikey="
		myethscankey := "AIZN4W3JXM3PGKU3KBAAYXMJ9FFKHJCGA5"

		finalstring := fmt.Sprintf("%s%s%s%s",myethscan , address.Hex() , tail , myethscankey);
		//log.Printf("%s", finalstring)

		resp, err := http.Get(finalstring)
		if err != nil {
				// handle error
				fmt.Println("get error", err)
				return;
			}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("body error", err)
		}
		//fmt.Println(string(body))
		replace := strings.Replace(string(body),"{","",-1)
		replace = strings.Replace(replace,"}","",-1)
		replace = strings.Replace(replace,`"`,"",-1)
		split := strings.Split(replace,":")
		result, _ := strconv.Atoi(split[3])
			log.Printf("Found address with ETH balance, priv: %s, addr: %s, ammount %d", priv.D, address.Hex(),result )
	//	result /= 1000000000000000000;

		if(result > 0) {
			writeToFound(fmt.Sprintf("Found address with ETH balance, priv: %s, addr: %s, ammount %d\n", priv.D, address.Hex(),result ))
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func writeToFound(text string) {
	foundFileName := "./found.txt"
	if _, err := os.Stat(foundFileName); os.IsNotExist(err) {
		_, _ = os.Create(foundFileName)
	}
	f, err := os.OpenFile(foundFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer f.Close()
	if err != nil {
		log.Printf(err.Error())
	}
	_, err = f.WriteString(text)
	if err != nil {
		log.Printf(err.Error())
	}
}

func incrementPrivKey(privKey []byte) {
	/*
	for i := 31; i > 0; i-- {
		if privKey[i]+1 == 255 {
			privKey[i] = 0
		} else {
			privKey[i] += 1
			break
		}
	}
	*/
	i := rand.Intn(32)
		if privKey[i]+1 == 255 {
			privKey[i] = 0
		} else {
			privKey[i] += 1
		}
}

func convertToPrivateKey(privKey []byte) (*ecdsa.PrivateKey) {
	return crypto.ToECDSAUnsafe(privKey)
}

