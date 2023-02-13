package tables

import "platcont/src/database/models"

func users_GetSchema() ([]models.Base, string) {
	var users []models.Base
	tableName := "_" + "users"
	users = append(users, models.Base{ //id
		Name:        "id",
		Description: "id",
		Required:    true,
		Update:      true,
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
			Min:       10.000000,
			Max:       100,
			UpperCase: true,
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
			Min:       20.000000,
			Max:       200,
			UpperCase: true,
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
			Min:       20.000000,
			Max:       200,
			UpperCase: true,
		},
	})
	return users, tableName
}
