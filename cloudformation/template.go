package cloudformation

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/drmmarsunited/goformation/v7/intrinsics"
	"gopkg.in/yaml.v3"
)

// Template represents an AWS CloudFormation template
// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html
type Template struct {
	AWSTemplateFormatVersion string                 `json:"AWSTemplateFormatVersion,omitempty"`
	Transform                *Transform             `json:"Transform,omitempty"`
	Description              string                 `json:"Description,omitempty"`
	Metadata                 map[string]interface{} `json:"Metadata,omitempty"`
	Parameters               Parameters             `json:"Parameters,omitempty"`
	Mappings                 map[string]interface{} `json:"Mappings,omitempty"`
	Conditions               map[string]interface{} `json:"Conditions,omitempty"`
	Resources                Resources              `json:"Resources,omitempty"`
	Outputs                  Outputs                `json:"Outputs,omitempty"`
	Globals                  Globals                `json:"Globals,omitempty"`
}

type Parameter struct {
	Type                  string        `json:"Type"`
	Description           *string       `json:"Description,omitempty"`
	Default               interface{}   `json:"Default,omitempty"`
	AllowedPattern        *string       `json:"AllowedPattern,omitempty"`
	AllowedValues         []interface{} `json:"AllowedValues,omitempty"`
	ConstraintDescription *string       `json:"ConstraintDescription,omitempty"`
	MaxLength             *int          `json:"MaxLength,omitempty"`
	MinLength             *int          `json:"MinLength,omitempty"`
	MaxValue              *float64      `json:"MaxValue,omitempty"`
	MinValue              *float64      `json:"MinValue,omitempty"`
	NoEcho                *bool         `json:"NoEcho,omitempty"`
}

type Output struct {
	Value       interface{} `json:"Value"`
	Description *string     `json:"Description,omitempty"`
	Export      *Export     `json:"Export,omitempty"`
	Condition   *string     `json:"Condition,omitempty"`
}

type Export struct {
	Name string `json:"Name,omitempty"`
}

type Resource interface {
	AWSCloudFormationType() string
}

type Parameters map[string]Parameter
type Resources map[string]Resource
type Globals map[string]Resource
type Outputs map[string]Output

func (resources *Resources) UnmarshalJSON(b []byte) error {
	// Resources
	var rawResources map[string]*json.RawMessage
	err := json.Unmarshal(b, &rawResources)

	if err != nil {
		return err
	}

	newResources := Resources{}
	for name, raw := range rawResources {
		res, err := unmarshallResource(name, raw)
		if err != nil {
			return err
		}
		newResources[name] = res
	}

	*resources = newResources
	return nil
}

func (globals *Globals) UnmarshalJSON(b []byte) error {
	// Globals

	var rawGlobals map[string]*json.RawMessage
	err := json.Unmarshal(b, &rawGlobals)

	if err != nil {
		return err
	}

	newGlobals := Globals{}
	for name, raw := range rawGlobals {
		res, err := unmarshallGlobal(name, raw)
		if err != nil {
			return err
		}
		newGlobals[name] = res
	}

	*globals = newGlobals
	return nil
}

func unmarshallResource(name string, raw_json *json.RawMessage) (Resource, error) {
	var err error

	type rType struct {
		Type string
	}

	var rtype rType
	if err = json.Unmarshal(*raw_json, &rtype); err != nil {
		return nil, err
	}

	if rtype.Type == "" {
		return nil, fmt.Errorf("Cannot find Type for %v", name)
	}

	// Custom Resource Handler
	var resourceStruct Resource

	if strings.HasPrefix(rtype.Type, "Custom::") {
		resourceStruct = &CustomResource{Type: rtype.Type}
	} else {
		resourceStruct = AllResources()[rtype.Type]
	}

	err = json.Unmarshal(*raw_json, resourceStruct)

	if err != nil {
		return nil, err
	}
	return resourceStruct, nil
}

func unmarshallGlobal(name string, raw_json *json.RawMessage) (Resource, error) {
	var err error

	// Global handler
	resourceStruct := AllResources()[name]

	err = json.Unmarshal(*raw_json, resourceStruct)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return resourceStruct, nil
}

type Transform struct {
	String *string

	StringArray *[]string
}

func (t Transform) value() interface{} {
	if t.String != nil {
		return t.String
	}

	if t.StringArray != nil {
		return t.StringArray
	}

	return nil
}

func (t *Transform) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.value())
}

func (t *Transform) UnmarshalJSON(b []byte) error {
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case string:
		t.String = &val

	case []string:
		t.StringArray = &val

	case []interface{}:
		var strslice []string
		for _, i := range val {
			switch str := i.(type) {
			case string:
				strslice = append(strslice, str)
			}
		}
		t.StringArray = &strslice
	}

	return nil
}

// NewTemplate creates a new AWS CloudFormation template struct
func NewTemplate() *Template {
	return &Template{
		AWSTemplateFormatVersion: "2010-09-09",
		Description:              "",
		Metadata:                 map[string]interface{}{},
		Parameters:               Parameters{},
		Mappings:                 map[string]interface{}{},
		Conditions:               map[string]interface{}{},
		Resources:                Resources{},
		Outputs:                  Outputs{},
		Globals:                  Globals{},
	}
}

// JSON converts an AWS CloudFormation template object to JSON
func (t *Template) JSON() ([]byte, error) {

	j, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		return nil, err
	}

	return intrinsics.ProcessJSON(j, nil)

}

// YAML converts an AWS CloudFormation template object to YAML
func (t *Template) YAML() ([]byte, error) {
	j, err := t.JSON()
	if err != nil {
		return nil, err
	}

	var jsonObj interface{}
	// We are using yaml.Unmarshal here (instead of json.Unmarshal) because the
	// Go JSON library doesn't try to pick the right number type (int, float,
	// etc.) when unmarshalling to interface{}, it just picks float64
	// universally. go-yaml does go through the effort of picking the right
	// number type, so we can preserve number type throughout this process.
	err = yaml.Unmarshal(j, &jsonObj)
	if err != nil {
		return nil, err
	}

	// Marshal this object into YAML
	return yaml.Marshal(jsonObj)
}
