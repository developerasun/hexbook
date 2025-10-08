package test

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"testing"

	"github.com/google/uuid"
	constant "github.com/hexbook/internal/constant"
	qrcode "github.com/skip2/go-qrcode"
)

func TestCreateQrCodeAtPath(t *testing.T) {
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)

	if err != nil {
		log.Println(err.Error())
		t.FailNow()
	}

	wd, gErr := os.Getwd()

	if gErr != nil {
		log.Println(gErr.Error())
		t.FailNow()
	}

	id := uuid.New().String()
	root := path.Dir(wd)
	filename := fmt.Sprintf("%s.png", id)
	t.Log("root: ", root)
	t.Log("filename: ", filename)

	targetPath := strings.Join([]string{root, "assets", "qrcode", filename}, "/")
	wErr := os.WriteFile(targetPath, png, constant.FilePermUserReadWriteGroupRead)
	if wErr != nil {
		log.Println(wErr.Error())
		t.FailNow()
	}
}
