package serverless

import (
	"bytes"
	"encoding/json"
	"io"
	"sort"

	"github.com/awslabs/goformation/v6/cloudformation/utils"
)

// HttpApi_CorsConfiguration is a helper struct that can hold either a Boolean or CorsConfigurationObject value
type HttpApi_CorsConfiguration struct {
	Boolean *bool

	CorsConfigurationObject *HttpApi_CorsConfigurationObject
}

func (r HttpApi_CorsConfiguration) value() interface{} {
	ret := []interface{}{}

	if r.Boolean != nil {
		ret = append(ret, r.Boolean)
	}

	if r.CorsConfigurationObject != nil {
		ret = append(ret, *r.CorsConfigurationObject)
	}

	sort.Sort(utils.ByJSONLength(ret)) // Heuristic to select best attribute
	if len(ret) > 0 {
		return ret[0]
	}

	return nil
}

func (r HttpApi_CorsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *HttpApi_CorsConfiguration) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case bool:
		r.Boolean = &val

	case map[string]interface{}:
		val = val // This ensures val is used to stop an error

		reader := bytes.NewReader(b)
		decoder := json.NewDecoder(reader)
		decoder.DisallowUnknownFields()
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.CorsConfigurationObject); err != nil {
			r.CorsConfigurationObject = nil
		}
		reader.Seek(0, io.SeekStart)

	case []interface{}:

	}

	return nil
}
