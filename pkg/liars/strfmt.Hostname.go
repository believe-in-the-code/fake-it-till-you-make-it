package liars

import (
	"strings"

	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtHostnameLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.Hostname",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.Hostname
		item.Scan(strings.ToLower(fake.DomainName()))
		return &item
	}

	return &liar
}
