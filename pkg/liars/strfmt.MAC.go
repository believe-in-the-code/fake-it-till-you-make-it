package liars

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
)

func NewStrfmtMACLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.MAC",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.MAC
		item.Scan(randomdata.MacAddress())
		return &item
	}

	return &liar
}
