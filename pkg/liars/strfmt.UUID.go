package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func NewStrfmtUUIDLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.UUID",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.UUID
		item.Scan(uuid.NewV1())
		return &item
	}

	return &liar
}
