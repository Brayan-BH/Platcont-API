package tables

import (
	"github.com/deybin/go_basic_orm"
	"github.com/google/uuid"
)

func Facturas_GetSchema() ([]go_basic_orm.Model, string) {
	var facturas []go_basic_orm.Model
	tableName := "facturas"
	facturas = append(facturas, go_basic_orm.Model{ //id_fact
		Name:        "id_fact",
		Description: "id_fact",
		Required:    true,
		Default:     uuid.New().String(),
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	facturas = append(facturas, go_basic_orm.Model{ //id_clipd
		Name:        "id_clipd",
		Description: "id_clipd",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	facturas = append(facturas, go_basic_orm.Model{ //s_impo
		Name:        "s_impo",
		Description: "s_impo",
		Required:    true,
		Update:      true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	facturas = append(facturas, go_basic_orm.Model{ //s_igv
		Name:        "s_igv",
		Description: "s_igv",
		Required:    true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	facturas = append(facturas, go_basic_orm.Model{ //s_desc
		Name:        "s_desc",
		Description: "s_desc",
		Required:    true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	facturas = append(facturas, go_basic_orm.Model{ //s_tota
		Name:        "s_tota",
		Description: "s_tota",
		Required:    true,
		Update:      true,
		Type:        "float64",
		Float:       go_basic_orm.Floats{},
	})
	facturas = append(facturas, go_basic_orm.Model{ //years
		Name:        "years",
		Description: "years",
		Required:    true,
		Type:        "uint64",
	})
	facturas = append(facturas, go_basic_orm.Model{ //k_stad
		Name:        "k_stad",
		Description: "k_stad",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	facturas = append(facturas, go_basic_orm.Model{ //months
		Name:        "months",
		Description: "months",
		Required:    true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	facturas = append(facturas, go_basic_orm.Model{ //f_pago
		Name:        "f_pago",
		Description: "f_pago",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Date: true,
		},
	})
	facturas = append(facturas, go_basic_orm.Model{ //l_obse
		Name:        "l_obse",
		Description: "l_obse",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       10,
			Max:       100,
			LowerCase: true,
		},
	})
	facturas = append(facturas, go_basic_orm.Model{ //n_docu
		Name:        "n_docu",
		Description: "n_docu",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       1,
			Max:       11,
			LowerCase: true,
		},
	})
	facturas = append(facturas, go_basic_orm.Model{ //c_comp
		Name:        "c_comp",
		Description: "c_comp",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       0,
			Max:       2,
			LowerCase: true,
		},
	})
	facturas = append(facturas, go_basic_orm.Model{ //n_seri
		Name:        "n_seri",
		Description: "n_seri",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       0,
			Max:       4,
			LowerCase: true,
		},
	})
	facturas = append(facturas, go_basic_orm.Model{ //n_com
		Name:        "n_com",
		Description: "n_com",
		Required:    true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	facturas = append(facturas, go_basic_orm.Model{ //f_venc
		Name:        "f_venc",
		Description: "f_venc",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Date: true,
		},
	})
	facturas = append(facturas, go_basic_orm.Model{ //f_comp
		Name:        "f_comp",
		Description: "f_comp",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Date: true,
		},
	})
	return facturas, tableName
}
