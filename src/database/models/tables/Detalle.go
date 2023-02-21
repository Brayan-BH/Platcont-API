package tables

import "platcont/src/database/models"

func FacturasDetalle_GetSchema() ([]models.Base, string) {
	var detalle []models.Base
	tableName := "facturas_" + "detalle"
	detalle = append(detalle, models.Base{ //s_desc
		Name:        "s_desc",
		Description: "s_desc",
		Required:    true,
		Update:      true,
	})
	detalle = append(detalle, models.Base{ //s_tota
		Name:        "s_tota",
		Description: "s_tota",
		Required:    true,
		Update:      true,
	})
	detalle = append(detalle, models.Base{ //n_item
		Name:        "n_item",
		Description: "n_item",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	detalle = append(detalle, models.Base{ //s_impo
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Update:      true,
	})
	detalle = append(detalle, models.Base{ //s_igv
		Name:        "s_igv",
		Description: "s_igv",
		Required:    true,
		Update:      true,
	})
	detalle = append(detalle, models.Base{ //id_pddt
		Name:        "id_pddt",
		Description: "id_pddt",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	detalle = append(detalle, models.Base{ //c_prod
		Name:        "c_prod",
		Description: "c_prod",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       0,
			Max:       8,
			UpperCase: true,
		},
	})
	detalle = append(detalle, models.Base{ //l_peri
		Name:        "l_peri",
		Description: "l_peri",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       0,
			Max:       7,
			UpperCase: true,
		},
	})
	detalle = append(detalle, models.Base{ //id_fact
		Name:        "id_fact",
		Description: "id_fact",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	return detalle, tableName
}
