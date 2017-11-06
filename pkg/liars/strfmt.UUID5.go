package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
	"github.com/satori/go.uuid"
)

func NewStrfmtUUID5Liar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.UUID5",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.UUID5
		NamespaceDNS, _ := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
		item.Scan(uuid.NewV5(NamespaceDNS, fake.DomainName()))
		return &item
	}

	return &liar
}
