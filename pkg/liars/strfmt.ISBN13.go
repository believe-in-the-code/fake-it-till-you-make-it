package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
)

func NewStrfmtISBN13Liar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.ISBN13",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.ISBN13
		item.Scan("")
		return &item
	}

	return &liar
}
