package main

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/CebEcloudTime/charitycc/protos"
	"github.com/CebEcloudTime/charitycc/utils"
)

func main() {

	fmt.Println("=========== Bargain ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	bargainAddr := "bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"
	fmt.Println("=========== bargainAddr ==========")
	fmt.Println(bargainAddr)

	bargainPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDM3R4R3hnWUtKcjlZK21iMTROcENKS0NtUgo4QmNRek9MSjNEK3EvUFo1Zk9xUzUzdFhvVko2QUZtNEwyelZLYUFkMWNOS0s4L2t3RktsV1E5YmJLZ1ZOV25zCnc4MjM0N05yRzQxaWZocFZ3dThJVHJSOGlMOC9pR3lMdnh4SGg0OWpmQ3RIWkFHV3hrWkFsVDBwdkRZYTNJMW0KTjZMKytZWCt2WW42WTdOOFB3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== bargainPublickKey ==========")
	fmt.Println(bargainPublickKey)

	bargain := genBargain()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(bargainAddr + bargainPublickKey + base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
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

var privateKey []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	privateKey, err = ioutil.ReadFile(filepath + "fund01private.pem")
	if err != nil {
		os.Exit(-1)
	}

}
