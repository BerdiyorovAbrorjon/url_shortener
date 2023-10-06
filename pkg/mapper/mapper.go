package mapper

import jsoniter "github.com/json-iterator/go"

func Map(input interface{}, output interface{}) error {
	rawBytes, err := jsoniter.Marshal(input)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(rawBytes, output)
}
