package tables

import (
	"github.com/deybin/go_basic_orm"
)

func Versiones_GetSchema() ([]go_basic_orm.Model, string) {
	var versiones []go_basic_orm.Model
	tableName := "versiones"
	versiones = append(versiones, go_basic_orm.Model{ //id_version
		Name:        "id_version",
		Description: "id_version",
		Required:    true,
		Update:      true,
		Type:        "uint64",
		Uint: go_basic_orm.Uints{
			Max: 10,
		},
	})
	versiones = append(versiones, go_basic_orm.Model{ //f_digi
		Name:        "f_digi",
		Description: "f_digi",
		Required:    true,
		Update:      true,
	})
	versiones = append(versiones, go_basic_orm.Model{ //c_vers
		Name:        "c_vers",
		Description: "c_vers",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       10,
			Max:       100,
			UpperCase: true,
		},
	})
	versiones = append(versiones, go_basic_orm.Model{ //id_file
		Name:        "id_file",
		Description: "id_file",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings:     go_basic_orm.Strings{},
	})
	versiones = append(versiones, go_basic_orm.Model{ //l_deta
		Name:        "l_deta",
		Description: "l_deta",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: go_basic_orm.Strings{
			Min:       50,
			Max:       500,
			UpperCase: true,
		},
	})
	return versiones, tableName
}
