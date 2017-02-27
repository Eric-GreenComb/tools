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
	genBargain01()
	genBargain02()
	genBargain03()
	genBargain04()
	genBargain05()
	genBargain06()
	genBargain07()
	genBargain08()
	genBargain09()
	genBargain10()
}

func genBargain01() {
	fmt.Println("")
	fmt.Println("=========== Bargain01 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	bargainAddr := "bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"
	fmt.Println("=========== bargainAddr ==========")
	fmt.Println(bargainAddr)

	bargainPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDM3R4R3hnWUtKcjlZK21iMTROcENKS0NtUgo4QmNRek9MSjNEK3EvUFo1Zk9xUzUzdFhvVko2QUZtNEwyelZLYUFkMWNOS0s4L2t3RktsV1E5YmJLZ1ZOV25zCnc4MjM0N05yRzQxaWZocFZ3dThJVHJSOGlMOC9pR3lMdnh4SGg0OWpmQ3RIWkFHV3hrWkFsVDBwdkRZYTNJMW0KTjZMKytZWCt2WW42WTdOOFB3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== bargainPublickKey ==========")
	fmt.Println(bargainPublickKey)

	bargain := initBargain01()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(bargainAddr + bargainPublickKey + base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
}

func initBargain01() protos.Bargain {
	var contract protos.Bargain

	contract.Addr = "bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"
	contract.Name = "宁夏母亲水窖项目一期施工"
	contract.Detail = "宁夏母亲水窖项目一期施工"
	contract.StartTime = "2017-01-01"
	contract.EndTime = "2017-03-01"
	contract.PartyA = "中国妇女发展基金会"
	contract.PartyB = "中国第三建筑公司"
	contract.DepositBank = "光大银行"
	contract.BankAccount = "6226660000000000"

	return contract
}

func genBargain02() {
	fmt.Println("")
	fmt.Println("=========== Bargain02 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	bargainAddr := "bargain02:527b016abdea151f014f3fd2f840f5a02b66ebde6ffc4172d2b2e50b4c46c423"
	fmt.Println("=========== bargainAddr ==========")
	fmt.Println(bargainAddr)

	bargainPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDN1lLUDRsMEdUMThzU0ZHcnU2VFhuTnNJZApYYkc1UUdLUjdzN2UrWEpQTWpDcGMzVkRzelhLMkRQTjhzYmNFbnFISEFRblR6aGt4V3NJVHB5MlYzTUtOZUZyCnRHc1J6amhYeTBqYmZFWmh0WDJHeTlaSXd0NFZnUlZwa0g4WGVEWDdKSStRNjVjbFZXa2l4MldBVTdhZ3dsM2gKRFdkbFZEUk13UVQrK3FwU3l3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== bargainPublickKey ==========")
	fmt.Println(bargainPublickKey)

	bargain := initBargain02()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(bargainAddr + bargainPublickKey + base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
}

func initBargain02() protos.Bargain {
	var contract protos.Bargain

	contract.Addr = "bargain02:527b016abdea151f014f3fd2f840f5a02b66ebde6ffc4172d2b2e50b4c46c423"
	contract.Name = "宁夏母亲水窖项目二期施工"
	contract.Detail = "宁夏母亲水窖项目二期施工"
	contract.StartTime = "2017-03-01"
	contract.EndTime = "2016-06-01"
	contract.PartyA = "中国妇女发展基金会"
	contract.PartyB = "中国第一建筑公司"
	contract.DepositBank = "光大银行"
	contract.BankAccount = "6226660000000000"

	return contract
}

func genBargain03() {
	fmt.Println("")
	fmt.Println("=========== Bargain03 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	bargainAddr := "bargain03:83a9c5019541d6b8f9170bfd59e79b77a217e3cab2989d6db706b0e08240ebb3"
	fmt.Println("=========== bargainAddr ==========")
	fmt.Println(bargainAddr)

	bargainPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEQ3lzVjZIZzJhNndBMTRzbVNQdzhRZnZYbgpkZkVHTVc0d25IT0thejZKc3hVTUZacitmYy9TV1d1TmZnVW0xU1NGNFhMRU1CWGk0dVpyY3lRVXhRQytydC8wCkpEUGhZZTZScWZlZzhWRHVxWmhuMEhEaG1uR2ZUZGU0Ni9QakJLdXkrNVJnM3ROWmhxQUoyejdBdkRqeXlKT3QKcDhkR3Jua2xpWHVsam11b0NRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== bargainPublickKey ==========")
	fmt.Println(bargainPublickKey)

	bargain := initBargain03()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(bargainAddr + bargainPublickKey + base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
}

func initBargain03() protos.Bargain {
	var contract protos.Bargain

	contract.Addr = "bargain03:83a9c5019541d6b8f9170bfd59e79b77a217e3cab2989d6db706b0e08240ebb3"
	contract.Name = "河北母亲水窖项目一期施工"
	contract.Detail = "河北母亲水窖项目一期施工"
	contract.StartTime = "2017-01-01"
	contract.EndTime = "2017-03-01"
	contract.PartyA = "中国妇女发展基金会"
	contract.PartyB = "中国第三建筑公司"
	contract.DepositBank = "光大银行"
	contract.BankAccount = "6226660000000000"

	return contract
}

func genBargain04() {
	fmt.Println("")
	fmt.Println("=========== Bargain04 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	bargainAddr := "bargain04:6a9eae00ae31d44ae7ad40cf2a66d7eee85ae624a1a737581fa7ddcf1dd20230"
	fmt.Println("=========== bargainAddr ==========")
	fmt.Println(bargainAddr)

	bargainPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEaXBzWG4xN0VzSC9HWjVxc0tSVEs4ajdBQgpJTnk2cXNQREovcWdNOUlmVEc3aURFaHJFbGlZQVFvT25taDJBN2lreEtkZ3F5U0EvY0xPTUhQaGsxaWs5QXNOCnBMRlI5K2wrZmxrazBoODY2ZGdzTU5zWG51eVRsbXN3UXorQ1NFWEE2RUN6aEI4cDFOWHVVNHdEQTdKK1VlMVkKUXVFQkVQcHVNTGx6bUoxSFZRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== bargainPublickKey ==========")
	fmt.Println(bargainPublickKey)

	bargain := initBargain04()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(bargainAddr + bargainPublickKey + base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
}

func initBargain04() protos.Bargain {
	var contract protos.Bargain

	contract.Addr = "bargain04:6a9eae00ae31d44ae7ad40cf2a66d7eee85ae624a1a737581fa7ddcf1dd20230"
	contract.Name = "河北母亲水窖项目二期施工"
	contract.Detail = "河北母亲水窖项目二期施工"
	contract.StartTime = "2017-03-01"
	contract.EndTime = "2017-06-01"
	contract.PartyA = "中国妇女发展基金会"
	contract.PartyB = "中国第一建筑公司"
	contract.DepositBank = "光大银行"
	contract.BankAccount = "6226660000000000"

	return contract
}

func genBargain05() {
	fmt.Println("")
	fmt.Println("=========== Bargain05 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	bargainAddr := "bargain05:0723e1bcdf3ec08dd94a83f50aab1114dc420c215fb869ae518ec755c320dab1"
	fmt.Println("=========== bargainAddr ==========")
	fmt.Println(bargainAddr)

	bargainPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEWmZ3STJkU1kwaW50QnBGNnlmMWVFYWc1NApKTENIUEJwRWdFYmdxY0ZzTUU2aW54NG9sRVVXRldQRDN0cDRwaE5USThDRXl6QVFmTHdNelJoT3dnQlJYOFBXClJuZ3dpU2MxRlMxelRQd2lHT09EbFJYM1ZxTnFQWk9aTnZMU2Z4TzFMNWVDai92Y1hjWExVRmZWUEFzb0VHdVUKZjlZd2hLWU00eks1SHFzUm5RSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== bargainPublickKey ==========")
	fmt.Println(bargainPublickKey)

	bargain := initBargain05()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(bargainAddr + bargainPublickKey + base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
}

func initBargain05() protos.Bargain {
	var contract protos.Bargain

	contract.Addr = "bargain05:0723e1bcdf3ec08dd94a83f50aab1114dc420c215fb869ae518ec755c320dab1"
	contract.Name = "陕西母亲水窖项目一期施工"
	contract.Detail = "陕西母亲水窖项目一期施工"
	contract.StartTime = "2017-01-01"
	contract.EndTime = "2017-03-01"
	contract.PartyA = "中国妇女发展基金会"
	contract.PartyB = "中国第三建筑公司"
	contract.DepositBank = "光大银行"
	contract.BankAccount = "6226660000000000"

	return contract
}

func genBargain06() {
	fmt.Println("")
	fmt.Println("=========== Bargain06 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	bargainAddr := "bargain06:990490e80e68525fc7c8cf276adec3087c261876f85934e7e553eaf0071ccb4f"
	fmt.Println("=========== bargainAddr ==========")
	fmt.Println(bargainAddr)

	bargainPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDM1JnV2phcnFYWnZzdUthQVFOd0syQ2UxegpKUW1lK2RDWXRJN0Zkb21nZXQrVWFhRTFUQVpLSjJRbDdtZ3J3WEVIS0xnZjB6V2FWRWdJSVlyaTJkaEE0L2YwCklJK1JuekZWSVpVblg4U0hFYk5kcE0vSHU0eEFLaTVyYkplajRLaFpsMU80Nmo2NEE1NUY1aURrR0xYcUNHSWMKcVdlTEp1N05SU1RHOXZocEVRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== bargainPublickKey ==========")
	fmt.Println(bargainPublickKey)

	bargain := initBargain06()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(bargainAddr + bargainPublickKey + base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
}

func initBargain06() protos.Bargain {
	var contract protos.Bargain

	contract.Addr = "bargain06:990490e80e68525fc7c8cf276adec3087c261876f85934e7e553eaf0071ccb4f"
	contract.Name = "陕西母亲水窖项目二期施工"
	contract.Detail = "陕西母亲水窖项目二期施工"
	contract.StartTime = "2017-01-01"
	contract.EndTime = "2017-03-01"
	contract.PartyA = "中国妇女发展基金会"
	contract.PartyB = "中国第一建筑公司"
	contract.DepositBank = "光大银行"
	contract.BankAccount = "6226660000000000"

	return contract
}

func genBargain07() {
	fmt.Println("")
	fmt.Println("=========== Bargain07 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	bargainAddr := "bargain07:765b5cf305978f41ec5a59e854f7b3f5eaa5a3889fd6b1c387f9447a2608d921"
	fmt.Println("=========== bargainAddr ==========")
	fmt.Println(bargainAddr)

	bargainPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDL2tkajJtLzVranNRS1o2MTVKdU1UUitmWApoc3pZa2FVc0N6Z1FRYjhNYmxoSUFLNWxHWFdQUjRleHltNHRZMkIzNUJZbGp1ZzBOb05MQy80aWI0cFNEbUhpClJnMWVlVXdNSHpNR01jZkZHeFlwSzQ5emh2ZEJQUkg5cERVL1JRZU8zZFB1REVLNWNQZUVwSTl2QnE1KzJwWG4KM1J4cUw4ZkpxdCtxQW8zVGFRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== bargainPublickKey ==========")
	fmt.Println(bargainPublickKey)

	bargain := initBargain07()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(bargainAddr + bargainPublickKey + base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
}

func initBargain07() protos.Bargain {
	var contract protos.Bargain

	contract.Addr = "bargain07:765b5cf305978f41ec5a59e854f7b3f5eaa5a3889fd6b1c387f9447a2608d921"
	contract.Name = "西藏母亲水窖项目一期施工"
	contract.Detail = "西藏母亲水窖项目一期施工"
	contract.StartTime = "2017-01-01"
	contract.EndTime = "2017-03-01"
	contract.PartyA = "中国妇女发展基金会"
	contract.PartyB = "中国第三建筑公司"
	contract.DepositBank = "光大银行"
	contract.BankAccount = "6226660000000000"

	return contract
}

func genBargain08() {
	fmt.Println("")
	fmt.Println("=========== Bargain08 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	bargainAddr := "bargain08:ea218af5f53e8e6ff824caab6fc521b5fa01fbce3525a46aa456cb63c884a522"
	fmt.Println("=========== bargainAddr ==========")
	fmt.Println(bargainAddr)

	bargainPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDazZFZlpKNmMyT2RLRlRKWnFINVdlcFYwRQpPSlpWK0U5NWtRN0lEdHpmN2t6WmN3c2R4cFZLN0xKdDlkT0grMDM5MHM2UFUwSjlnTFFNSi9XRlhQbGRDb09UCk9MUElndU5SS0FMRldtbnE1aysyUU1rQmJwNnVyNnVWUnNCOG5jL0s1NVBoQlJpV3B5TXo0M0FSSWs0aEdrc2wKNXBGdXU5NXRjYmhFU2k3SGNRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== bargainPublickKey ==========")
	fmt.Println(bargainPublickKey)

	bargain := initBargain08()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(bargainAddr + bargainPublickKey + base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
}

func initBargain08() protos.Bargain {
	var contract protos.Bargain

	contract.Addr = "bargain08:ea218af5f53e8e6ff824caab6fc521b5fa01fbce3525a46aa456cb63c884a522"
	contract.Name = "西藏母亲水窖项目二期施工"
	contract.Detail = "西藏母亲水窖项目二期施工"
	contract.StartTime = "2017-03-01"
	contract.EndTime = "2017-06-01"
	contract.PartyA = "中国妇女发展基金会"
	contract.PartyB = "中国第一建筑公司"
	contract.DepositBank = "光大银行"
	contract.BankAccount = "6226660000000000"

	return contract
}

func genBargain09() {
	fmt.Println("")
	fmt.Println("=========== Bargain09 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	bargainAddr := "bargain09:0fb2a0cdddddcfa33052c81b5e8ebbb05da19b69880487a49f44d2cfb37a7773"
	fmt.Println("=========== bargainAddr ==========")
	fmt.Println(bargainAddr)

	bargainPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEZ09IZytxeEVhV1g1SC94S0k0VEtwTjhwUApDcnJ2RUMwZytRZ0pHY2RVN2laREtUb3lRYS9CS3VIZlB2NC9OQXpiTEJLbGNRMWxyUzJldXZNdFZGOTg4S3VxCjdJc3BlTVZ1QWVHUEFDU2tNbWorN21HcjgxMVpYR0lhUGIrc0p1eGVJTktlckRGVVRiOTVjUlNtZzJYeVVWM0sKejJLZE5rOFUreGptaWFMbnB3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== bargainPublickKey ==========")
	fmt.Println(bargainPublickKey)

	bargain := initBargain09()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(bargainAddr + bargainPublickKey + base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
}

func initBargain09() protos.Bargain {
	var contract protos.Bargain

	contract.Addr = "bargain09:0fb2a0cdddddcfa33052c81b5e8ebbb05da19b69880487a49f44d2cfb37a7773"
	contract.Name = "甘肃母亲水窖项目一期施工"
	contract.Detail = "甘肃母亲水窖项目一期施工"
	contract.StartTime = "2017-01-01"
	contract.EndTime = "2017-03-01"
	contract.PartyA = "中国妇女发展基金会"
	contract.PartyB = "中国第三建筑公司"
	contract.DepositBank = "光大银行"
	contract.BankAccount = "6226660000000000"

	return contract
}

func genBargain10() {
	fmt.Println("")
	fmt.Println("=========== Bargain10 ==========")

	foundationAddr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== foundationAddr ==========")
	fmt.Println(foundationAddr)

	bargainAddr := "bargain10:9d681edb9d219cbb9885cc49d9bf76060febfa5e404e8c3a3fc269753250b62d"
	fmt.Println("=========== bargainAddr ==========")
	fmt.Println(bargainAddr)

	bargainPublickKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEUFdMOEo1TGtRZ0hHTzlnOTZ1dXdGVFNXegpuZ0puM0djelI4M3VtQmNENGhmN3hJanRWRHd0VkUrVmxXT1haV1hGbVd1TDk1azNJd3VXb2MvSzh5dmQyOXFEClM1c2IwQ3RjMUFEM3Q5akEwMDVVN0R0aXNpKzZ2NHBvdjA2M3MvK3RlVjR5NE45dStMSzVGaExoSko2a2JLbGkKOW5wZDBQK3Z6UUlRQXYwU3N3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== bargainPublickKey ==========")
	fmt.Println(bargainPublickKey)

	bargain := initBargain10()

	jsonBargain, _ := json.Marshal(bargain)

	base64BargainString := base64.StdEncoding.EncodeToString(jsonBargain)

	fmt.Println("=========== data : ")
	fmt.Println(base64BargainString)

	bargainHashed := sha256.Sum256([]byte(bargainAddr + bargainPublickKey + base64BargainString))
	_signBargain, _ := utils.RsaSign(crypto.SHA256, bargainHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signBargain))
}

func initBargain10() protos.Bargain {
	var contract protos.Bargain

	contract.Addr = "bargain10:9d681edb9d219cbb9885cc49d9bf76060febfa5e404e8c3a3fc269753250b62d"
	contract.Name = "甘肃母亲水窖项目二期施工"
	contract.Detail = "甘肃母亲水窖项目二期施工"
	contract.StartTime = "2017-03-01"
	contract.EndTime = "2017-06-01"
	contract.PartyA = "中国妇女发展基金会"
	contract.PartyB = "中国第一建筑公司"
	contract.DepositBank = "光大银行"
	contract.BankAccount = "6226660000000000"

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
