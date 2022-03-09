/*
 * @Author: your name
 * @Date: 2022-03-09 10:23:40
 * @LastEditTime: 2022-03-09 15:18:29
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_code/src/cryptogolang/transfer.go
 */
package main
 
import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
)
func main(){
	client,err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil{
		log.Fatal(err)
	}

	// privateKey, err := crypto.HexTOECDSA("20f85aa6bc340efe8c2ab6aea458533bbfa57ef7e464cd21bcf8b31977835724")
	privateKey, err := crypto.HexToECDSA("20f85aa6bc340efe8c2ab6aea458533bbfa57ef7e464cd21bcf8b31977835724")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce,err := client.PendingNonceAt(context.Background(),fromAddress)
	if err != nil{
		log.Fatal()
	}
	value := big.NewInt(5000000000000000000)
	gasLimit := uint64(210000)
	gasPrice,err := client.SuggestGasPrice(context.Background())
	if err != nil{
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x72995D75636b3c18f02A7e5fEA14AC5e93E76C53")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx,types.HomesteadSigner{} , privateKey)
	if err != nil{
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(),signedTx)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("Sent %s wei to %s :%s\n",value.String(),toAddress.Hex(),signedTx.Hash().Hex())
	// fmt.Printf()
}