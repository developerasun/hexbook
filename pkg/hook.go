package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	constant "github.com/hexbook/internal/constant"
	qrcode "github.com/skip2/go-qrcode"
)

func validateAddress(address string) {
	_, found := strings.CutPrefix(address, "0x")

	if !found || len(address) != 42 {
		error := errors.New("validateAddress.go: invalid ethereum address")
		log.Fatalln(error.Error(), "| address length: ", len(address))
	}
}

func validateDuplicate(address string) bool {
	wd, gErr := os.Getwd()

	if gErr != nil {
		log.Fatalln("validateDuplicate: ", gErr.Error())
	}

	targetPath := strings.Join([]string{wd, "assets", "qrcode"}, "/")
	entries, rErr := os.ReadDir(targetPath)

	if rErr != nil {
		log.Fatalln("validateDuplicate: ", rErr.Error())
	}

	var isExisting bool = false

	for _, v := range entries {
		if v.Name() == address {
			isExisting = true
			break
		}
	}

	return isExisting
}

/*
@docs https://dev-docs.dcentwallet.com/dynamic-link/eip-681-transaction-payment-request#eip681-dynamic-link-format
TODO check metamask qrcode not working
*/
func makeResourceEip681Compatible(address string) string {
	// @dev ethereum mainnet, 0.001 ether
	protocol := "ethereum"
	chainId := 1
	link := fmt.Sprintf("%s:%s@%d?value=1000000000000000", protocol, address, chainId)

	return link
}

func GenerateQrCode(wallet string) string {
	validateAddress(wallet)
	isExisting := validateDuplicate(wallet)
	link := makeResourceEip681Compatible(wallet)
	filename := fmt.Sprintf("%s.png", wallet)

	if !isExisting {
		log.Println("GenerateQrCode: detecting new entry for qrcode, starting encoding...", filename)
		png, err := qrcode.Encode(link, qrcode.Medium, 256)

		if err != nil {
			log.Fatalln(err.Error())
		}

		wd, gErr := os.Getwd()

		if gErr != nil {
			log.Fatalln(gErr.Error())
		}

		targetPath := strings.Join([]string{wd, "assets", "qrcode", filename}, "/")
		wErr := os.WriteFile(targetPath, png, constant.FilePermUserReadWriteGroupRead)

		if wErr != nil {
			log.Fatalln(wErr.Error())
		}

		return filename
	}

	log.Println("GenerateQrCode: detecting existing entry for qrcode, terminating...", filename)
	return filename
}
