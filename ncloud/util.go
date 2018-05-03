package ncloud

import (
	"strings"
	"encoding/xml"
	"encoding/json"
	"reflect"
)

// TODO: Documenting

func Unmarshal(data []byte, v interface{}) (err error) {

	if strings.HasPrefix(string(data), "<") {
		err = xml.Unmarshal(data, v)
	} else {
		err = json.Unmarshal(data, v)
	}

	return nil
}

func Marshal(v interface{}) ([]byte, error)  {
	return nil, nil
}

func String(v interface{}, s string) (string, error) {
	var e error

	switch s {
	case "json":
		b, err := json.MarshalIndent(v, "", "\t")
		if err == nil {
			return string(b[:]), nil
		}
		e = err
	case "xml":
		b, err := xml.MarshalIndent(v, "", "\t")
		if err == nil {
			return string(b[:]), nil
		}
		e = err
	default:
		b, err := json.MarshalIndent(v, "", "\t")
		if err == nil {
			return string(b[:]), nil
		}
		e = err
	}

	return reflect.TypeOf(v).String() + " Marshal fail", e
}