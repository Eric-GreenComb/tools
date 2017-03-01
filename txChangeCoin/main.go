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

	fmt.Println("============== base64String")
	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], cebbankprivate)

	fmt.Println("============== sign")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))

}

func genTx() protos.TX {
	var tx protos.TX

	tx.Version = 170101
	tx.Timestamp = time.Now().UTC().Unix()

	tx.Txin = genTxin()

	tx.Txout = genTxout()

	tx.InputData = "donoruuid"

	tx.Founder = "eric"

	return tx
}

func genTxin() []*protos.TX_TXIN {
	var txins []*protos.TX_TXIN
	var txin protos.TX_TXIN
	txin.Addr = "cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924"
	txin.SourceTxHash = "aaaaa"
	txin.Idx = 0

	txins = append(txins, &txin)

	return txins
}

func genTxout() []*protos.TX_TXOUT {
	var txouts []*protos.TX_TXOUT

	changeCoinTxout := genChangeCoinTxout()
	txouts = append(txouts, &changeCoinTxout)

	return txouts
}

func genRechangeTxout() protos.TX_TXOUT {
	var txout protos.TX_TXOUT

	txout.Value = (100000000 - 1000) * 100 * 1000
	txout.Addr = "cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924"

	txDataInfo := fmt.Sprintf("%s%d", txout.Addr, txout.Value)
	hashed := sha256.Sum256([]byte(txDataInfo))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], cebbankprivate)

	txout.Sign = base64.StdEncoding.EncodeToString(_sign)

	return txout
}

func genChangeCoinTxout() protos.TX_TXOUT {
	var txout protos.TX_TXOUT

	txout.Value = 1000 * 100 * 1000
	txout.Addr = "donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1"
	txout.Attr = "donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1,donorUUID"

	return txout
}

var cebbankprivate, donor01private []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	cebbankprivate, err = ioutil.ReadFile(filepath + "cebbankprivate.pem")
	if err != nil {
		os.Exit(-1)
	}

	donor01private, err = ioutil.ReadFile(filepath + "donor01private.pem")
	if err != nil {
		os.Exit(-1)
	}

}
