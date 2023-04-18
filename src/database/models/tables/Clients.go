package tables

import (
	"github.com/deybin/go_basic_orm"
)

func Clients_GetSchema() ([]go_basic_orm.Model, string) {
	var clients []go_basic_orm.Model
	tableName := "clients"
	clients = append(clients, go_basic_orm.Model{ //id_clie
		Name:        "id_clie",
		Description: "id_clie",
		Required:    true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	clients = append(clients, go_basic_orm.Model{ //n_docu
		Name:        "n_docu",
		Description: "n_docu",
		Required:    true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       8,
			Max:       11,
			LowerCase: true,
		},
	})
	clients = append(clients, go_basic_orm.Model{ //l_orga
		Name:        "l_orga",
		Description: "l_orga",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       10,
			Max:       150,
			LowerCase: true,
		},
	})
	clients = append(clients, go_basic_orm.Model{ //l_dire
		Name:        "l_dire",
		Description: "l_dire",
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       10,
			Max:       200,
			LowerCase: true,
		},
	})
	clients = append(clients, go_basic_orm.Model{ //n_celu
		Name:        "n_celu",
		Description: "n_celu",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       9,
			Max:       25,
			LowerCase: true,
		},
	})
	clients = append(clients, go_basic_orm.Model{ //l_emai
		Name:        "l_emai",
		Description: "l_emai",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       10,
			Max:       50,
			LowerCase: true,
		},
	})
	clients = append(clients, go_basic_orm.Model{ //n_repr
		Name:        "n_repr",
		Description: "n_repr",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       0,
			Max:       8,
			LowerCase: true,
		},
	})
	clients = append(clients, go_basic_orm.Model{ //l_repr
		Name:        "l_repr",
		Description: "l_repr",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       7,
			Max:       70,
			LowerCase: true,
		},
	})
	return clients, tableName
}
