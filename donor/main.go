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
)

func main() {
	donorAddr := "donor01:275e74b0e340f54135496e46d829b25af699984e6787f9a7b13191ad991a1eb1"

	fmt.Println("=========== donor ===========")
	donor := genDonor(donorAddr)

	jsonDonor, _ := json.Marshal(donor)

	base64DonorString := base64.StdEncoding.EncodeToString(jsonDonor)

	fmt.Println("=========== data : ")
	fmt.Println(base64DonorString)

	donorHashed := sha256.Sum256([]byte(base64DonorString))
	_signDonor, _ := RsaSign(crypto.SHA256, donorHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signDonor))

	fmt.Println("")
	fmt.Println("=========== contribution ===========")
	contribution := genContribution()

	jsonContribution, _ := json.Marshal(contribution)

	base64ContributionString := base64.StdEncoding.EncodeToString(jsonContribution)

	fmt.Println("=========== data : ")
	fmt.Println(base64ContributionString)

	contributionHashed := sha256.Sum256([]byte(donorAddr + base64ContributionString))
	_signContribution, _ := RsaSign(crypto.SHA256, contributionHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signContribution))

	fmt.Println("")
	fmt.Println("=========== track ===========")
	track := genTrack()

	jsonTrack, _ := json.Marshal(track)

	base64TrackString := base64.StdEncoding.EncodeToString(jsonTrack)

	fmt.Println("=========== data : ")
	fmt.Println(base64TrackString)

	trackHashed := sha256.Sum256([]byte(donorAddr + base64TrackString))
	_signTrack, _ := RsaSign(crypto.SHA256, trackHashed[:], privateKey)

	fmt.Println("=========== sign : ")
	fmt.Println(base64.StdEncoding.EncodeToString(_signTrack))
}

func genDonor(donorAddr string) Donor {
	var donor Donor

	donor.Addr = donorAddr
	donor.Name = "donorUser01"
	donor.Total = 1000 * 100 * 1000

	return donor
}

func genContribution() DonorContribution {
	var contribution DonorContribution

	contribution.Name = "donorUser01"
	contribution.SmartContractName = "宁夏西部地区母亲水窖项目"
	contribution.SmartContractAddr = "smartcontract01:1d54a8713923af1718e8eeabec3e4d8596dbbdf2da3f69ea23aeb8c7a5ab73d8"
	contribution.Amount = 1000 * 100 * 1000
	contribution.Timestamp = time.Now().UTC().Unix()

	return contribution
}

func genTrack() DonorTrack {
	var track DonorTrack

	track.Name = "donorUser01"
	track.AccountName = "宁夏西部地区母亲水窖项目XX县XX村水窖"
	track.AccountAddr = "bargain01:8fcc58ea7ed212f7c1ba359d15bea144e67c390044d953797548cf67fd62534a"
	track.Amount = 1000 * 100 * 1000
	track.Timestamp = time.Now().UTC().Unix()

	return track
}

var privateKey, publicKey []byte

func init() {
	var err error

	filepath := "/home/eric/go/src/github.com/CebEcloudTime/rsaserver/"

	privateKey, err = ioutil.ReadFile(filepath + "channel01private.pem")
	if err != nil {
		os.Exit(-1)
	}

	publicKey, err = ioutil.ReadFile(filepath + "channel01public.pem")
	if err != nil {
		os.Exit(-1)
	}

}
