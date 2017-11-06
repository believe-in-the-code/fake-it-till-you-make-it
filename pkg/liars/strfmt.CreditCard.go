package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtCreditCardLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.CreditCard",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.CreditCard
		item.Scan(fake.CreditCardNum(""))
		return &item
	}

	return &liar
}
