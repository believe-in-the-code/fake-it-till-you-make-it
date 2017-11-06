package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtEmailLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.Email",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.Email
		item.Scan(fake.EmailAddress())
		return &item
	}

	return &liar
}
