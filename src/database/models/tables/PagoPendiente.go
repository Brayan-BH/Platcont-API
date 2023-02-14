package tables

import (
	"platcont/src/database/models"

	"github.com/google/uuid"
)

func PagoPendiente_GetSchema() ([]models.Base, string) {
	var pagopendiente []models.Base
	tableName := "pagopendiente"
	pagopendiente = append(pagopendiente, models.Base{ //n_fact
		Name:        "n_fact",
		Description: "n_fact",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	pagopendiente = append(pagopendiente, models.Base{ //n_period
		Name:        "n_period",
		Description: "n_period",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	pagopendiente = append(pagopendiente, models.Base{ //s_impo
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Update:      true,
		Type:        "float64",
		Float:       models.Floats{},
	})
	pagopendiente = append(pagopendiente, models.Base{ //id_pago
		Name:        "id_pago",
		Description: "id_pago",
		Default:     uuid.New().String(),
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	pagopendiente = append(pagopendiente, models.Base{ //n_clie
		Name:        "n_clie",
		Description: "n_clie",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       100,
			UpperCase: true,
		},
	})
	pagopendiente = append(pagopendiente, models.Base{ //l_detalle
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
	return pagopendiente, tableName
}
