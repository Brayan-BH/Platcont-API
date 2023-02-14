package tables

import (
	"platcont/src/database/models"

	"github.com/google/uuid"
)

func Clientproducts_GetSchema() ([]models.Base, string) {
	var clientproducts []models.Base
	tableName := "clientproducts"
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
			Min:       10,
			Max:       100,
			UpperCase: true,
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
			Min:       2,
			Max:       20,
			UpperCase: true,
		},
	})
	clientproducts = append(clientproducts, models.Base{ //data_base
		Name:        "data_base",
		Description: "data_base",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       2,
			Max:       20,
			UpperCase: true,
		},
	})
	clientproducts = append(clientproducts, models.Base{ //password
		Name:        "password",
		Description: "password",
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
	clientproducts = append(clientproducts, models.Base{ //modulos
		Name:        "modulos",
		Description: "modulos",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       50,
			UpperCase: true,
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
	clientproducts = append(clientproducts, models.Base{ //uid
		Name:        "uid",
		Description: "uid",
		Default:     uuid.New().String(),
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	clientproducts = append(clientproducts, models.Base{ //id
		Name:        "id",
		Description: "id",
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	clientproducts = append(clientproducts, models.Base{ //id_pago
		Name:        "id_pago",
		Description: "id_pago",
		Important:   true,
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	return clientproducts, tableName
}
