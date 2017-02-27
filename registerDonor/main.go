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
	GenDonor01Info()
	GenDonor02Info()
	GenDonor03Info()
	GenDonor04Info()
	GenDonor05Info()
	GenDonor06Info()
	GenDonor07Info()
	GenDonor08Info()
	GenDonor09Info()
	GenDonor10Info()
}

func GenDonor01Info() {
	fmt.Println("")
	fmt.Println("=========== donor01 ===========")

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

func GenDonor02Info() {
	fmt.Println("")
	fmt.Println("=========== donor02 ===========")

	channelAddr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fmt.Println("=========== channelAddr ===========")
	fmt.Println(channelAddr)

	donorAddr := "donor02:b931578857dbe9cb37d0e59cd9cc5fb3758758d0cee1e905e25fe62f9e9e2688"
	fmt.Println("=========== donorAddr ===========")
	fmt.Println(donorAddr)

	donorPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERjVCdlJKU2dWR01rZVBBUkRGWlZsZGdTVQpOWnJmY1ZzREJjYjRSVmpxVmdWRGw0azJsdXJRQXpLejZocEgvdnJGV3dKdDM4R0ZLVDBiYTRQL1Vxakg0TE5aCmc4SGp2cnpsZUk1TDJ0cHpwSXVXMlY5YTZCbWFIOEpnajlPa0p0Vk9NMEJOL3VoeEZRQmlOT3M5TkE2VlltY04KUGFXaElwQWJUS1dKaW5FVHlRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== donorPublicKey ===========")
	fmt.Println(donorPublicKey)

	donorHashed := sha256.Sum256([]byte(donorAddr + donorPublicKey))
	_signDonor, _ := utils.RsaSign(crypto.SHA256, donorHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signDonor))
}

func GenDonor03Info() {
	fmt.Println("")
	fmt.Println("=========== donor03 ===========")

	channelAddr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fmt.Println("=========== channelAddr ===========")
	fmt.Println(channelAddr)

	donorAddr := "donor03:a65557075090886383f6c67220e22b095dcfe1f3a6fcb47f6fb622a3ce0f1c2d"
	fmt.Println("=========== donorAddr ===========")
	fmt.Println(donorAddr)

	donorPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEQ0Z4T3JXNFJZMnBlcVhreEJieDhVa2oxVgpZVkVTMGd0QmhzVk5rZFNyaTRzL2d1dnlNcjF2b1g1VGJmNnpHclVRbjFPQ3F2aHdFTFJjb1pXMnp4OEpmWHhTCjRGTXhQWmJSL3NPaXI0K2Nhc1M3YlUyVG9zOU5DSjJMdlVNbC9wMldHUklQVGgzeitCLzdORXBqNndEYkdyVDgKYTY1V1U1ZFlKMXJrUWRZUmpRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== donorPublicKey ===========")
	fmt.Println(donorPublicKey)

	donorHashed := sha256.Sum256([]byte(donorAddr + donorPublicKey))
	_signDonor, _ := utils.RsaSign(crypto.SHA256, donorHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signDonor))
}

func GenDonor04Info() {
	fmt.Println("")
	fmt.Println("=========== donor04 ===========")

	channelAddr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fmt.Println("=========== channelAddr ===========")
	fmt.Println(channelAddr)

	donorAddr := "donor04:55e79613ab432a03a9c894e3efbbe5d3f0a445feea82ff49c98cb03302ea6541"
	fmt.Println("=========== donorAddr ===========")
	fmt.Println(donorAddr)

	donorPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEQ25JaWxhNHpZbTVvYytmTEpHS0s4dTFZNQo1QzlTck1YMSsranViWUJDYkNwRS9DblVoUFl3WnNlejNXOHRKSThYWDNwaFBSN2MzV2Mvd1E3SFpaK1J3SGcyCjhtQjV3L3BaNHlJUFozb2loSzMrU1BEZVVjMTR1cjVCdXliRzJrcy9JOHVOT3JEM0RLS0hCNkY2Z1g2RGNkeHMKeDd6SjJhV1Zmd1dkSk1CbHh3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== donorPublicKey ===========")
	fmt.Println(donorPublicKey)

	donorHashed := sha256.Sum256([]byte(donorAddr + donorPublicKey))
	_signDonor, _ := utils.RsaSign(crypto.SHA256, donorHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signDonor))
}

func GenDonor05Info() {
	fmt.Println("")
	fmt.Println("=========== donor05 ===========")

	channelAddr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fmt.Println("=========== channelAddr ===========")
	fmt.Println(channelAddr)

	donorAddr := "donor05:f799bc349d140248bb1ea2bd1f9a820ae99161b9bc59c0ac212f6d276e9a2468"
	fmt.Println("=========== donorAddr ===========")
	fmt.Println(donorAddr)

	donorPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDanpGb2xsejlqYlNsamhqcEFJS2VZczF4dwpDWG9pQWcwT0E3OVdVZnA1a0tvZmRlbTNpY0JaOS9XZEhwVmZPZHNmTTZOWUFQVHE0RURObUloWnlBbmVsSTBZCkFrQ3hhUWVFLy9zd2czWHVZVHpYRlFIWktvNEkxdFJHSTJWaG0vdXNDbm44bUxpYnhVeE5wUlF0SEVJT1AwaXYKNDhwRUFPcjZ2Tnd1MUpxYVh3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== donorPublicKey ===========")
	fmt.Println(donorPublicKey)

	donorHashed := sha256.Sum256([]byte(donorAddr + donorPublicKey))
	_signDonor, _ := utils.RsaSign(crypto.SHA256, donorHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signDonor))
}

func GenDonor06Info() {
	fmt.Println("")
	fmt.Println("=========== donor06 ===========")

	channelAddr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fmt.Println("=========== channelAddr ===========")
	fmt.Println(channelAddr)

	donorAddr := "donor06:293db3e941705aaabe9505e4a513ceac8b51b8ae1407c28b7eb8ff88193eb2f6"
	fmt.Println("=========== donorAddr ===========")
	fmt.Println(donorAddr)

	donorPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDcndFakJEUUwzdGtVNy9HUkZPYmZvbHhidQpMOW9RK0dNZlR4a2M1eXFXeHlJZ3djSkNidSthV3FYUFc2WFVXY0pJOVlRUkVHN1pXZWRlUVRwWVpFSGloR1pWClVFQzdVQlFMWGV0ZjdNR1UwUUJOc0tOZ054cUpRUHp6ZWxreC90ZGpQWHo4QklIUkMyNGRmRXdqTVhwWU5mRFQKeUJLWC9yYnhGRTgwRVlmR1B3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== donorPublicKey ===========")
	fmt.Println(donorPublicKey)

	donorHashed := sha256.Sum256([]byte(donorAddr + donorPublicKey))
	_signDonor, _ := utils.RsaSign(crypto.SHA256, donorHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signDonor))
}

func GenDonor07Info() {
	fmt.Println("")
	fmt.Println("=========== donor07 ===========")

	channelAddr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fmt.Println("=========== channelAddr ===========")
	fmt.Println(channelAddr)

	donorAddr := "donor07:0abc5a0def68504b23bd4fafde9fa089315f60ae6077535643ee6705643115f1"
	fmt.Println("=========== donorAddr ===========")
	fmt.Println(donorAddr)

	donorPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEV0pteXlSVDNxTzlSYlFFVHY3Y0tZU29HSAo0VlNOWTBqUVFCN1l2Ukg3WlR4UkovNFlCSTZiNnlER2NmWDNtTVZhdzJyNkg2Zk9wUVFWT1UrWDVjWjcvYmdFClpEdmNQa01oL3RzSFRGbklIS01KeEN0T0g5Y01jKytEaWVUWkFYKzRjUFFOVHdUWWpibXZmV29IYkJwazhTRnEKYjIwZjR0SEVoY1g4TXpyYmVRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== donorPublicKey ===========")
	fmt.Println(donorPublicKey)

	donorHashed := sha256.Sum256([]byte(donorAddr + donorPublicKey))
	_signDonor, _ := utils.RsaSign(crypto.SHA256, donorHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signDonor))
}

func GenDonor08Info() {
	fmt.Println("")
	fmt.Println("=========== donor08 ===========")

	channelAddr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fmt.Println("=========== channelAddr ===========")
	fmt.Println(channelAddr)

	donorAddr := "donor08:dc0e0c13816ba9270f23fb67338c21e68459dc977c85dc8ef7f4eff0e5df6189"
	fmt.Println("=========== donorAddr ===========")
	fmt.Println(donorAddr)

	donorPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDZ09hSHBwRmhIREJWTVNMRTFOWXhqS2RVbQpNSzQ1MXR3V3daV3B2R2RTdEJGQzhMaUc0b0VobjdjbEN4NTFqa3NhazVLdGJBSGxtRVNucW1JKzVCR2xRblhiCm11OEEvajVKSHVtK1RxSEpPQlJaQW96YVliWFNEUFhBZzJpNnQvNDcvbmdKWDhrVllObjVTM1pXQXlhak1HNisKejBHcU5rVGE1ZnB3N1JBOWF3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== donorPublicKey ===========")
	fmt.Println(donorPublicKey)

	donorHashed := sha256.Sum256([]byte(donorAddr + donorPublicKey))
	_signDonor, _ := utils.RsaSign(crypto.SHA256, donorHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signDonor))
}

func GenDonor09Info() {
	fmt.Println("")
	fmt.Println("=========== donor09 ===========")

	channelAddr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fmt.Println("=========== channelAddr ===========")
	fmt.Println(channelAddr)

	donorAddr := "donor09:e98462be3aebacb5bf8acf6b0d3ef46331016b4263b787a0867dea25613fb6ae"
	fmt.Println("=========== donorAddr ===========")
	fmt.Println(donorAddr)

	donorPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDZ2Foc0dhcjJWcUR0bUQ0SlFuZlhKaG9LbwpHR0luY1hmZmpYUTk2czFKMGJNeTBVYWFDWHNPMUM2UHp3MmhsUHhQTTZmRGpOb1o2WHRNcDlnd0FNYVVMN3RUCmpqekI4UnpXeGkvbThzTnYwUys4ZUNmVUw3N0haZS9CRzFlL3VuTkNsZTFvKy9oL2U1SUVuNWtQaTExQVR6N2QKL1F2MGFybHdKeTFhbk8zQ0l3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	fmt.Println("=========== donorPublicKey ===========")
	fmt.Println(donorPublicKey)

	donorHashed := sha256.Sum256([]byte(donorAddr + donorPublicKey))
	_signDonor, _ := utils.RsaSign(crypto.SHA256, donorHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signDonor))
}

func GenDonor10Info() {
	fmt.Println("")
	fmt.Println("=========== donor10 ===========")

	channelAddr := "channel01:9c8b43ce948010efc3b7d102aae502165ccd5e0714a3e765fe1a8f444936785a"
	fmt.Println("=========== channelAddr ===========")
	fmt.Println(channelAddr)

	donorAddr := "donor10:79289520e19c9d80ca228d22ca2c11dc775348cdfece34cfc32f99dd5ef89243"
	fmt.Println("=========== donorAddr ===========")
	fmt.Println(donorAddr)

	donorPublicKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDbTZoTFh4ak5tS0Jmemp2T3pkNU1EQUx5cApnU2piMmcreU00Qmx3RjJsU083ZnZFckcxRGF3UmZpMVJ0MHBnTW5saU93VDc1RGFySUIwT0pzdDY5Y2U2WWFXCjZIWFoyTjdqVDVCTmc2THZPbE5vSW13S2dUL1Q2WVQzVjZVNlRQV0FjY2hNWlZxSE1hbloyK1lzSUFiUVB1OEYKQ00zNUZZR3drMVJjdkZnU0tRSURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
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
