package utils

import (
	"encoding/json"
	"fmt"
)

func Clone(obj interface{}) interface{} {
	marshal, err := json.Marshal(obj)

	if err != nil {
		return fmt.Errorf("could not clone, serialization failed: %w", err)
	}

	var out interface{}
	err = json.Unmarshal(marshal, &out)

	if err != nil {
		return fmt.Errorf("could not clone, deserialization failed: %w", err)
	}

	return out
}
