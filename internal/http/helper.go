package http

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func GenerateQueryString(object interface{}) string {
	_object, err := json.Marshal(object)
	if err != nil {
		log.Panic(err)
	}

	maps := map[string]interface{}{}
	err = json.Unmarshal(_object, &maps)
	if err != nil {
		log.Panic(err)
	}

	queries := []string{}
	for k, v := range maps {
		if v != nil {
			switch v.(type) {
			case string:
				if len(v.(string)) > 0 {
					queries = append(queries, fmt.Sprintf("%s=%+v", k, v))
				}
			default:
				queries = append(queries, fmt.Sprintf("%s=%+v", k, v))
			}
		}
	}
	return strings.Join(queries, "&")
}
