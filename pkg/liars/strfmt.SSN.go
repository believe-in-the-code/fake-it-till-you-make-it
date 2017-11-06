package liars

import (
	"fmt"

	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtSSNLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.SSN",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.SSN
		item.Scan(fmt.Sprintf("%s-%s-%s", fake.DigitsN(3), fake.DigitsN(2), fake.DigitsN(4)))
		return &item
	}

	return &liar
}
