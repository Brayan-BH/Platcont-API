package tables

import (
	"platcont/src/database/models"

	"github.com/google/uuid"
)

func Productosdetalle_GetSchema() ([]models.Base, string) {
	var productosdetalle []models.Base
	tableName := "productosdetalle"
	productosdetalle = append(productosdetalle, models.Base{ //s_impo
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Update:      true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	productosdetalle = append(productosdetalle, models.Base{ //months
		Name:        "months",
		Description: "months",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	productosdetalle = append(productosdetalle, models.Base{ //years
		Name:        "years",
		Description: "years",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	productosdetalle = append(productosdetalle, models.Base{ //id_pddt
		Name:        "id_pddt",
		Description: "id_pddt",
		Required:    true,
		Default:     uuid.New().String(),
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	productosdetalle = append(productosdetalle, models.Base{ //id_fact
		Name:        "id_fact",
		Description: "id_fact",
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	productosdetalle = append(productosdetalle, models.Base{ //id_clipd
		Name:        "id_clipd",
		Description: "id_clipd",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	productosdetalle = append(productosdetalle, models.Base{ //l_deta
		Name:        "l_deta",
		Description: "l_deta",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       100,
			UpperCase: true,
		},
	})
	return productosdetalle, tableName
}
