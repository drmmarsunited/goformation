package serverless

import (
	"bytes"
	"encoding/json"
	"io"
	"sort"

	"github.com/awslabs/goformation/v6/cloudformation/utils"
)

// StateMachine_Properties is a helper struct that can hold either a CloudWatchEventEvent, EventBridgeRuleEvent, ScheduleEvent, or ApiEvent value
type StateMachine_Properties struct {
	CloudWatchEventEvent *StateMachine_CloudWatchEventEvent
	EventBridgeRuleEvent *StateMachine_EventBridgeRuleEvent
	ScheduleEvent        *StateMachine_ScheduleEvent
	ApiEvent             *StateMachine_ApiEvent
}

func (r StateMachine_Properties) value() interface{} {
	ret := []interface{}{}

	if r.CloudWatchEventEvent != nil {
		ret = append(ret, *r.CloudWatchEventEvent)
	}

	if r.EventBridgeRuleEvent != nil {
		ret = append(ret, *r.EventBridgeRuleEvent)
	}

	if r.ScheduleEvent != nil {
		ret = append(ret, *r.ScheduleEvent)
	}

	if r.ApiEvent != nil {
		ret = append(ret, *r.ApiEvent)
	}

	sort.Sort(utils.ByJSONLength(ret)) // Heuristic to select best attribute
	if len(ret) > 0 {
		return ret[0]
	}

	return nil
}

func (r StateMachine_Properties) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *StateMachine_Properties) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case map[string]interface{}:
		val = val // This ensures val is used to stop an error

		reader := bytes.NewReader(b)
		decoder := json.NewDecoder(reader)
		decoder.DisallowUnknownFields()
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.CloudWatchEventEvent); err != nil {
			r.CloudWatchEventEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.EventBridgeRuleEvent); err != nil {
			r.EventBridgeRuleEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.ScheduleEvent); err != nil {
			r.ScheduleEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.ApiEvent); err != nil {
			r.ApiEvent = nil
		}
		reader.Seek(0, io.SeekStart)

	case []interface{}:

	}

	return nil
}
