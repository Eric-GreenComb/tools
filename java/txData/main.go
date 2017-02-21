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

	base64TxData := "eyJ0eG91dCI6W3sic2lnbiI6IiIsImFkZHIiOiJzbWFydGNvbnRyYWN0MDE6MWQ1NGE4NzEzOTIzYWYxNzE4ZThlZWFiZWMzZTRkODU5NmRiYmRmMmRhM2Y2OWVhMjNhZWI4YzdhNWFiNzNkOCIsImF0dHIiOiJkb25vcjAxOjI3NWU3NGIwZTM0MGY1NDEzNTQ5NmU0NmQ4MjliMjVhZjY5OTk4NGU2Nzg3ZjlhN2IxMzE5MWFkOTkxYTFlYjEsYzMzMjg2NGUtZmU5OS00MDE5LTg0MzgtYzQwYmUwOWIzNDVmIiwidmFsdWUiOjUwMDAwMDAwMDAwfV0sImlucHV0RGF0YSI6IiIsImZvdW5kZXIiOiJzbWFydGNvbnRyYWN0MDE6MWQ1NGE4NzEzOTIzYWYxNzE4ZThlZWFiZWMzZTRkODU5NmRiYmRmMmRhM2Y2OWVhMjNhZWI4YzdhNWFiNzNkOCIsInZlcnNpb24iOjIwMTcwMTAxLCJ0aW1lc3RhbXAiOjE0ODc1OTE0NzQsInR4aW4iOlt7InNvdXJjZVR4SGFzaCI6ImJlM2U3Y2NlZTA5OGFiMmQzZTg3NGNmMmJkNTg0YTZiMjA2NDExYTdlYjRmZDFhM2U3OWRmZmM1OTgyNWQxYWUiLCJhZGRyIjoiZG9ub3IwMToyNzVlNzRiMGUzNDBmNTQxMzU0OTZlNDZkODI5YjI1YWY2OTk5ODRlNjc4N2Y5YTdiMTMxOTFhZDk5MWExZWIxIiwiaWR4IjoxfV19"
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
