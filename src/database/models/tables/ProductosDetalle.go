package tables

import (
	"github.com/deybin/go_basic_orm"
	"github.com/google/uuid"
)

func Productosdetalle_GetSchema() ([]go_basic_orm.Model, string) {
	var productosdetalle []go_basic_orm.Model
	tableName := "productosdetalle"
	productosdetalle = append(productosdetalle, go_basic_orm.Model{ //s_impo
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Update:      true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	productosdetalle = append(productosdetalle, go_basic_orm.Model{ //months
		Name:        "months",
		Description: "months",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	productosdetalle = append(productosdetalle, go_basic_orm.Model{ //years
		Name:        "years",
		Description: "years",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	productosdetalle = append(productosdetalle, go_basic_orm.Model{ //id_pddt
		Name:        "id_pddt",
		Description: "id_pddt",
		Required:    true,
		Default:     uuid.New().String(),
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	productosdetalle = append(productosdetalle, go_basic_orm.Model{ //id_fact
		Name:        "id_fact",
		Description: "id_fact",
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	productosdetalle = append(productosdetalle, go_basic_orm.Model{ //id_clipd
		Name:        "id_clipd",
		Description: "id_clipd",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	productosdetalle = append(productosdetalle, go_basic_orm.Model{ //l_deta
		Name:        "l_deta",
		Description: "l_deta",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       10,
			Max:       100,
			UpperCase: true,
		},
	})
	return productosdetalle, tableName
}
