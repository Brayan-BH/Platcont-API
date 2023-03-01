package tables

import (
	"platcont/src/database/models"

	"github.com/google/uuid"
)

func Clients_GetSchema() ([]models.Base, string) {
	var clients []models.Base
	tableName := "clients"
	clients = append(clients, models.Base{ //id_clie
		Name:        "id_clie",
		Description: "id_clie",
		Required:    true,
		Important:   true,
		Default:     uuid.New().String(),
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
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
			Min:       10,
			Max:       150,
			LowerCase: true,
		},
	})
	clients = append(clients, models.Base{ //l_dire
		Name:        "l_dire",
		Description: "l_dire",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       200,
			LowerCase: true,
		},
	})
	clients = append(clients, models.Base{ //n_celu
		Name:        "n_celu",
		Description: "n_celu",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       9,
			Max:       25,
			LowerCase: true,
		},
	})
	clients = append(clients, models.Base{ //l_emai
		Name:        "l_emai",
		Description: "l_emai",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       15,
			Max:       50,
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
	return clients, tableName
}
