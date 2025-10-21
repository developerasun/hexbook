package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	hook "github.com/hexbook/pkg"
)

func TestQRCodeURI(t *testing.T) {
	assert := assert.New(t)
	supported := "metamask"
	unsupported := "binance"

	assert.Panics(func() {
		hook.BuildBaseUrlByWallet(unsupported)
	})

	var baseUrl string
	assert.NotPanics(func() {
		baseUrl = hook.BuildBaseUrlByWallet(supported)
	})

	t.Log("baseUrl: ", baseUrl)
}

func TestMetamaskDeeplink(t *testing.T) {
	assert := assert.New(t)
	generated := "https://metamask.app.link/send/pay-0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11@1?value=1e15"
	qd := hook.QRCodeData{
		AppType:   "metamask",
		Wallet:    "0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11",
		ChainId:   1,
		Amount:    "1e15",
		TokenType: "eth",
		Decimal:   18,
	}
	// eip681 scheme
	options := hook.UriOption{Prefix: "pay"}
	deeplink := hook.BuildQRCodeDeeplink(qd, &options)

	assert.Equal(generated, deeplink)
	t.Log(deeplink)
}
