package models

import (
	"bytes"
	"encoding/json"
)

// openapi generator bug: https://github.com/OpenAPITools/openapi-generator/issues/11374
func newStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()

	return dec
}
