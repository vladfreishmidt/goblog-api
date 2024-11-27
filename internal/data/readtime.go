package data

import (
	"fmt"
	"strconv"
)

type ReadTime int32

func (r ReadTime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
