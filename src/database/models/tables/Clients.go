package tables

import (
	"platcont/src/database/models"

	"github.com/google/uuid"
)

func Clients_GetSchema() ([]models.Base, string) {
	var clients []models.Base
	tableName := "clients"
	clients = append(clients, models.Base{ //id
		Name:        "id",
		Description: "id",
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	clients = append(clients, models.Base{ //l_clie
		Name:        "l_clie",
		Description: "l_clie",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       100,
			LowerCase: true,
		},
	})
	clients = append(clients, models.Base{ //n_docu
		Name:        "n_docu",
		Description: "n_docu",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       8,
			Max:       11,
			LowerCase: true,
		},
	})
	clients = append(clients, models.Base{ //l_orga
		Name:        "l_orga",
		Description: "l_orga",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       15,
			Max:       150,
			LowerCase: true,
		},
	})
	clients = append(clients, models.Base{ //l_dire
		Name:        "l_dire",
		Description: "l_dire",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       20,
			Max:       200,
			LowerCase: true,
		},
	})
	clients = append(clients, models.Base{ //n_celu
		Name:        "n_celu",
		Description: "n_celu",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       2,
			Max:       25,
			LowerCase: true,
		},
	})
	clients = append(clients, models.Base{ //l_emai
		Name:        "l_emai",
		Description: "l_emai",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       15,
			Max:       150,
			LowerCase: true,
		},
	})
	clients = append(clients, models.Base{ //n_repr
		Name:        "n_repr",
		Description: "n_repr",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       0,
			Max:       8,
			LowerCase: true,
		},
	})
	clients = append(clients, models.Base{ //l_repr
		Name:        "l_repr",
		Description: "l_repr",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       7,
			Max:       70,
			LowerCase: true,
		},
	})
	clients = append(clients, models.Base{ //uid
		Name:        "uid",
		Description: "uid",
		Default:     uuid.New().String(),
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	clients = append(clients, models.Base{ //id_prod
		Name:        "id_prod",
		Description: "id_prod",
		Important:   true,
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	return clients, tableName
}
