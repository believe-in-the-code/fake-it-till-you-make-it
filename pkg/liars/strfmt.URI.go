package liars

import (
	"fmt"
	"strings"

	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtURILiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.URI",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.Hostname
		item.Scan(fmt.Sprintf("https://%s/", strings.ToLower(fake.DomainName())))
		return &item
	}

	return &liar
}
