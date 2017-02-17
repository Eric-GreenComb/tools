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
	donorUUID := "15F332E9906ED10294CC634747ADD787"
	smartContractAddr := "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"

	fmt.Println("============== donorAddr")
	fmt.Println(donorAddr)

	fmt.Println("============== donorUUID")
	fmt.Println(donorUUID)

	fmt.Println("============== smartContractAddr")
	fmt.Println(smartContractAddr)

	tx := genTx(donorAddr, donorUUID)

	json, _ := json.Marshal(tx)

	base64String := base64.StdEncoding.EncodeToString(json)

	fmt.Println("============== base64String")
	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(donorAddr + donorUUID + smartContractAddr + base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], donor01private)

	fmt.Println("============== sign")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))

}

func genTx(donorAddr, donorUUID string) protos.TX {
	var tx protos.TX

	tx.Version = 170101
	tx.Timestamp = time.Now().UTC().Unix()

	tx.Txin = genTxin(donorAddr)

	tx.Txout = genTxout(donorAddr, donorUUID)

	tx.InputData = donorUUID

	tx.Founder = "channel01"

	return tx
}

func genTxin(donorAddr string) []*protos.TX_TXIN {
	var txins []*protos.TX_TXIN
	var txin protos.TX_TXIN
	txin.Addr = donorAddr
	txin.SourceTxHash = "700e7f51da673619e602b5d45f0c30b10f03410181c569ebdf2583fdbeb11371"
	txin.Idx = 1

	txins = append(txins, &txin)

	return txins
}

func genTxout(donorAddr, donorUUID string) []*protos.TX_TXOUT {
	var txouts []*protos.TX_TXOUT

	smartContractTxout := genSmartContractTxout(donorAddr, donorUUID)
	txouts = append(txouts, &smartContractTxout)

	// channelTxout := genChannelTxout(donorAddr)
	// txouts = append(txouts, &channelTxout)

	// fundTxout := genFundTxout(donorAddr)
	// txouts = append(txouts, &fundTxout)

	return txouts
}

func genSmartContractTxout(donorAddr, donorUUID string) protos.TX_TXOUT {

	var txout protos.TX_TXOUT

	txout.Addr = "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	txout.Value = 1000 * 100 * 1000

	txout.Attr = donorAddr + "," + donorUUID

	return txout
}

func genChannelTxout(donorAddr string) protos.TX_TXOUT {

	var txout protos.TX_TXOUT

	txout.Addr = "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	txout.Value = 1000 * 100 * 2

	txout.Attr = donorAddr

	txDataInfo := fmt.Sprintf("%s%d", txout.Addr, txout.Value)
	hashed := sha256.Sum256([]byte(txDataInfo))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], channel01private)

	txout.Sign = base64.StdEncoding.EncodeToString(_sign)

	return txout
}

func genFundTxout(donorAddr string) protos.TX_TXOUT {

	var txout protos.TX_TXOUT

	txout.Addr = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	txout.Value = 1000 * 100 * 3

	txout.Attr = donorAddr

	txDataInfo := fmt.Sprintf("%s%d", txout.Addr, txout.Value)
	hashed := sha256.Sum256([]byte(txDataInfo))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], fund01private)

	txout.Sign = base64.StdEncoding.EncodeToString(_sign)

	return txout
}

var donor01private, channel01private, fund01private, smartcontract01private []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	donor01private, err = ioutil.ReadFile(filepath + "donor01private.pem")
	if err != nil {
		os.Exit(-1)
	}

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
