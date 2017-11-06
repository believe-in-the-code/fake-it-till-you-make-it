package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
)

func NewStrfmtISBN10Liar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.ISBN10",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.ISBN10
		item.Scan("")
		return &item
	}

	return &liar
}
