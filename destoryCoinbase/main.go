package main

import (
	"crypto"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	targetAddr := "bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"

	fmt.Println("=========== data : ")
	fmt.Println(targetAddr)

	targetAddrHashed := sha256.Sum256([]byte(targetAddr))
	_sign, _ := RsaSign(crypto.SHA256, targetAddrHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_sign))
}

var privateKey []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	privateKey, err = ioutil.ReadFile(filepath + "cebbankprivate.pem")
	if err != nil {
		os.Exit(-1)
	}

}
