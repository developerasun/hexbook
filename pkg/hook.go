package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
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

@example2 erc20(usdt)
https://metamask.app.link/send/pay-0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1/transfer?address=0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11&uint256=1e6

@example3
https://metamask.app.link/send/0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1?value=1e15

@example4
https://metamask.app.link/send/0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1/transfer?address=0x7dBF026bd945295b2b492458FcA47Ed503F6e45F&uint256=1e6
*/
func BuildQRCodeDeeplink(qd QRCodeData, option *UriOption) string {
	baseUrl := BuildBaseUrlByAppType(qd.AppType)
	deeplink := ""

	if qd.AppType == "metamask" {
		switch qd.TokenType {
		case "ether":
			if option != nil {
				// @dev build eip681 uri with `pay` prefix
				deeplink = fmt.Sprintf("%s/%s-%s@%d?value=%s", baseUrl, option.Prefix, qd.Wallet, qd.ChainId, qd.Amount)
			} else {
				deeplink = fmt.Sprintf("%s/%s@%d?value=%s", baseUrl, qd.Wallet, qd.ChainId, qd.Amount)
			}
		case "usdt":
			if option != nil {
				// @dev build eip681 uri with `pay` prefix
				deeplink = fmt.Sprintf("%s/%s-%s@%d/transfer?address=%s&uint256=%s", baseUrl, option.Prefix, constant.ETH_USDT_ADDRESS, qd.ChainId, qd.Wallet, qd.Amount)
			} else {
				deeplink = fmt.Sprintf("%s/%s@%d/transfer?address=%s&uint256=%s", baseUrl, constant.ETH_USDT_ADDRESS, qd.ChainId, qd.Wallet, qd.Amount)
			}

		default:
			error := errors.New("BuildQRCodeDeeplink.go: unsupported token type: ")
			log.Fatalln(error.Error() + qd.TokenType)
		}
	}

	if qd.AppType == "trust" {
		switch qd.TokenType {
		case "ether":
			deeplink = fmt.Sprintf("%s:%s@%d?value=%s", baseUrl, qd.Wallet, qd.ChainId, toWei(qd.Amount, qd.TokenType))
		case "usdt":
			baseUrl = "https://link.trustwallet.com/send?coin=60"
			deeplink = fmt.Sprintf("%s&address=%s&amount=%s&token_id=%s", baseUrl, qd.Wallet, qd.Amount, constant.ETH_USDT_ADDRESS)
		}
	}

	log.Println("apptype: ", qd.AppType)
	log.Println("deeplink: ", deeplink)
	log.Println("tokentype: ", qd.TokenType)

	return deeplink
}

func validateAppType(appType string) error {
	if appType != "metamask" && appType != "trust" {
		error := errors.New("validateAppType: invalid wallet app type")
		return error
	}

	return nil
}

func validateAddress(address string) error {
	_, found := strings.CutPrefix(address, "0x")

	if !found || len(address) != 42 {
		error := errors.New("validateAddress.go: invalid ethereum address")
		return error
	}

	return nil
}

/*
@return e.g `1000000000000000000`
*/
func toWei(_amount string, _tokenType string) string {
	amount, err := decimal.NewFromString(_amount)

	if err != nil {
		log.Fatalln(err.Error())
	}

	targetDecimal := constant.ETH_DECIMAL
	if _tokenType == "usdt" {
		targetDecimal = constant.USDT_DECIMAL
	}

	plain := decimal.NewFromFloat(targetDecimal)
	calculated := amount.Mul(plain)

	return calculated.String()
}

/*
@return e.g `N*1e18`
*/
func toWeiAsExponent(_amount string, _tokenType string) string {
	targetDecimal := constant.ETH_DECIMAL
	if _tokenType == "usdt" {
		targetDecimal = constant.USDT_DECIMAL
	}

	toFloat, _ := strconv.ParseFloat(_amount, 64)
	target := fmt.Sprintf("%.e", toFloat*targetDecimal)

	converted := strings.Replace(target, "+0", "", 1)
	return converted
}

func GenerateQrCode(appType string, wallet string, amount string, tokenType string) (string, error) {
	if appTypeErr := validateAppType(appType); appTypeErr != nil {
		return "", appTypeErr
	}

	if addrTypeErr := validateAddress(wallet); addrTypeErr != nil {
		return "", addrTypeErr
	}

	var deeplink string
	if appType == "metamask" {
		deeplink = BuildQRCodeDeeplink(QRCodeData{
			AppType:   "metamask",
			Wallet:    wallet,
			ChainId:   1,
			Amount:    toWeiAsExponent(amount, tokenType), // hardcoded
			TokenType: tokenType,
		}, nil)
	} else {
		deeplink = BuildQRCodeDeeplink(QRCodeData{
			AppType:   "trust",
			Wallet:    wallet,
			ChainId:   1,
			Amount:    amount, // hardcoded
			TokenType: tokenType,
		}, nil)
	}

	filename := fmt.Sprintf("%s.png", uuid.New().String())
	log.Println("GenerateQrCode: detecting new request for qrcode, starting encoding...", filename)
	png, eErr := qrcode.Encode(deeplink, qrcode.Medium, 256)

	if eErr != nil {
		return filename, eErr
	}

	wd, gErr := os.Getwd()

	if gErr != nil {
		return filename, gErr
	}

	targetPath := strings.Join([]string{wd, "assets", "qrcode", filename}, "/")
	wErr := os.WriteFile(targetPath, png, constant.FilePermUserReadWriteGroupRead)

	if wErr != nil {
		return filename, wErr
	}

	return filename, nil
}
