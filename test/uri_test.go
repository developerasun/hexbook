package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	hook "github.com/hexbook/pkg"
)

func TestQRCodeURI(t *testing.T) {
	t.Skip()
	assert := assert.New(t)
	supported := "metamask"
	unsupported := "binance"

	assert.Panics(func() {
		hook.BuildBaseUrlByAppType(unsupported)
	})

	var baseUrl string
	assert.NotPanics(func() {
		baseUrl = hook.BuildBaseUrlByAppType(supported)
	})

	t.Log("baseUrl: ", baseUrl)
}

func TestMetamaskDeeplink(t *testing.T) {
	t.Skip()
	assert := assert.New(t)
	generated := "https://metamask.app.link/send/pay-0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1?value=1e15"
	qd := hook.QRCodeData{
		AppType:   "metamask",
		Wallet:    "0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11",
		ChainId:   1,
		Amount:    "1e15",
		TokenType: "eth",
	}
	// eip681 scheme
	options := hook.UriOption{Prefix: "pay"}
	deeplink := hook.BuildQRCodeDeeplink(qd, &options)

	assert.Equal(generated, deeplink)
	t.Log(deeplink)
}

/*
@doc see here: https://goethereumbook.org/util-go/
*/
func TestToWeiConversion(t *testing.T) {
	t.Skip()
	assert := assert.New(t)
	hardcodedAmount := "0.002"
	amount, _ := decimal.NewFromString(hardcodedAmount)
	plain := decimal.NewFromFloat(1e18)
	calculated := amount.Mul(plain)

	t.Log("calculated: ", calculated)
	assert.Equal(calculated.String(), "2000000000000000")
}

func TestToWeiWithExponent(t *testing.T) {
	t.Skip()
	assert := assert.New(t)
	toFloat, _ := strconv.ParseFloat("0.004", 64)
	target := fmt.Sprintf("%.e", toFloat*1e18)

	converted := strings.Replace(target, "+", "", 1)
	t.Log("target: ", target)
	t.Log("toFloat: ", toFloat)
	t.Log("converted: ", converted)
	assert.Equal("4e15", converted)
}

func TestUniqueFilename(t *testing.T) {
	assert := assert.New(t)
	filename := fmt.Sprintf("%s.png", uuid.New().String())

	t.Log("filename: ", filename)
	assert.Greater(len(filename), 0)
}
