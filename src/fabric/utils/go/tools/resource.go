package tools

import "encoding/json"

// Define the resource structure
type Resource struct {
	hashToString       string `json:"hash"`
}

func (r Resource) ToBytes() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return b
}

func NewResource(b []byte) (Resource, error) {
	r := Resource{}
	err := json.Unmarshal(b, &r)
	return r, err
}
