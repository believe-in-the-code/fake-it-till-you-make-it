package liars

import (
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
)

func NewStrfmtDurationLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "int64",
		Type: "strfmt.Duration",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.Duration
		duration := "1d"
		if args.HasIndex(0) {
			duration = args.GetIndex(0)
		}
		item.Scan(duration)
		return &item
	}

	return &liar
}
