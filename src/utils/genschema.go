package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	basePath := "./tmp/"
	table := "pagopendiente"
	resultTables := consultar("SELECT table_name FROM Information_Schema.TABLES WHERE table_name='" + table + "'")
	// resultTables := consultar("select * from requ_organization")
	// fmt.Println("resultTables: ", resultTables)
	for _, v := range resultTables {
		var codigo_struct string
		tableName := v["table_name"].(string)

		temp_table_name := strings.Split(tableName, "_")
		modulo := ""
		table := ""
		if len(temp_table_name) > 1 {
			switch temp_table_name[0] {
			case "Caja":
				codigo_struct += "package caja\n"
				// basePath += "caja/"
			case "Fina":
				codigo_struct += "package financiera\n"
				// basePath += "financiera/"
			case "Requ":
				codigo_struct += "package required\n"
				// basePath += "required/"
			case "Cont":
				codigo_struct += "package contabilidad\n"
				// basePath += "contabilidad/"
			case "Stck":
				codigo_struct += "package stock\n"
				// basePath += "stock/"
			case "Cost":
				codigo_struct += "package costos\n"
				// basePath += "costos/"
			case "Ctae":
				codigo_struct += "package cuentaCorriente\n"
				// basePath += "cuentaCorriente/"
			case "Rrhh":
				codigo_struct += "package recursosHumanos\n"
				// basePath += "recursosHumanos/"
			default:
				codigo_struct += "package tables\n"
				// basePath += "tables/"
				modulo = ""
				table = temp_table_name[0]

			}
			modulo = temp_table_name[0]
			table = temp_table_name[1]
		} else {
			codigo_struct += "package tables\n"
			basePath += "tables/"
			modulo = ""
			table = temp_table_name[0]
		}
		codigo_struct += "import \"server-go/src/models\" \n"
		codigo_struct += fmt.Sprintf("func %s_GetSchema() ([]models.Base, string) {\n", table)
		codigo_struct += fmt.Sprintf("\tvar %s []models.Base\n", table)
		codigo_struct += fmt.Sprintf("\ttableName := \"%s_\" + \"%s\"\n", modulo, table)

		query_sql := fmt.Sprintf("SELECT * FROM Information_Schema.Columns where  TABLE_NAME='%s'", tableName)
		var resultColumns []map[string]interface{}
		resultColumns = consultar(query_sql)
		// fmt.Println(resultColumns)
		var codigo_schema string
		fmt.Println("tabla:", tableName, "   columnas:", len(resultColumns))

		for _, column := range resultColumns {
			codigo_struct += fmt.Sprintf("\t%s = append(%s, models.Base{//%s\n", table, table, column["column_name"])
			codigo_struct += fmt.Sprintf("\t\tName:\"%s\",\n", column["column_name"])
			codigo_struct += fmt.Sprintf("\t\tDescription:\"%s\",\n", column["column_name"])
			codigo_struct += "\t\tRequired: true,\n"
			codigo_struct += "\t\tUpdate: true,\n"
			if column["udt_name"] == "varchar" || column["udt_name"] == "char" {
				codigo_struct += fmt.Sprintf("\t\tType:\"%s\",\n", "string")
				codigo_struct += "\t\tStrings: models.Strings{\n"
				codigo_struct += "\t\t\tExpr:      *models.Null(),\n"
				max_length := int(column["character_maximum_length"].(int64))
				if max_length == 10 && column["udt_name"] == "varchar" {
					codigo_struct += fmt.Sprintf("\t\t\tDate:%v,\n", true)
				} else {
					if max_length != 36 {
						if column["udt_name"] == "char" {
							codigo_struct += fmt.Sprintf("\t\t\tMin:%d,\n", max_length)
							codigo_struct += fmt.Sprintf("\t\t\tMax:%d,\n", max_length)
						} else {
							codigo_struct += fmt.Sprintf("\t\t\tMin:%f,\n", float64(max_length)*0.1)
							codigo_struct += fmt.Sprintf("\t\t\tMax:%d,\n", max_length)
							codigo_struct += fmt.Sprintf("\t\t\tUpperCase:%v,\n", true)

						}

					}
				}
				codigo_struct += "\t\t},\n"
			} else if column["data_type"] == "integer" {
				codigo_struct += fmt.Sprintf("\t\tType:\"%s\",\n", "uint64")
				codigo_struct += "\t\tUint: models.Uints{\n"
				codigo_struct += "\t\t\tMax: 10,\n"
				codigo_struct += "\t\t},\n"
			} else if column["udt_name"] == "numeric" || column["udt_name"] == "real" {
				codigo_struct += fmt.Sprintf("\t\tType:\"%s\",\n", "float64")
				codigo_struct += "\t\tFloat: models.Floats{\n"
				codigo_struct += "\t\t},\n"
			}
			codigo_struct += "\t})\n"
		}
		codigo_struct += fmt.Sprintf("\treturn %s, tableName\n", table)
		codigo_struct += "}\n"
		codigo_schema += codigo_struct
		texto := []byte(codigo_struct)
		errs := ioutil.WriteFile(fmt.Sprintf("%s%s.go", basePath, table), texto, 0644)
		if errs != nil {
			log.Fatal(errs)
		}

	}
}
func consultar(query string) []map[string]interface{} {
	db := Connection()
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Query: ", query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	cols, _ := rows.Columns()
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			log.Fatal(err)
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			l := *val

			if l != nil {
				if strings.Contains(reflect.TypeOf(l).String(), "[]uint8") {
					m[colName] = string(l.([]uint8))
				} else {
					m[colName] = l
				}
			} else {
				m[colName] = l
			}

		}

		result = append(result, m)
	}
	return result
}

func Connection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error configuraciÃ³n de variables de entorno")
	}
	server := os.Getenv("ENV_DDBB_SERVER")
	user := os.Getenv("ENV_DDBB_USER")
	password := os.Getenv("ENV_DDBB_PASSWORD")
	port := os.Getenv("ENV_DDBB_PORT")
	database := "Platcont"
	// fmt.Println("database: ", database)
	connection_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", server, port, user, password, database)
	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		log.Fatal("Error connection: ", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error creating connection: ", err.Error())
	}
	return db
}