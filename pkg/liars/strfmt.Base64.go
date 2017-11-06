package liars

import (
	"encoding/base64"

	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
)

func NewStrfmtBase64Liar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "string",
		Type: "strfmt.Base64",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.Base64
		item.Scan(base64.StdEncoding.EncodeToString([]byte(fake.Words())))
		return &item
	}

	return &liar
}
