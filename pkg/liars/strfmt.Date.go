package liars

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/believe-in-the-code/liar/pkg"
	"github.com/go-openapi/strfmt"
)

func NewStrfmtDateLiar() *pkg.Liar {
	liar := pkg.Liar{
		Kind: "struct",
		Type: "strfmt.Date",
	}

	liar.Fill = func(args pkg.Tag) interface{} {
		var item strfmt.Date
		dateTime, _ := strfmt.ParseDateTime(randomdata.FullDate())
		item.Scan(dateTime.String())
		return &item
	}

	return &liar
}
