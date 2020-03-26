package responses

import (
	"encoding/json"
	r "github.com/luqmansen/hanako/api/rest/utils"
	"github.com/micro/go-micro/v2/errors"
	"log"
	"net/http"
	"reflect"
)

func JSON(data []interface{}) (string, int32) {
	dict := make(map[string]interface{})
	if reflect.ValueOf(data[0]).Len() == 0 {
		return "No content", http.StatusNoContent
	}
	dict["data"] = data
	b, err := json.Marshal(dict)
	if err != nil {
		log.Fatal(errors.InternalServerError(r.ServiceName+".search", err.Error()))
	}
	return string(b), http.StatusOK
}
