package tables

import (
	"github.com/deybin/go_basic_orm"
)

func FacturasDetalle_GetSchema() ([]go_basic_orm.Model, string) {
	var detalle []go_basic_orm.Model
	tableName := "facturas_" + "detalle"
	detalle = append(detalle, go_basic_orm.Model{ //s_desc
		Name:        "s_desc",
		Description: "s_desc",
		Required:    true,
		Update:      true,
	})
	detalle = append(detalle, go_basic_orm.Model{ //s_tota
		Name:        "s_tota",
		Description: "s_tota",
		Required:    true,
		Update:      true,
	})
	detalle = append(detalle, go_basic_orm.Model{ //n_item
		Name:        "n_item",
		Description: "n_item",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	detalle = append(detalle, go_basic_orm.Model{ //s_impo
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Update:      true,
	})
	detalle = append(detalle, go_basic_orm.Model{ //s_igv
		Name:        "s_igv",
		Description: "s_igv",
		Required:    true,
		Update:      true,
	})
	detalle = append(detalle, go_basic_orm.Model{ //id_pddt
		Name:        "id_pddt",
		Description: "id_pddt",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	detalle = append(detalle, go_basic_orm.Model{ //c_prod
		Name:        "c_prod",
		Description: "c_prod",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       0,
			Max:       8,
			UpperCase: true,
		},
	})
	detalle = append(detalle, go_basic_orm.Model{ //l_peri
		Name:        "l_peri",
		Description: "l_peri",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       0,
			Max:       7,
			UpperCase: true,
		},
	})
	detalle = append(detalle, go_basic_orm.Model{ //id_fact
		Name:        "id_fact",
		Description: "id_fact",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	return detalle, tableName
}
