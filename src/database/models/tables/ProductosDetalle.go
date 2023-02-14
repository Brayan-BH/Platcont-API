package tables

import (
	"platcont/src/database/models"

	"github.com/google/uuid"
)

func Productosdetalle_GetSchema() ([]models.Base, string) {
	var productosdetalle []models.Base
	tableName := "productosdetalle"
	productosdetalle = append(productosdetalle, models.Base{ //id_detail
		Name:        "id_detail",
		Description: "id_detail",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	productosdetalle = append(productosdetalle, models.Base{ //s_impo
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Update:      true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	productosdetalle = append(productosdetalle, models.Base{ //l_detalle
		Name:        "l_detalle",
		Description: "l_detalle",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       100,
			LowerCase: true,
		},
	})
	productosdetalle = append(productosdetalle, models.Base{ //id_fact
		Name:        "id_fact",
		Description: "id_fact",
		Important:   true,
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	productosdetalle = append(productosdetalle, models.Base{ //id_prod
		Name:        "id_prod",
		Description: "id_prod",
		Default:     uuid.New().String(),
		Important:   true,
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	return productosdetalle, tableName
}
