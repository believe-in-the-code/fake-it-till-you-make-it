package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtPasswordLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.Password",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.Password
		item.Scan(fake.SimplePassword())
		return &item
	}

	return &liar
}
