package tables

import (
	"github.com/deybin/go_basic_orm"
)

func Clientproducts_GetSchema() ([]go_basic_orm.Model, string) {
	var clientproducts []go_basic_orm.Model
	tableName := "clientproducts"
	clientproducts = append(clientproducts, go_basic_orm.Model{ //id_clipd
		Name:        "id_clipd",
		Description: "id_clipd",
		Required:    true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	clientproducts = append(clientproducts, go_basic_orm.Model{ //id_clie
		Name:        "id_clie",
		Description: "id_clie",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	clientproducts = append(clientproducts, go_basic_orm.Model{ //multi
		Name:        "multi",
		Description: "multi",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	clientproducts = append(clientproducts, go_basic_orm.Model{ //host
		Name:        "host",
		Description: "host",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       5,
			Max:       100,
			LowerCase: true,
		},
	})
	clientproducts = append(clientproducts, go_basic_orm.Model{ //users
		Name:        "users",
		Description: "users",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       5,
			Max:       20,
			LowerCase: true,
		},
	})
	clientproducts = append(clientproducts, go_basic_orm.Model{ //data_base
		Name:        "data_base",
		Description: "data_base",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       2,
			Max:       20,
			LowerCase: true,
		},
	})
	clientproducts = append(clientproducts, go_basic_orm.Model{ //password
		Name:        "password",
		Description: "password",
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       8,
			Max:       60,
			LowerCase: true,
		},
	})
	clientproducts = append(clientproducts, go_basic_orm.Model{ //modulos
		Name:        "modulos",
		Description: "modulos",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       4,
			Max:       50,
			LowerCase: true,
		},
	})
	clientproducts = append(clientproducts, go_basic_orm.Model{ //date_facture
		Name:        "date_facture",
		Description: "date_facture",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Date: true,
		},
	})
	return clientproducts, tableName
}
