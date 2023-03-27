package tables

import (
	"platcont/src/database/models"
)

func Clientproducts_GetSchema() ([]models.Base, string) {
	var clientproducts []models.Base
	tableName := "clientproducts"
	clientproducts = append(clientproducts, models.Base{ //id_clipd
		Name:        "id_clipd",
		Description: "id_clipd",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	clientproducts = append(clientproducts, models.Base{ //id_clie
		Name:        "id_clie",
		Description: "id_clie",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	clientproducts = append(clientproducts, models.Base{ //multi
		Name:        "multi",
		Description: "multi",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	clientproducts = append(clientproducts, models.Base{ //host
		Name:        "host",
		Description: "host",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       100,
			LowerCase: true,
		},
	})
	clientproducts = append(clientproducts, models.Base{ //users
		Name:        "users",
		Description: "users",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       20,
			LowerCase: true,
		},
	})
	clientproducts = append(clientproducts, models.Base{ //data_base
		Name:        "data_base",
		Description: "data_base",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       2,
			Max:       20,
			LowerCase: true,
		},
	})
	clientproducts = append(clientproducts, models.Base{ //password
		Name:        "password",
		Description: "password",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       8,
			Max:       60,
			LowerCase: true,
		},
	})
	clientproducts = append(clientproducts, models.Base{ //modulos
		Name:        "modulos",
		Description: "modulos",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       4,
			Max:       50,
			LowerCase: true,
		},
	})
	clientproducts = append(clientproducts, models.Base{ //date_facture
		Name:        "date_facture",
		Description: "date_facture",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	return clientproducts, tableName
}
