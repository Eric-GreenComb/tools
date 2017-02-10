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

	fmt.Println("=========== SmartContract ==========")
	smartContract := genSmartContract()

	jsonSmartContract, _ := json.Marshal(smartContract)

	base64String := base64.StdEncoding.EncodeToString(jsonSmartContract)

	fmt.Println("=========== data : ")
	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))

	fmt.Println("=========== Bargain ==========")
	bargain := genBargain()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
}

func genSmartContract() protos.SmartContract {
	var treaty protos.SmartContract

	treaty.Addr = "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	treaty.Name = "宁夏西部地区母亲水窖项目"
	treaty.Detail = "宁夏西部地区母亲水窖项目"
	treaty.Goal = 10000000 * 100 * 1000
	treaty.PartyA = "某基金会"
	treaty.PartyB = "某地区"
	treaty.Status = 0
	treaty.Foundation = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.ChannelFee = 2
	treaty.FundManangerFee = 3
	treaty.CreateTimestamp = time.Now().UTC().Unix()

	return treaty
}

func genBargain() protos.Bargain {
	var contract protos.Bargain

	contract.Addr = "bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"
	contract.Name = "宁夏西部地区母亲水窖项目XX县XX村水窖"
	contract.Detail = "宁夏西部地区母亲水窖项目XX县XX村水窖"
	contract.StartTime = "2017-01-01"
	contract.EndTime = "2018-01-01"
	contract.PartyA = "某基金会"
	contract.PartyB = "某地区XX县XX村施工队"
	contract.DepositBank = "光大银行"
	contract.BankAccount = "1298hfak09kkljadskf"

	return contract
}

var privateKey, publicKey []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	privateKey, err = ioutil.ReadFile(filepath + "fund01private.pem")
	if err != nil {
		os.Exit(-1)
	}

	publicKey, err = ioutil.ReadFile(filepath + "fund01public.pem")
	if err != nil {
		os.Exit(-1)
	}

}
