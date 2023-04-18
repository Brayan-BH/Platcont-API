package tables

import (
	"github.com/deybin/go_basic_orm"
	"github.com/google/uuid"
)

func Users_GetSchema() ([]go_basic_orm.Model, string) {
	var users []go_basic_orm.Model
	tableName := "users"
	users = append(users, go_basic_orm.Model{ //id_user
		Name:        "id_user",
		Description: "id_user",
		Required:    true,
		Important:   true,
		Default:     uuid.New().String(),
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	users = append(users, go_basic_orm.Model{ //email
		Name:        "email",
		Description: "email",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       10,
			Max:       100,
			LowerCase: true,
		},
	})
	users = append(users, go_basic_orm.Model{ //password
		Name:        "password",
		Description: "password",
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Encriptar: true,
		},
	})
	users = append(users, go_basic_orm.Model{ //password_admin
		Name:        "password_admin",
		Description: "password_admin",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Encriptar: true,
		},
	})
	return users, tableName
}
