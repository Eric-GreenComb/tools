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
	tx := genTx()

	json, _ := json.Marshal(tx)

	fmt.Println(string(json))

	base64String := base64.StdEncoding.EncodeToString(json)

	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], cebbankprivate)

	fmt.Println(base64.StdEncoding.EncodeToString(_sign))

}

func genTx() protos.TX {
	var tx protos.TX

	tx.Version = 170101
	tx.Timestamp = time.Now().UTC().Unix()

	tx.Txin = genTxin()

	tx.Txout = genTxout()

	tx.InputData = "donoruuid"

	tx.Founder = "cebbank"

	return tx
}

func genTxin() []*protos.TX_TXIN {
	var txins []*protos.TX_TXIN
	var txin protos.TX_TXIN
	txin.Addr = "null"
	txin.SourceTxHash = "null"
	txin.Idx = 0

	txins = append(txins, &txin)

	return txins
}

func genTxout() []*protos.TX_TXOUT {
	var txouts []*protos.TX_TXOUT
	var txout protos.TX_TXOUT

	txout.Value = 10000000 * 100 * 1000
	txout.Addr = "cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924"

	txDataInfo := fmt.Sprintf("%s%d", txout.Addr, txout.Value)
	hashed := sha256.Sum256([]byte(txDataInfo))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], cebbankprivate)

	txout.Sign = base64.StdEncoding.EncodeToString(_sign)

	txouts = append(txouts, &txout)

	return txouts
}

var cebbankprivate, cebbankpublic []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	cebbankprivate, err = ioutil.ReadFile(filepath + "cebbankprivate.pem")
	if err != nil {
		os.Exit(-1)
	}

	cebbankpublic, err = ioutil.ReadFile(filepath + "cebbankpublic.pem")
	if err != nil {
		os.Exit(-1)
	}

}
