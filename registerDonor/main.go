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

	fmt.Println("=========== donor ===========")

	channelAddr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fmt.Println("=========== channelAddr ===========")
	fmt.Println(channelAddr)

	donorAddr := "donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1"
	fmt.Println("=========== donorAddr ===========")
	fmt.Println(donorAddr)

	donorPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERS9HZ3VWbUk0QmRnUU5oYThPMmtYK3dHQgpNSkxQbGNPZllLMUdsQkpsZzlLUm9RNmZrWGd5aXRCb2FORzI5dVpnV3loVjZSMmhzUXpmRyt3cHE0VmsyQkEyClBxZEtFcVRyQVhMU0EvWkw0bUpPVEx4bjV0Qm5QL0p5OGRGaitXSjBlSWtzOGZhMUxORS9QVnR6UDlEZDV2bmEKbEZvTXNJa09maDlMcWxmaFRRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== donorPublicKey ===========")
	fmt.Println(donorPublicKey)

	donorHashed := sha256.Sum256([]byte(donorAddr + donorPublicKey))
	_signDonor, _ := utils.RsaSign(crypto.SHA256, donorHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signDonor))

}

var privateKey []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	privateKey, err = ioutil.ReadFile(filepath + "channel01private.pem")
	if err != nil {
		os.Exit(-1)
	}

}
