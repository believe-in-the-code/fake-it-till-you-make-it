package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtIPv4Liar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.IPv4",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.IPv4
		item.Scan(fake.IPv4())
		return &item
	}

	return &liar
}
