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

func GenerateQrCode(wallet string) string {
	validateAddress(wallet)
	isExisting := validateDuplicate(wallet)
	filename := fmt.Sprintf("%s.png", wallet)

	if !isExisting {
		log.Println("GenerateQrCode: detecting new entry for qrcode, starting encoding...", filename)
		png, err := qrcode.Encode(wallet, qrcode.Medium, 256)

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
