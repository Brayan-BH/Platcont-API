package tables

import (
	"platcont/src/database/models"

	"github.com/google/uuid"
)

func ProductosDetalle_GetSchema() ([]models.Base, string) {
	var productosdetalle []models.Base
	tableName := "productosdetalle"
	productosdetalle = append(productosdetalle, models.Base{ //id_detail
		Name:        "id_detail",
		Description: "id_detail",
		Required:    true,
		Default:     uuid.New().String(),
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
			Min:       10.000000,
			Max:       100,
			UpperCase: true,
		},
	})
	productosdetalle = append(productosdetalle, models.Base{ //id_fact
		Name:        "id_fact",
		Description: "id_fact",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	return productosdetalle, tableName
}
