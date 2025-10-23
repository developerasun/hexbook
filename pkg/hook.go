package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	constant "github.com/hexbook/internal/constant"
	"github.com/shopspring/decimal"
	qrcode "github.com/skip2/go-qrcode"
)

type QRCodeData struct {
	AppType   string // metamask & trust wallet only
	Wallet    string // 0x123...789
	ChainId   uint
	Amount    string // 1e15, this is string since it's for uri
	TokenType string
}

type UriOption struct {
	Prefix string
}

func BuildBaseUrlByAppType(appType string) string {
	var baseUrl string

	switch appType {
	case "metamask":
		baseUrl = "https://metamask.app.link/send"

	case "trust":
		baseUrl = "ethereum" // eip681 protocol

	default:
		error := errors.New("BuildBaseUrlByAppType.go: unsupported wallet type")
		log.Panicln(error.Error())
	}

	return baseUrl
}

/*
@docs https://dev-docs.dcentwallet.com/dynamic-link/eip-681-transaction-payment-request#eip681-dynamic-link-format

@demo https://metamask.github.io/metamask-deeplinks/

@example1 eth
https://metamask.app.link/send/pay-0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1?value=1e15

@example2 erc20
https://metamask.app.link/send/pay-0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1/transfer?address=0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11&uint256=1e6

@example3
https://metamask.app.link/send/0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1?value=1e15

@example4
https://metamask.app.link/send/0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1/transfer?address=0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11&uint256=1e6
*/
func BuildQRCodeDeeplink(qd QRCodeData, option *UriOption) string {
	baseUrl := BuildBaseUrlByAppType(qd.AppType)
	deeplink := ""

	if qd.AppType == "metamask" {
		switch qd.TokenType {
		case "eth":
			if option != nil {
				// @dev build eip681 uri with `pay` prefix
				deeplink = fmt.Sprintf("%s/%s-%s@%d?value=%s", baseUrl, option.Prefix, qd.Wallet, qd.ChainId, qd.Amount)
			} else {
				deeplink = fmt.Sprintf("%s/%s@%d?value=%s", baseUrl, qd.Wallet, qd.ChainId, qd.Amount)
			}
		case "erc20":
			if option != nil {
				// @dev build eip681 uri with `pay` prefix
				deeplink = fmt.Sprintf("%s/%s-%s@%d/transfer?address=%s&uint256=%s", baseUrl, option.Prefix, qd.Wallet, qd.ChainId, qd.Wallet, qd.Amount)
			} else {
				deeplink = fmt.Sprintf("%s/%s@%d/transfer?address=%s&uint256=%s", baseUrl, qd.Wallet, qd.ChainId, qd.Wallet, qd.Amount)
			}

		default:
			error := errors.New("BuildQRCodeDeeplink.go: unsupported token type")
			log.Fatalln(error.Error())
		}
	}

	if qd.AppType == "trust" {
		deeplink = fmt.Sprintf("%s:%s@%d?value=%s", baseUrl, qd.Wallet, qd.ChainId, qd.Amount)
	}

	log.Println("apptype: ", qd.AppType, " deeplink: ", deeplink)

	return deeplink
}

func validateAppType(appType string) {
	if appType != "metamask" && appType != "trust" {
		log.Fatalln("validateAppType: invalid wallet app type")
	}
}

func validateAddress(address string) {
	_, found := strings.CutPrefix(address, "0x")

	if !found || len(address) != 42 {
		error := errors.New("validateAddress.go: invalid ethereum address")
		log.Fatalln(error.Error(), "| address length: ", len(address))
	}
}

func validateDuplicate(address string, appType string) bool {
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

	filename := fmt.Sprintf("%s-%s.png", address, appType)
	for _, v := range entries {
		if v.Name() == filename {
			isExisting = true
			break
		}
	}

	return isExisting
}

/*
@return e.g `1000000000000000000`
*/
func toWei(_amount string) string {
	amount, err := decimal.NewFromString(_amount)

	if err != nil {
		log.Fatalln(err.Error())
	}
	plain := decimal.NewFromFloat(constant.ETH_DECIMAL)
	calculated := amount.Mul(plain)

	return calculated.String()
}

/*
@return e.g `N*1e18`
*/
func toWeiAsExponent(_amount string) string {
	toFloat, _ := strconv.ParseFloat(_amount, 64)
	target := fmt.Sprintf("%.e", toFloat*constant.ETH_DECIMAL)

	converted := strings.Replace(target, "+", "", 1)
	return converted
}

func GenerateQrCode(appType string, wallet string, amount string) string {
	validateAppType(appType)
	validateAddress(wallet)
	isExisting := validateDuplicate(wallet, appType)

	var deeplink string
	if appType == "metamask" {
		deeplink = BuildQRCodeDeeplink(QRCodeData{
			AppType: "metamask",
			Wallet:  wallet,
			ChainId: 1,
			// TODO replace hardcoded decimals
			Amount:    toWeiAsExponent(amount), // hardcoded
			TokenType: "eth",
		}, nil)
	} else {
		deeplink = BuildQRCodeDeeplink(QRCodeData{
			AppType:   "trust",
			Wallet:    wallet,
			ChainId:   1,
			Amount:    toWei(amount), // hardcoded
			TokenType: "eth",
		}, nil)
	}

	filename := fmt.Sprintf("%s-%s.png", wallet, appType)

	if !isExisting {
		log.Println("GenerateQrCode: detecting new entry for qrcode, starting encoding...", filename)
		png, err := qrcode.Encode(deeplink, qrcode.Medium, 256)

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
