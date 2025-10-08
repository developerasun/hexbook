package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"

	// "path"
	"strings"

	"github.com/google/uuid"
	constant "github.com/hexbook/internal/constant"
	qrcode "github.com/skip2/go-qrcode"
)

func validateAddress(address string) {
	_, found := strings.CutPrefix(address, "0x")

	if !found || len(address) != 42 {
		error := errors.New("validateAddress.go: invalid ethereum address\n")
		log.Fatalln(error.Error(), "address length: ", len(address))
	}
}

func GenerateQrCode(wallet string) string {
	validateAddress(wallet)
	png, err := qrcode.Encode(wallet, qrcode.Medium, 256)

	if err != nil {
		log.Fatalln(err.Error())
	}

	wd, gErr := os.Getwd()

	if gErr != nil {
		log.Fatalln(gErr.Error())
	}

	id := uuid.New().String()
	filename := fmt.Sprintf("%s.png", id)

	log.Println("wd: ", wd, "filename: ", filename)

	targetPath := strings.Join([]string{wd, "assets", "qrcode", filename}, "/")
	wErr := os.WriteFile(targetPath, png, constant.FilePermUserReadWriteGroupRead)

	if wErr != nil {
		log.Fatalln(wErr.Error())
	}

	return filename
}
