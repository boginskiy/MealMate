package pkg

import "encoding/json"

type ExEncoder interface {
	Deserialization([]byte, any) error
}

type ExtraEncode struct {
}

func NewExtraEncode() *ExtraEncode {
	return &ExtraEncode{}
}

func (e *ExtraEncode) Deserialization(sliceByte []byte, pStruct any) error {
	return json.Unmarshal(sliceByte, pStruct)
}
