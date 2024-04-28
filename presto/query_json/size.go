package query_json

import (
	"encoding/json"
	"fmt"
	"github.com/docker/go-units"
)

type SISize float64

func (s *SISize) UnmarshalJSON(bytes []byte) error {
	var v any
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*s = SISize(value)
		return nil
	case string:
		sz, err := units.RAMInBytes(value)
		if err != nil {
			return err
		}
		*s = SISize(sz)
		return nil
	default:
		return fmt.Errorf("invalid SI size")
	}
}

func (s *SISize) MarshalJSON() ([]byte, error) {
	return json.Marshal(units.BytesSize(float64(*s)))
}
