package liars

import (
	"encoding/hex"
	"fmt"

	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtRGBColorLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.RGBColor",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.RGBColor

		rgbHex := []byte(fake.HexColor())
		red, err := hex.DecodeString(string(rgbHex[0:1]))
		if err != nil {
			panic(err)
		}
		green, err := hex.DecodeString(string(rgbHex[2:3]))
		if err != nil {
			panic(err)
		}
		blue, err := hex.DecodeString(string(rgbHex[4:5]))
		if err != nil {
			panic(err)
		}

		item.Scan(fmt.Sprintf("rgb(%s, %s, %s)", red, green, blue))
		return &item
	}

	return &liar
}
