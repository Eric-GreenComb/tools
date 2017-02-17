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

	cebbankAddr := "cebbank:29731d0e6c6ca9cb985eabf9fe716d1644c624cae5265c36c9b7a46702003924"
	fmt.Println("=========== cebbankAddr ==========")
	fmt.Println(cebbankAddr)

	cebbankPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDenFYaVVyRFZzanhIZzdyM1RMc1NSaGlaTApqeE1nczBIU3JsRlV3c0s1eFVCcGFwSHdSazlHeGJrMWtOd2tkSzdPeHlselJjbmxzbEd5VnhhU21KYzNqbmpvClVRVFphbVRaemViMzNNZC9oYUVkN3BhSXFkSS9pZ1Z6TEtpaStXcFJkUE02VlFaV0pHam01eEMwMHhWUWw3aWcKOExBUXV6ek1OMnNCRkhsSmN3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== cebbankPublicKey ==========")
	fmt.Println(cebbankPublicKey)

	hashed := sha256.Sum256([]byte(cebbankAddr + cebbankPublicKey))
	_sign, _ := utils.RsaSign(crypto.SHA256, hashed[:], systemadminprivate)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))

	fmt.Println("")
	fund01Addr := "fund01:25ab580a2093776ca2e1dd1775e96dfec5f1ffbcc9565129351cb330cf0712d7"
	fmt.Println("=========== fund01Addr ==========")
	fmt.Println(fund01Addr)

	fund01PublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEanZlcE94YkI1SGlQbVgwcnJTUThJMG90MQp2c0ljNno0RThtVnhubEtlMCt3MVpOT1J1UEdwSmE1TENNZWNqNDk4dllDcEljbk1IMkJCZjJQMk4wcnVkOW5PCllOVWZiOHlxdk5aODZNNHZBRVM4d2QxT1MzYmY1bm12Ry92Ykl5aWQwSmdIK0tRNU1LVDFZSjNGRE5TcWYrbHQKRHVTOWFUeDhtb3BKODlSME93SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== fund01PublicKey ==========")
	fmt.Println(fund01PublicKey)

	hashedFund01 := sha256.Sum256([]byte(fund01Addr + fund01PublicKey))
	_signFund01, _ := utils.RsaSign(crypto.SHA256, hashedFund01[:], systemadminprivate)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signFund01))

	fmt.Println("")
	channel01Addr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fmt.Println("=========== channel01Addr ==========")
	fmt.Println(channel01Addr)

	channel01PublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERGVNNzJVRnkvK1VTaUpTU3BkWlF2Tmt2NgpySDAvQ25raUJMRmRBUEVIUUR6bTY0Wks3bWR0bUE2ayt1alJtbENXbjhIUVRSK0xHcGp4WGxhVFc1a1Z2a2xVCks4cUdkRUUzNW5JTW0zQnc2ZldsQUlvV2I4dUxEQjlVQmpIVGxIb1dlaEE0SzV3WEh6ZmpqaWQ0U2dBNEhwbzUKdlhlazdzc3RESEpBZUhuNXdRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== channel01PublicKey ==========")
	fmt.Println(channel01PublicKey)

	hashedchannel01 := sha256.Sum256([]byte(channel01Addr + channel01PublicKey))
	_signchannel01, _ := utils.RsaSign(crypto.SHA256, hashedchannel01[:], systemadminprivate)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signchannel01))
}

var systemadminprivate []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	systemadminprivate, err = ioutil.ReadFile(filepath + "systemadminprivate.pem")
	if err != nil {
		os.Exit(-1)
	}

}
