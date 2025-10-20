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

type QRCodeData struct {
	Wallet    string // metamask & trust wallet only
	ChainId   uint
	Amount    uint
	Decimal   uint
	TokenType string
}

func buildBaseUrlByWallet(wallet string) string {
	var baseUrl string

	switch wallet {
	case "metamask":
		baseUrl = "https://metamask.app.link/send"
		// break

	case "trust":
		baseUrl = "ethereum:"
		// break

	default:
		error := errors.New("buildBaseUrlByWallet.go: unsupported wallet type")
		log.Fatalln(error.Error())
	}

	return baseUrl
}

func buildMetamaskDeeplinkWithPrefix(qd QRCodeData) string {
	baseUrl := buildBaseUrlByWallet(qd.Wallet)
	prefix := "pay"
	deeplink := ""

	switch qd.TokenType {
	case "eth":
		deeplink = fmt.Sprintf("%s/%s-%s@%d?value=%d", baseUrl, prefix, qd.Wallet, qd.ChainId, qd.Amount)

	case "erc20":

	default:
		error := errors.New("buildMetamaskDeeplinkWithPrefix.go: unsupported token type")
		log.Fatalln(error.Error())

	}
	// eth-pay
	// https://metamask.app.link/send/pay-0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1?value=1e15
	// erc20 transfer-pay
	// https://metamask.app.link/send/pay-0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1/transfer?address=0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11&uint256=1e6
	return deeplink
}

func buildMetamaskDeeplink() string {
	// eth-send
	// https://metamask.app.link/send/0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1?value=1e15
	// erc20 transfer-send(contract address@chainid)
	// https://metamask.app.link/send/0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1/transfer?address=0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11&uint256=1e6

	return ""
}

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
		if v.Name() == address+".png" {
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
