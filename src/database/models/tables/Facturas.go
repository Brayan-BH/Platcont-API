package tables

import (
	"platcont/src/database/models"

	"github.com/google/uuid"
)

func facturas_GetSchema() ([]models.Base, string) {
	var facturas []models.Base
	tableName := "facturas"
	facturas = append(facturas, models.Base{ //id_fact
		Name:        "id_fact",
		Description: "id_fact",
		Required:    true,
		Default:     uuid.New().String(),
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	facturas = append(facturas, models.Base{ //n_fact
		Name:        "n_fact",
		Description: "n_fact",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	facturas = append(facturas, models.Base{ //s_impo
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Update:      true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	facturas = append(facturas, models.Base{ //k_stad
		Name:        "k_stad",
		Description: "k_stad",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	facturas = append(facturas, models.Base{ //n_period
		Name:        "n_period",
		Description: "n_period",
		Required:    true,
		Update:      true,
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	facturas = append(facturas, models.Base{ //id_clie
		Name:        "id_clie",
		Description: "id_clie",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	facturas = append(facturas, models.Base{ //n_clie
		Name:        "n_clie",
		Description: "n_clie",
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
	facturas = append(facturas, models.Base{ //l_detalle
		Name:        "l_detalle",
		Description: "l_detalle",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       8,
			Max:       80,
			UpperCase: true,
		},
	})
	return facturas, tableName
}
