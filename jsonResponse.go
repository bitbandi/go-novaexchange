package novaexchange

import (
	"encoding/json"
	"reflect"
	"strings"
)

type jsonResponse struct {
	Status  string                     `json:"status"`
	Message string                     `json:"message"`
	Result  map[string]json.RawMessage `json:"-"`
}

func (t *jsonResponse) UnmarshalJSON(b []byte) error {
	type Alias jsonResponse
	t2 := Alias{}
	err := json.Unmarshal(b, &t2)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &(t2.Result))
	if err != nil {
		return err
	}

	typ := reflect.TypeOf(t2)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]
		if jsonTag != "" && jsonTag != "-" {
			delete(t2.Result, jsonTag)
		}
	}

	*t = jsonResponse(t2)

	return nil
}
