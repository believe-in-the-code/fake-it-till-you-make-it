package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtHexColorLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.HexColor",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.HexColor
		item.Scan(fake.HexColor())
		return &item
	}

	return &liar
}
