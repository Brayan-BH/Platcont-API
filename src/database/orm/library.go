package orm

import (
	"fmt"
	"regexp"
	"strings"
)

func Query_Cross_Update(query string) string {

	queryregEx := regexp.MustCompile(`ADD(.*?)_SUMA=`)
	array := queryregEx.FindAllStringSubmatch(query, -1)
	for _, v := range array {

		new := fmt.Sprintf("%s=%s+", v[1], v[1])

		query = strings.Replace(query, v[0], new, -1)
	}

	return query
}
