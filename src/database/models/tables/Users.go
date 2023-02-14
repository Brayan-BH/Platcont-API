package tables

import (
	"platcont/src/database/models"

	"github.com/google/uuid"
)

func Users_GetSchema() ([]models.Base, string) {
	var users []models.Base
	tableName := "users"
	users = append(users, models.Base{ //id
		Name:        "id",
		Description: "id",
		Default:     uuid.New().String(),
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	users = append(users, models.Base{ //email
		Name:        "email",
		Description: "email",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       15,
			Max:       100,
			LowerCase: true,
		},
	})
	users = append(users, models.Base{ //password
		Name:        "password",
		Description: "password",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       200,
			LowerCase: true,
		},
	})
	users = append(users, models.Base{ //password_admin
		Name:        "password_admin",
		Description: "password_admin",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       200,
			LowerCase: true,
		},
	})
	return users, tableName
}
