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

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	smartContractAddr := "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	fmt.Println("=========== smartContractAddr ==========")
	fmt.Println(smartContractAddr)

	smartContractPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEMG80Q3dIdE9ISzEvUXZZQzlCL0o4NTZmKwoyQTFtRGE5elpiTE5yZXhHdVlPZ1Yvb1UwTDdTaVBnSUF3Tm9HSGRENmtQakxkK2N3aEVaMHdQRnluZjhicDh5CithK00wRUJjeStkK1VQQ1MrQmNlYjdrWDVOK3VYUm9za3VzWWx3MTVRc3UyeDBrZ3hXVUJIUXZiOGxEV0xuNVUKckptSVVUdG5XWkExMHB0em13SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== smartContractPublickKey ==========")
	fmt.Println(smartContractPublickKey)

	smartContract := genSmartContract()

	jsonSmartContract, _ := json.Marshal(smartContract)

	base64String := base64.StdEncoding.EncodeToString(jsonSmartContract)

	fmt.Println("=========== data : ")
	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(smartContractAddr + smartContractPublickKey + base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))

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
	treaty.ChannelAddr = "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	treaty.ChannelName = "某ChannelName"
	treaty.ChannelFee = 2
	treaty.FundAddr = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.FundName = "某基金会"
	treaty.FundManangerFee = 3
	treaty.CreateTimestamp = time.Now().UTC().Unix()

	return treaty
}

var privateKey []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	privateKey, err = ioutil.ReadFile(filepath + "fund01private.pem")
	if err != nil {
		os.Exit(-1)
	}

}
