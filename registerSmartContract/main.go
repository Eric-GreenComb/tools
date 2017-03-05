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
	genSmartContract01()
	genSmartContract02()
	genSmartContract03()
	genSmartContract04()
	genSmartContract05()
}

func genSmartContract01() {
	fmt.Println("")
	fmt.Println("=========== SmartContract01 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	smartContractAddr := "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	fmt.Println("=========== smartContractAddr ==========")
	fmt.Println(smartContractAddr)

	smartContractPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEMG80Q3dIdE9ISzEvUXZZQzlCL0o4NTZmKwoyQTFtRGE5elpiTE5yZXhHdVlPZ1Yvb1UwTDdTaVBnSUF3Tm9HSGRENmtQakxkK2N3aEVaMHdQRnluZjhicDh5CithK00wRUJjeStkK1VQQ1MrQmNlYjdrWDVOK3VYUm9za3VzWWx3MTVRc3UyeDBrZ3hXVUJIUXZiOGxEV0xuNVUKckptSVVUdG5XWkExMHB0em13SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== smartContractPublickKey ==========")
	fmt.Println(smartContractPublickKey)

	smartContract := initSmartContract01()

	jsonSmartContract, _ := json.Marshal(smartContract)

	base64String := base64.StdEncoding.EncodeToString(jsonSmartContract)

	fmt.Println("=========== data : ")
	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(smartContractAddr + smartContractPublickKey + base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))
}

func initSmartContract01() protos.SmartContract {
	var treaty protos.SmartContract

	treaty.Addr = "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	treaty.Name = "宁夏母亲水窖"
	treaty.Detail = `宁夏母亲水窖`
	treaty.Goal = 10000000 * 100 * 1000
	treaty.PartyA = "中国妇女发展基金会"
	treaty.PartyB = "宁夏母亲水窖项目"
	treaty.Status = 1
	treaty.Foundation = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.ChannelAddr = "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	treaty.ChannelName = "瑶瑶缴费"
	treaty.ChannelFee = 10
	treaty.FundAddr = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.FundName = "中国妇女发展基金会"
	treaty.FundManangerFee = 72
	treaty.CreateTimestamp = 1483228800
	treaty.UtilTimestamp = 1483228800
	treaty.TerminationTimestamp = 1609459200

	return treaty
}

func genSmartContract02() {
	fmt.Println("")
	fmt.Println("=========== SmartContract02 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	smartContractAddr := "smartcontract02:f14e694150c33750690d2c3baad7bb3406799263f7de5920a088072da5797800"
	fmt.Println("=========== smartContractAddr ==========")
	fmt.Println(smartContractAddr)

	smartContractPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERUVHZnBDNUxJSnJBejh1ZUhHMjlma1MxUwpFTEpEdFJLZjBPM0E0cjZ0enNRNjlJa0pFc0h6b1V4N1dPVlBtQmJsWkFOZ1Y0VlJ1UVU4M3RSUFUwK2xVZ0dpCmhoNy9TWE9OWjZMb0ozR3Q2bUVFTUJ1Zk92OFVZMEN2a2QyWjVzSjREOW5IbnVWT1JpeS83MDdVZVBVbEkxRnUKeFd4LzBTa1BrVlhKbDBNYkJ3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== smartContractPublickKey ==========")
	fmt.Println(smartContractPublickKey)

	smartContract := initSmartContract02()

	jsonSmartContract, _ := json.Marshal(smartContract)

	base64String := base64.StdEncoding.EncodeToString(jsonSmartContract)

	fmt.Println("=========== data : ")
	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(smartContractAddr + smartContractPublickKey + base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))
}

func initSmartContract02() protos.SmartContract {
	var treaty protos.SmartContract

	treaty.Addr = "smartcontract02:f14e694150c33750690d2c3baad7bb3406799263f7de5920a088072da5797800"
	treaty.Name = "河北母亲水窖"
	treaty.Detail = `河北母亲水窖`
	treaty.Goal = 10000000 * 100 * 1000
	treaty.PartyA = "中国妇女发展基金会"
	treaty.PartyB = "河北母亲水窖项目"
	treaty.Status = 1
	treaty.Foundation = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.ChannelAddr = "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	treaty.ChannelName = "瑶瑶缴费"
	treaty.ChannelFee = 30
	treaty.FundAddr = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.FundName = "中国妇女发展基金会"
	treaty.FundManangerFee = 64
	treaty.CreateTimestamp = 1483228800
	treaty.UtilTimestamp = 1483228800
	treaty.TerminationTimestamp = 1609459200

	return treaty
}

func genSmartContract03() {
	fmt.Println("")
	fmt.Println("=========== SmartContract03 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	smartContractAddr := "smartcontract03:9564142592ce9c66edb005f8098a725f72325d8a0411028ab78911f5d6c14718"
	fmt.Println("=========== smartContractAddr ==========")
	fmt.Println(smartContractAddr)

	smartContractPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDY3NjbzIxdnUzRWZmeG0vN21NWlZxak44cQpIV1k3aWFMNndvZm82eHJtcWxjaTZac21DTXFYSVdRMGN3MTdIcmlWWEl5anM4NlkyOEtFMVpreko1RnYrVUJOCm1KVndvZmVHQjVlSHNUbDNWK01tam80M1B0TUZ5ZGZqeFJzNW5JQmZYZ3ZyNHVHbjFmcDdaL0FteUw1WDhzYmUKM0FrUlRJZFlXTXRZOURUa2RRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== smartContractPublickKey ==========")
	fmt.Println(smartContractPublickKey)

	smartContract := initSmartContract03()

	jsonSmartContract, _ := json.Marshal(smartContract)

	base64String := base64.StdEncoding.EncodeToString(jsonSmartContract)

	fmt.Println("=========== data : ")
	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(smartContractAddr + smartContractPublickKey + base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))
}

func initSmartContract03() protos.SmartContract {
	var treaty protos.SmartContract

	treaty.Addr = "smartcontract03:9564142592ce9c66edb005f8098a725f72325d8a0411028ab78911f5d6c14718"
	treaty.Name = "甘肃母亲水窖"
	treaty.Detail = `甘肃母亲水窖`
	treaty.Goal = 10000000 * 100 * 1000
	treaty.PartyA = "中国妇女发展基金会"
	treaty.PartyB = "甘肃母亲水窖项目"
	treaty.Status = 1
	treaty.Foundation = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.ChannelAddr = "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	treaty.ChannelName = "瑶瑶缴费"
	treaty.ChannelFee = 30
	treaty.FundAddr = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.FundName = "中国妇女发展基金会"
	treaty.FundManangerFee = 64
	treaty.CreateTimestamp = 1483228800
	treaty.UtilTimestamp = 1483228800
	treaty.TerminationTimestamp = 1609459200

	return treaty
}

func genSmartContract04() {
	fmt.Println("")
	fmt.Println("=========== SmartContract04 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	smartContractAddr := "smartcontract04:facd57477f723c07a85296c20f53ebad70c8ac4edbd67094ea3121bc8d1e8994"
	fmt.Println("=========== smartContractAddr ==========")
	fmt.Println(smartContractAddr)

	smartContractPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDMmtleWxua0c1T3ZBSlhwL3BKdlFJQzJqawpkQTNIc0lsK3FJSjArczNHczdSS2pwSWYyaEs4R1puWkhKSWpiM1V1SGN3dmZudWdZUE1XWThJYm1RYXBUbFFlCjhBVUhITWZNZjljbS9xQjVuZXlXNVdNQmhyc3lVdTY4T25aUFh3Mm1YWjMvcUN2WGkyUG54VWZzQXJBQWUyRkEKU1dRcnVEUXA4RjlSNVd0T2h3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== smartContractPublickKey ==========")
	fmt.Println(smartContractPublickKey)

	smartContract := initSmartContract04()

	jsonSmartContract, _ := json.Marshal(smartContract)

	base64String := base64.StdEncoding.EncodeToString(jsonSmartContract)

	fmt.Println("=========== data : ")
	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(smartContractAddr + smartContractPublickKey + base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))
}

func initSmartContract04() protos.SmartContract {
	var treaty protos.SmartContract

	treaty.Addr = "smartcontract04:facd57477f723c07a85296c20f53ebad70c8ac4edbd67094ea3121bc8d1e8994"
	treaty.Name = "陕西母亲水窖"
	treaty.Detail = `陕西母亲水窖`
	treaty.Goal = 10000000 * 100 * 1000
	treaty.PartyA = "中国妇女发展基金会"
	treaty.PartyB = "陕西母亲水窖项目"
	treaty.Status = 0
	treaty.Foundation = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.ChannelAddr = "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	treaty.ChannelName = "瑶瑶缴费"
	treaty.ChannelFee = 30
	treaty.FundAddr = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.FundName = "中国妇女发展基金会"
	treaty.FundManangerFee = 64
	treaty.CreateTimestamp = 1483228800
	treaty.UtilTimestamp = 1483228800
	treaty.TerminationTimestamp = 1609459200

	return treaty
}

func genSmartContract05() {
	fmt.Println("")
	fmt.Println("=========== SmartContract05 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	smartContractAddr := "smartcontract05:5e8f644f5d79ad3d515f86ac545bf665595c9558e236b3c87965d4ff122a8286"
	fmt.Println("=========== smartContractAddr ==========")
	fmt.Println(smartContractAddr)

	smartContractPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDaVk5bzhSaS90Y3EvTGpLZW5JL0Z5MnVYZApyVVlPc3VJc3pOTWVaY1VtK3V4dXZDUHRlSWFlT0h2eU1QN3hoclR3cXVMODZ1MjdselRwZENVK0cxeC9wRkMyCjJTeHRWTVRESUlUR3JDTmY1V1NyK3hZTFhleDVIUTlCcmNJeG9wT0ZydXhSSGNSRkc0cmg1MEl6SXBlcTRrdVQKajg0YmRMdWlKNG03S2pxNTV3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== smartContractPublickKey ==========")
	fmt.Println(smartContractPublickKey)

	smartContract := initSmartContract05()

	jsonSmartContract, _ := json.Marshal(smartContract)

	base64String := base64.StdEncoding.EncodeToString(jsonSmartContract)

	fmt.Println("=========== data : ")
	fmt.Println(base64String)

	hashed := sha256.Sum256([]byte(smartContractAddr + smartContractPublickKey + base64String))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))
}

func initSmartContract05() protos.SmartContract {
	var treaty protos.SmartContract

	treaty.Addr = "smartcontract05:5e8f644f5d79ad3d515f86ac545bf665595c9558e236b3c87965d4ff122a8286"
	treaty.Name = "西藏母亲水窖"
	treaty.Detail = `西藏母亲水窖`
	treaty.Goal = 10000000 * 100 * 1000
	treaty.PartyA = "中国妇女发展基金会"
	treaty.PartyB = "西藏母亲水窖项目"
	treaty.Status = 0
	treaty.Foundation = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.ChannelAddr = "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	treaty.ChannelName = "瑶瑶缴费"
	treaty.ChannelFee = 30
	treaty.FundAddr = "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	treaty.FundName = "中国妇女发展基金会"
	treaty.FundManangerFee = 64
	treaty.CreateTimestamp = 1483228800
	treaty.UtilTimestamp = 1483228800
	treaty.TerminationTimestamp = 1609459200

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
