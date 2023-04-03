package tables

import "platcont/src/database/models"

func Versiones_GetSchema() ([]models.Base, string) {
	var versiones []models.Base
	tableName := "versiones"
	versiones = append(versiones, models.Base{ //id_version
		Name:        "id_version",
		Description: "id_version",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	versiones = append(versiones, models.Base{ //f_digi
		Name:        "f_digi",
		Description: "f_digi",
		Required:    true,
		Update:      true,
	})
	versiones = append(versiones, models.Base{ //c_vers
		Name:        "c_vers",
		Description: "c_vers",
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
	versiones = append(versiones, models.Base{ //id_file
		Name:        "id_file",
		Description: "id_file",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	versiones = append(versiones, models.Base{ //l_deta
		Name:        "l_deta",
		Description: "l_deta",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       50,
			Max:       500,
			UpperCase: true,
		},
	})
	return versiones, tableName
}
