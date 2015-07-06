package pingpp

import (
	"bytes"
	"encoding/json"
)

func JsonEncode(v interface{}) ([]byte, error) {
	return json.Marshal(&v)
}

func JsonDecode(p []byte, v interface{}) error {
	obj := json.NewDecoder(bytes.NewBuffer(p))
	obj.UseNumber()
	return obj.Decode(&v)
}
