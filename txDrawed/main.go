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
	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("============== foundationAddr")
	fmt.Println(foundationAddr)

	drawUUID := "68926CF4F6EE035A2DC2E0B606D012A2"
	fmt.Println("============== drawUUID")
	fmt.Println(drawUUID)

	smartContractAddr := "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	fmt.Println("============== smartContractAddr")
	fmt.Println(smartContractAddr)

	bargainAddr := "bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"
	fmt.Println("============== bargainAddr")
	fmt.Println(bargainAddr)

	amount := "90000000"
	fmt.Println("============== amount")
	fmt.Println(amount)

	txoutAttr := "donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1,15F332E9906ED10294CC634747ADD787"

	tx := genTx(drawUUID, txoutAttr, foundationAddr)

	json, _ := json.Marshal(tx)

	base64String := base64.StdEncoding.EncodeToString(json)

	fmt.Println("============== base64String")
	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(drawUUID + smartContractAddr + bargainAddr + amount + base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], fund01private)

	fmt.Println("============== sign")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))

}

func genTx(drawUUID, txoutAttr, foundationAddr string) protos.TX {
	var tx protos.TX

	tx.Version = 170101
	tx.Timestamp = time.Now().UTC().Unix()

	tx.Txin = genTxin()

	tx.Txout = genTxout(txoutAttr)

	tx.InputData = drawUUID

	tx.Founder = foundationAddr

	return tx
}

func genTxin() []*protos.TX_TXIN {
	var txins []*protos.TX_TXIN
	var txin1 protos.TX_TXIN
	txin1.Addr = "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	txin1.SourceTxHash = "460247fc1bbf377dac34694132fa907b7fc0955264266d898d49b2771b6d1497"
	txin1.Idx = 0

	txins = append(txins, &txin1)

	return txins
}

func genTxout(txoutAttr string) []*protos.TX_TXOUT {
	var txouts []*protos.TX_TXOUT

	rechangeTxout := genRechangeTxout(txoutAttr)
	txouts = append(txouts, &rechangeTxout)

	contractTxout := genContractTxout(txoutAttr)
	txouts = append(txouts, &contractTxout)

	return txouts
}

func genRechangeTxout(txoutAttr string) protos.TX_TXOUT {

	var txout protos.TX_TXOUT

	txout.Addr = "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	txout.Value = 1000 * 100 * 95

	txout.Attr = txoutAttr

	txDataInfo := fmt.Sprintf("%s%d", txout.Addr, txout.Value)
	hashed := sha256.Sum256([]byte(txDataInfo))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], smartcontract01private)

	txout.Sign = base64.StdEncoding.EncodeToString(_sign)

	return txout
}

func genContractTxout(txoutAttr string) protos.TX_TXOUT {

	var txout protos.TX_TXOUT

	txout.Addr = "bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"
	txout.Value = 1000 * 100 * 900

	txout.Attr = txoutAttr

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
