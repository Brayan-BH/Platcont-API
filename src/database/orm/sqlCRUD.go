package orm

import (
	"context"
	"errors"
	"fmt"
	"platcont/src/auth"
	"platcont/src/database/connection"
	"platcont/src/database/models"
	"platcont/src/libraries/cryptoAes"
	"platcont/src/libraries/date"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type SqlExec struct {
	Ob     []map[string]interface{} //datos para observación
	Data   []map[string]interface{} //datos para insertar o actualizar o eliminar
	Query  []map[string]interface{}
	Table  string
	action string
}

/**
 * Reception datos para crear query para insertar, actualizar o eliminar
 * datos {[]map[string]interface{}}: datos a insertar, actualizar o eliminar
 * name {string}: nombre de la tabla
 * returns {*SqlExec} retorna SqlExec struct
 */
func (sq *SqlExec) New(datos []map[string]interface{}, name string) *SqlExec {
	sq.Ob = datos
	sq.Table = name
	return sq
}

/**
 * Valida los datos para insertar y crea el query para insertar
 * schema {[]models.Base}: modelo de la tabla
 * returns {error}: retorna errores ocurridos en la validación
 */
func (sq *SqlExec) Insert(schema []models.Base) error {
	data := sq.Ob
	length := len(data)

	if length > 0 {
		var sqlExec = make([]map[string]interface{}, 0)
		var data_insert []map[string]interface{}

		for _, item := range data {
			preArray, err := _checkInsertSchema(schema, item)
			if err == nil {
				data_insert = append(data_insert, preArray)
				var lineSqlExec = make(map[string]interface{}, 2)
				sqlPreparateInsert := ""
				sqlPreparateValues := ""
				var i int
				var p uint64
				length_newMap := len(preArray)
				var valuesExec []interface{}
				char := "$"
				for k, v := range preArray {
					p++
					if i+1 < length_newMap {
						sqlPreparateInsert += k + ", "
						sqlPreparateValues += char + strconv.FormatUint(p, 10) + ", "
					} else {
						sqlPreparateInsert += k
						sqlPreparateValues += char + strconv.FormatUint(p, 10)
					}
					valuesExec = append(valuesExec, v)
					i++
				}

				sqlPreparate := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", sq.Table, sqlPreparateInsert, sqlPreparateValues)
				lineSqlExec["sqlPreparate"] = sqlPreparate
				lineSqlExec["valuesExec"] = valuesExec
				sqlExec = append(sqlExec, lineSqlExec)

			} else {
				return err
			}
		}
		sq.Query = sqlExec
		sq.Data = data_insert
	} else {
		return errors.New("No existen datos para insertar")
	}
	sq.action = "INSERT"
	return nil
}

/**
 * Valida los datos para actualizar y crea el query para actualizar
 * schema {[]models.Base}: modelo de la tabla
 * returns {error}: retorna errores ocurridos en la validación
 */
func (sq *SqlExec) Update(schema []models.Base) error {
	data := sq.Ob
	length := len(data)

	if length > 0 {
		var sqlExec = make([]map[string]interface{}, 0)
		var data_update []map[string]interface{}
		for _, item := range data {
			where := make(map[string]interface{})

			length_where := 0
			if item["where"] != nil {
				where = item["where"].(map[string]interface{})
				length_where = len(where)
				delete(item, "where")
			}
			preArray, err := _checkUpdate(schema, item)
			if err != nil {
				return err
			}
			preArray_where := make(map[string]interface{})
			if length_where > 0 {
				preArray, err := _checkWhere(schema, where)
				if err != nil {
					return err
				}
				preArray_where = preArray
			}

			data_update = append(data_update, preArray)
			var lineSqlExec = make(map[string]interface{}, 2)
			sqlPreparateUpdate := ""
			sqlWherePreparateUpdate := ""
			var i int
			var p uint64
			length_newMap := len(preArray)
			var valuesExec []interface{}
			char := "$"
			for k, v := range preArray {
				p++

				if i+1 < length_newMap {
					sqlPreparateUpdate += k + "= " + char + strconv.FormatUint(p, 10) + ", "
				} else {
					sqlPreparateUpdate += k + "= " + char + strconv.FormatUint(p, 10)
				}
				valuesExec = append(valuesExec, v)

				i++
			}
			if length_where > 0 {
				length_newMapWhere := len(preArray_where)
				i = 0

				for k, v := range preArray_where {
					p++
					if i+1 < length_newMapWhere {
						// sqlWherePreparateUpdate += fmt.Sprintf("%s = '%s' AND ", ke, va)
						sqlWherePreparateUpdate += k + " = " + char + strconv.FormatUint(p, 10) + " AND "
					} else {
						//sqlWherePreparateUpdate += fmt.Sprintf("%s = '%s'", ke, va)
						sqlWherePreparateUpdate += k + " = " + char + strconv.FormatUint(p, 10)
					}
					valuesExec = append(valuesExec, v)
					i++
				}
				if length_newMapWhere > 0 {
					sqlWherePreparateUpdate = "WHERE " + sqlWherePreparateUpdate
				}
			}
			// fmt.Println(sqlPreparateUpdate)
			sqlPreparate := fmt.Sprintf("UPDATE %s SET %s %s", sq.Table, sqlPreparateUpdate, sqlWherePreparateUpdate)
			lineSqlExec["sqlPreparate"] = sqlPreparate
			lineSqlExec["valuesExec"] = valuesExec
			sqlExec = append(sqlExec, lineSqlExec)

		}
		sq.Query = sqlExec
		sq.Data = data_update
	} else {
		return errors.New("No existen datos para actualizar")
	}
	sq.action = "UPDATE"
	return nil
}

/**
 * Valida los datos para Eliminar y crea el query para Eliminar
 * schema {[]models.Base}: modelo de la tabla
 * returns {error}: retorna errores ocurridos en la validación
 */
func (sq *SqlExec) Delete(schema []models.Base) error {
	data := sq.Ob
	length := len(data)

	if length > 0 {
		var sqlExec = make([]map[string]interface{}, 0)
		var data_delete []map[string]interface{}
		for _, item := range data {

			preArray, err := _checkWhere(schema, item)
			if err != nil {
				return err
			}

			data_delete = append(data_delete, preArray)
			var lineSqlExec = make(map[string]interface{}, 2)
			sqlWherePreparateDelete := ""
			var i int
			var p uint64
			length_newMap := len(preArray)
			var valuesExec []interface{}
			if length_newMap > 0 {
				sqlWherePreparateDelete += " WHERE "
			}
			char := "$"
			for k, v := range preArray {
				p++
				if i+1 < length_newMap {
					// sqlWherePreparateUpdate += fmt.Sprintf("%s = '%s' AND ", ke, va)
					sqlWherePreparateDelete += k + " = " + char + strconv.FormatUint(p, 10) + " AND "
				} else {
					//sqlWherePreparateUpdate += fmt.Sprintf("%s = '%s'", ke, va)
					sqlWherePreparateDelete += k + " = " + char + strconv.FormatUint(p, 10)
				}
				valuesExec = append(valuesExec, v)
				i++
			}

			sqlPreparate := fmt.Sprintf("DELETE FROM %s %s", sq.Table, sqlWherePreparateDelete)
			lineSqlExec["sqlPreparate"] = sqlPreparate
			lineSqlExec["valuesExec"] = valuesExec
			sqlExec = append(sqlExec, lineSqlExec)
		}
		sq.Query = sqlExec
		sq.Data = data_delete
	} else {
		return errors.New("No existen datos para actualizar")
	}
	sq.action = "DELETE"
	return nil
}

/**
 * Ejecuta el query
 * returns {error}: retorna errores ocurridos durante la ejecución
 */
func (sq *SqlExec) Exec(params ...bool) error {
	cnn := connection.Connection()
	ctx := context.Background()
	err_cnn := cnn.PingContext(ctx)
	if err_cnn != nil {
		return errors.New(fmt.Sprint("Error Sql PING: ", err_cnn))
	}
	cross := false
	if len(params) == 1 {
		cross = params[0]
	}
	dataExec := sq.Query
	defer cnn.Close()
	for _, item := range dataExec {
		sqlPre := item["sqlPreparate"].(string)
		if cross {
			if sq.action == "UPDATE" {
				sqlPre = Query_Cross_Update(sqlPre)
			}
		}
		// fmt.Println("PREPARED: ", sqlPre)
		stmt, err_prepare := cnn.Prepare(sqlPre)
		if err_prepare != nil {
			return errors.New(fmt.Sprint("Error Sql PREPARE: ", err_prepare))
		}
		valuesExec := item["valuesExec"].([]interface{})
		_, err_exec := stmt.Exec(valuesExec...)
		if err_exec != nil {
			return errors.New(fmt.Sprintf("Error Sql %s: %s", sq.action, err_exec.Error()))
		}
	}
	return nil
}

func _checkInsertSchema(schema []models.Base, tabla_map map[string]interface{}) (map[string]interface{}, error) {

	// var err_cont uint64 = 0
	var err_cont uint
	var error string

	data := make(map[string]interface{})

	for _, item := range schema {
		// fmt.Println(item.Name, tabla_map[item.Name])
		isNil := tabla_map[item.Name] == nil
		defaultIsNil := item.Default == nil
		if !isNil {
			value := tabla_map[item.Name]
			new_value, err := checkDataTypes(item.Type, value)
			if err != nil {
				err_cont++
				error += fmt.Sprintf("%d.- El campo %s %s", err_cont, item.Description, err.Error())
			}
			var val interface{}
			switch item.Type {
			case "string":
				val, err = caseString(new_value.(string), item.Strings)
			case "float64":
				val, err = caseFloat(new_value.(float64), item.Float)
			case "uint64":
				val, err = caseUint(new_value.(uint64), item.Uint)
			case "int64":
				val, err = caseInt(new_value.(int64), item.Int)
			default:
				val, err = nil, errors.New("tipo de dato no asignado")
			}
			if err == nil {
				data[item.Name] = val
			} else {
				err_cont++
				error += fmt.Sprintf("%d.- Se encontró fallas al validar el campo %s \n %s\n", err_cont, item.Description, err.Error())
			}
		} else {
			if !defaultIsNil {
				data[item.Name] = item.Default
			} else {
				if item.Required {
					err_cont++
					error += fmt.Sprintf("%d.- El campo %s es Requerido\n", err_cont, item.Description)
				}
			}
		}

	}
	if err_cont > 0 {
		return nil, errors.New(error)
	} else {
		return data, nil
	}

}
func _checkUpdate(schema []models.Base, tabla_map map[string]interface{}) (map[string]interface{}, error) {
	var err_cont uint
	var error string
	data := make(map[string]interface{})
	for _, item := range schema {
		isNil := tabla_map[item.Name] == nil
		if !isNil {
			if item.Update {
				value := tabla_map[item.Name]
				new_value, err := checkDataTypes(item.Type, value)
				if err != nil {
					err_cont++
					error += fmt.Sprintf("%d.- El campo %s %s", err_cont, item.Description, err.Error())
				}
				var val interface{}
				switch item.Type {
				case "string":
					if new_value.(string) == "" {
						if !item.Empty {
							err_cont++
							error += fmt.Sprintf("%d.- El campo %s no puede estar vació\n", err_cont, item.Description)
						}
					} else {
						val, err = caseString(new_value.(string), item.Strings)
					}
				case "float64":
					val, err = caseFloat(new_value.(float64), item.Float)
				case "uint64":
					val, err = caseUint(new_value.(uint64), item.Uint)
				case "int64":
					val, err = caseInt(new_value.(int64), item.Int)
				default:
					val, err = nil, errors.New("tipo de dato no asignado")
				}
				if err == nil {
					data[item.Name] = val
				} else {
					err_cont++
					error += fmt.Sprintf("%d.- Se encontró fallas al validar el campo %s \n %s\n", err_cont, item.Description, err.Error())
				}
			} else {
				err_cont++
				error += fmt.Sprintf("%d.- El campo %s no puede ser modificado\n", err_cont, item.Description)
			}
		}
	}
	if err_cont > 0 {
		return nil, errors.New(error)
	} else {
		return data, nil
	}
}

func _checkWhere(schema []models.Base, table_where map[string]interface{}) (map[string]interface{}, error) {
	var err_cont uint
	var error string
	data := make(map[string]interface{})
	for _, item := range schema {
		isNil := table_where[item.Name] == nil
		if !isNil {
			value := table_where[item.Name]
			if !item.Where && !item.Important {
				err_cont++
				error += fmt.Sprintf("%d.- El campo %s no puede ser utilizado de esta forma\n", err_cont, item.Description)
			} else {
				if value.(string) == "" {
					err_cont++
					error += fmt.Sprintf("%d.- El campo %s esta vació verificar\n", err_cont, item.Description)
				} else {
					data[item.Name] = value
				}

			}
		} else {
			if item.Important {
				err_cont++
				error += fmt.Sprintf("%d.- El campo %s es obligatorio\n", err_cont, item.Description)
			}
		}
	}
	if err_cont > 0 {
		return nil, errors.New(error)
	} else {
		return data, nil
	}
}

func caseString(value string, schema models.Strings) (string, error) {

	err_ := ""
	value = strings.TrimSpace(value)
	if !schema.Expr.MatchString(value) {
		err_ += "- No Cumple con las características\n"
		return "", errors.New(err_)
	} else {
		if schema.Date {
			err := date.CheckDate(value)
			if err != nil {
				err_ += fmt.Sprintf("- %s\n", err.Error())
				return "", errors.New(err_)
			} else {
				return value, nil
			}
		} else {
			if schema.Encriptar {
				result, _ := bcrypt.GenerateFromPassword([]byte(value), 13)
				value = string(result)
				return value, nil
			} else {
				if schema.Cifrar {
					hash, _ := cryptoAes.AesEncrypt_PHP([]byte(value), auth.GetKey_PrivateCrypto())
					value = hash
					return value, nil
				} else {
					if schema.Min > 0 {
						if len(value) < schema.Min {
							err_ += fmt.Sprintf("- No Cumple los caracteres mínimos que debe tener (%v)\n", schema.Min)
							return "", errors.New(err_)
						}
					}
					if schema.Max > 0 {
						if len(value) > schema.Max {
							err_ += fmt.Sprintf("- No Cumple los caracteres máximos que debe tener (%v)\n", schema.Max)
							return "", errors.New(err_)
						}
					}
					if err_ == "" {
						if schema.UpperCase {
							value = strings.ToUpper(value)
						} else if schema.LowerCase {
							value = strings.ToLower(value)
						}
						return value, nil
					} else {
						return value, errors.New(err_)
					}

				}
			}
		}
	}
}

func caseFloat(value float64, schema models.Floats) (float64, error) {
	error := ""
	err_cont := 0
	if schema.Menor != 0 {
		if value < schema.Menor {
			err_cont++
			error += fmt.Sprintf("- No puede se menor a %f", schema.Menor)
		}
	}
	if schema.Mayor != 0 {
		if value > schema.Mayor {
			err_cont++
			error += fmt.Sprintf("- No puede se mayor a %f", schema.Mayor)
		}
	}
	if !schema.Negativo {
		if value < float64(0) {
			err_cont++
			error += fmt.Sprintf("- No puede ser negativo")
		}
	}
	if schema.Porcentaje {
		value = value / float64(100)
	}
	if err_cont > 0 {
		return 0, errors.New(error)
	} else {
		return value, nil
	}

}
func caseInt(value int64, schema models.Ints) (int64, error) {
	error := ""
	err_cont := 0
	if !schema.Negativo {
		if value < int64(0) {
			err_cont++
			error += fmt.Sprintf("- No puede ser negativo")
		}
	}
	if schema.Min != 0 {
		if value < schema.Min {
			err_cont++
			error += fmt.Sprintf("- No puede se menor a %d", schema.Min)
		}
	}
	if schema.Max != 0 {
		if value > schema.Max {
			err_cont++
			error += fmt.Sprintf("- No puede se mayor a %d", schema.Max)
		}
	}
	if err_cont > 0 {
		return int64(0), errors.New(error)
	} else {
		return value, nil
	}

}
func caseUint(value uint64, schema models.Uints) (uint64, error) {
	if schema.Max > 0 {
		if value > schema.Max {
			error := fmt.Sprintf("- No esta en el rango permitido")
			return 0, errors.New(error)
		}
	}
	return value, nil
}

func convertStringToType(types string, value_undefined interface{}) (val interface{}, err error) {
	value := fmt.Sprintf("%v", value_undefined)
	switch types {
	case "uint64":
		val, err = strconv.ParseUint(value, 10, 64)
		return
	case "int64":
		val, err = strconv.ParseInt(value, 10, 64)
		return
	case "float64":
		val, err = strconv.ParseFloat(value, 64)
		return
	default:
		return nil, errors.New("No se puede convertir el tipo de dato")
	}
}

func checkDataTypes(types string, values interface{}) (interface{}, error) {
	type_value := reflect.TypeOf(values).String()
	switch types {
	case "string":
		if types == type_value {
			return values, nil
		}
		return nil, errors.New("tipo de dato incorrecto")
	case "float64":
		if types == type_value {
			return values, nil
		}

		if type_value == "string" {
			new_value, err := convertStringToType("float64", values)
			if err != nil {
				return nil, err
			}
			return new_value, nil
		} else if type_value == "int64" {
			new_value := float64(values.(int64))
			return new_value, nil
		}

		return nil, errors.New("tipo de dato incorrecto")
	case "uint64":
		if types == type_value {
			return values, nil
		}

		if type_value == "float64" {
			new_value := uint64(values.(float64))
			return new_value, nil
		} else if type_value == "string" {
			new_value, err := convertStringToType("uint64", values)
			if err != nil {
				return nil, err
			}
			return new_value, nil
		} else if type_value == "int64" {
			new_value := uint64(values.(int64))
			return new_value, nil
		}

		return nil, errors.New("tipo de dato incorrecto")
	case "int64":
		if types == type_value {
			return values, nil
		}

		if type_value == "float64" {
			new_value := int64(values.(float64))
			return new_value, nil
		} else if type_value == "string" {
			new_value, err := convertStringToType("int64", values)
			if err != nil {
				return nil, err
			}
			return new_value, nil
		} else if type_value == "uint64" {
			new_value := int64(values.(uint64))
			return new_value, nil
		}

		return nil, errors.New("tipo de dato incorrecto")
	default:
		return nil, errors.New("No se puede convertir el tipo de dato")
	}
}
