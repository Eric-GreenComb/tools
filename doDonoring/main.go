package main

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/CebEcloudTime/charitycc/utils"
)

func main() {

	donorUUID := "donorUUID"
	donorAddr := "donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1"
	smartContractAddr := "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	amount := "100000000"
	bankAddr := "cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924"
	channelAddr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fundAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"

	hashed := sha256.Sum256([]byte(donorUUID + donorAddr + smartContractAddr + amount + bankAddr + channelAddr + fundAddr))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], donor01private)

	fmt.Println("============== sign")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))

}

var donor01private []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	donor01private, err = ioutil.ReadFile(filepath + "donor01private.pem")
	if err != nil {
		os.Exit(-1)
	}

}
