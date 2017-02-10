package main

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/CebEcloudTime/charitycc/protos"
	"github.com/CebEcloudTime/charitycc/utils"
)

func main() {
	donorAddr := "donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1"

	tx := genTx(donorAddr)

	json, _ := json.Marshal(tx)

	fmt.Println("============== json")
	fmt.Println(string(json))

	base64String := base64.StdEncoding.EncodeToString(json)

	fmt.Println("============== base64String")
	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], fund01private)

	fmt.Println("============== sign")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))

}

func genTx(donorAddr string) protos.TX {
	var tx protos.TX

	tx.Version = 170101
	tx.Timestamp = time.Now().UTC().Unix()

	tx.Txin = genTxin()

	tx.Txout = genTxout(donorAddr)

	tx.InputData = "donoruuid"

	tx.Founder = "fund01"

	return tx
}

func genTxin() []*protos.TX_TXIN {
	var txins []*protos.TX_TXIN
	var txin1 protos.TX_TXIN
	txin1.Addr = "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	txin1.SourceTxHash = "ed21c857da821e827380c8faafa99d94d4102a08ff1a7e64b05183c42f6db228"
	txin1.Idx = 0

	txins = append(txins, &txin1)

	return txins
}

func genTxout(donorAddr string) []*protos.TX_TXOUT {
	var txouts []*protos.TX_TXOUT

	rechangeTxout := genRechangeTxout(donorAddr)
	txouts = append(txouts, &rechangeTxout)

	contractTxout := genContractTxout(donorAddr)
	txouts = append(txouts, &contractTxout)

	return txouts
}

func genRechangeTxout(donorAddr string) protos.TX_TXOUT {

	var txout protos.TX_TXOUT

	txout.Addr = "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	txout.Value = 1000 * 100 * 95

	txout.Attr = donorAddr

	txDataInfo := fmt.Sprintf("%s%d", txout.Addr, txout.Value)
	hashed := sha256.Sum256([]byte(txDataInfo))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], smartcontract01private)

	txout.Sign = base64.StdEncoding.EncodeToString(_sign)

	return txout
}

func genContractTxout(donorAddr string) protos.TX_TXOUT {

	var txout protos.TX_TXOUT

	txout.Addr = "bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"
	txout.Value = 1000 * 100 * 900

	txout.Attr = donorAddr

	txDataInfo := fmt.Sprintf("%s%d", txout.Addr, txout.Value)
	hashed := sha256.Sum256([]byte(txDataInfo))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], bargain01private)

	txout.Sign = base64.StdEncoding.EncodeToString(_sign)

	return txout
}

var fund01private, bargain01private, smartcontract01private []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	fund01private, err = ioutil.ReadFile(filepath + "fund01private.pem")
	if err != nil {
		os.Exit(-1)
	}

	bargain01private, err = ioutil.ReadFile(filepath + "bargain01private.pem")
	if err != nil {
		os.Exit(-1)
	}

	smartcontract01private, err = ioutil.ReadFile(filepath + "smartcontract01private.pem")
	if err != nil {
		os.Exit(-1)
	}

}
