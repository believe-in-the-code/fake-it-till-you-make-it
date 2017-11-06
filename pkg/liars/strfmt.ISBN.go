package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
)

func NewStrfmtISBNLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.ISBN",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.ISBN
		item.Scan("")
		return &item
	}

	return &liar
}
