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

	hashed := sha256.Sum256([]byte(donorAddr + base64String))
	_sign, _ := RsaSign(crypto.SHA256, hashed[:], channel01private)

	fmt.Println("============== sign")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))

}

func genTx(donorAddr string) TX {
	var tx TX

	tx.Version = 170101
	tx.Timestamp = time.Now().UTC().Unix()

	tx.Txin = genTxin(donorAddr)

	tx.Txout = genTxout(donorAddr)

	tx.InputData = "donoruuid"

	tx.Founder = "channel01"

	return tx
}

func genTxin(donorAddr string) []*TX_TXIN {
	var txins []*TX_TXIN
	var txin TX_TXIN
	txin.Addr = donorAddr
	txin.SourceTxHash = "17e0e8d20d130dc751cc5203db00e49b52568ce1edc3d8f6116a6b4cef78e08a"
	txin.Idx = 1

	txins = append(txins, &txin)

	return txins
}

func genTxout(donorAddr string) []*TX_TXOUT {
	var txouts []*TX_TXOUT

	smartContractTxout := genSmartContractTxout(donorAddr)
	txouts = append(txouts, &smartContractTxout)

	channelTxout := genChannelTxout(donorAddr)
	txouts = append(txouts, &channelTxout)

	fundTxout := genFundTxout(donorAddr)
	txouts = append(txouts, &fundTxout)

	return txouts
}

func genChannelTxout(donorAddr string) TX_TXOUT {

	var txout TX_TXOUT

	txout.Addr = "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	txout.Value = 1000 * 100 * 2

	txout.Attr = donorAddr

	txDataInfo := fmt.Sprintf("%s%d", txout.Addr, txout.Value)
	hashed := sha256.Sum256([]byte(txDataInfo))
	_sign, _ := RsaSign(crypto.SHA256, hashed[:], channel01private)

	txout.Sign = base64.StdEncoding.EncodeToString(_sign)

	return txout
}

func genFundTxout(donorAddr string) TX_TXOUT {

	var txout TX_TXOUT

	txout.Addr = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	txout.Value = 1000 * 100 * 3

	txout.Attr = donorAddr

	txDataInfo := fmt.Sprintf("%s%d", txout.Addr, txout.Value)
	hashed := sha256.Sum256([]byte(txDataInfo))
	_sign, _ := RsaSign(crypto.SHA256, hashed[:], fund01private)

	txout.Sign = base64.StdEncoding.EncodeToString(_sign)

	return txout
}

func genSmartContractTxout(donorAddr string) TX_TXOUT {

	var txout TX_TXOUT

	txout.Addr = "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	txout.Value = 1000 * 100 * 995

	txout.Attr = donorAddr

	txDataInfo := fmt.Sprintf("%s%d", txout.Addr, txout.Value)
	hashed := sha256.Sum256([]byte(txDataInfo))
	_sign, _ := RsaSign(crypto.SHA256, hashed[:], smartcontract01private)

	txout.Sign = base64.StdEncoding.EncodeToString(_sign)

	return txout
}

var channel01private, fund01private, smartcontract01private []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	channel01private, err = ioutil.ReadFile(filepath + "channel01private.pem")
	if err != nil {
		os.Exit(-1)
	}

	fund01private, err = ioutil.ReadFile(filepath + "fund01private.pem")
	if err != nil {
		os.Exit(-1)
	}

	smartcontract01private, err = ioutil.ReadFile(filepath + "smartcontract01private.pem")
	if err != nil {
		os.Exit(-1)
	}

}
