package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/CebEcloudTime/charitycc/errors"
	// "github.com/CebEcloudTime/charitycc/protos"
	"github.com/CebEcloudTime/charitycc/utils"
)

func main() {

	base64TxData := "eyJ0eG91dCI6W3sic2lnbiI6IiIsImFkZHIiOiJjZWJiYW5rOjI5NzMxZDBlNmM2Y2E5Y2I5ODVlYWJmOWZlNzE2ZDE2NDRjNjI0Y2FlNTI2NWMzNmM5YjdhNDY3MDIwMDM5MjQiLCJhdHRyIjoiIiwidmFsdWUiOjk5OTg1NTAwMDAwMDB9LHsic2lnbiI6IiIsImFkZHIiOiJkb25vcjAyOmI5MzE1Nzg4NTdkYmU5Y2IzN2QwZTU5Y2Q5Y2M1ZmIzNzU4NzU4ZDBjZWUxZTkwNWUyNWZlNjJmOWU5ZTI2ODgiLCJhdHRyIjoiZG9ub3IwMjpiOTMxNTc4ODU3ZGJlOWNiMzdkMGU1OWNkOWNjNWZiMzc1ODc1OGQwY2VlMWU5MDVlMjVmZTYyZjllOWUyNjg4LDg2ZmE4YmEwLTJjYjItNDY4Yi04YjRjLWQ3N2NhMGUyOTQzYiIsInZhbHVlIjozMDAwMDAwMDB9XSwiaW5wdXREYXRhIjoiIiwiZm91bmRlciI6ImNlYmJhbms6Mjk3MzFkMGU2YzZjYTljYjk4NWVhYmY5ZmU3MTZkMTY0NGM2MjRjYWU1MjY1YzM2YzliN2E0NjcwMjAwMzkyNCIsInZlcnNpb24iOjIwMTcwMTAxLCJ0aW1lc3RhbXAiOjE0ODgyMDI3MjYsInR4aW4iOlt7InNvdXJjZVR4SGFzaCI6IjdkZjliNTQzMjJlZDhiNjhhNDRmZDk4YTliZmYyNmI4OWNhZjgzNGJiZjU3YzQ4YzJmMjRkMTZkMDFhZDczOGEiLCJhZGRyIjoiY2ViYmFuazoyOTczMWQwZTZjNmNhOWNiOTg1ZWFiZjlmZTcxNmQxNjQ0YzYyNGNhZTUyNjVjMzZjOWI3YTQ2NzAyMDAzOTI0IiwiaWR4IjowfV19"

	// base64TxData := ""
	txData, err := base64.StdEncoding.DecodeString(base64TxData)
	if err != nil {
		fmt.Println(errors.Base64Decoding)
		os.Exit(1)
	}

	newTX, err := utils.ParseTxByJsonBytes(txData)
	if err != nil {
		fmt.Println(errors.InvalidTX)
		os.Exit(1)
	}

	if newTX.Founder == "" {
		fmt.Println(errors.TxNoFounder)
		os.Exit(1)
	}

	fmt.Println(newTX)
}
