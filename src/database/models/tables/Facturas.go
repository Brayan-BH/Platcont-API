package tables

import (
	"platcont/src/database/models"

	"github.com/google/uuid"
)

func Facturas_GetSchema() ([]models.Base, string) {
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
	facturas = append(facturas, models.Base{ //id_clipd
		Name:        "id_clipd",
		Description: "id_clipd",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	facturas = append(facturas, models.Base{ //s_impo
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Update:      true,
	})
	facturas = append(facturas, models.Base{ //s_igv
		Name:        "s_igv",
		Description: "s_igv",
		Required:    true,
		Update:      true,
	})
	facturas = append(facturas, models.Base{ //s_desc
		Name:        "s_desc",
		Description: "s_desc",
		Required:    true,
		Update:      true,
	})
	facturas = append(facturas, models.Base{ //s_tota
		Name:        "s_tota",
		Description: "s_tota",
		Required:    true,
		Update:      true,
	})
	facturas = append(facturas, models.Base{ //years
		Name:        "years",
		Description: "years",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
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
	facturas = append(facturas, models.Base{ //months
		Name:        "months",
		Description: "months",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	facturas = append(facturas, models.Base{ //f_pago
		Name:        "f_pago",
		Description: "f_pago",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	facturas = append(facturas, models.Base{ //l_obse
		Name:        "l_obse",
		Description: "l_obse",
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
	facturas = append(facturas, models.Base{ //n_docu
		Name:        "n_docu",
		Description: "n_docu",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       1,
			Max:       11,
			UpperCase: true,
		},
	})
	facturas = append(facturas, models.Base{ //c_comp
		Name:        "c_comp",
		Description: "c_comp",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       0,
			Max:       2,
			UpperCase: true,
		},
	})
	facturas = append(facturas, models.Base{ //n_seri
		Name:        "n_seri",
		Description: "n_seri",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       0,
			Max:       4,
			UpperCase: true,
		},
	})
	facturas = append(facturas, models.Base{ //n_com
		Name:        "n_com",
		Description: "n_com",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	facturas = append(facturas, models.Base{ //f_venc
		Name:        "f_venc",
		Description: "f_venc",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	facturas = append(facturas, models.Base{ //f_comp
		Name:        "f_comp",
		Description: "f_comp",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	return facturas, tableName
}
