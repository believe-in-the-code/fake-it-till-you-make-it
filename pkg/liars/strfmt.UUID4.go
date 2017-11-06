package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func NewStrfmtUUID4Liar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.UUID4",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.UUID5
		item.Scan(uuid.NewV4())
		return &item
	}

	return &liar
}
