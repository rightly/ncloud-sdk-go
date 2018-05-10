package ncloud

import (
	"strings"
	"encoding/xml"
	"encoding/json"
	"reflect"
	"errors"
)

// TODO: Documenting

func Unmarshal2(data []byte, v interface{}) (err error) {

	if strings.HasPrefix(string(data), "<") {
		err = xml.Unmarshal(data, v)
	} else {
		err = json.Unmarshal(data, v)
	}
	return
}

func Unmarshal(r *Request) error {
	cType := r.HTTPResponse.Header.Get("Content-Type")

	if strings.Contains(cType, "xml") {
		err := xml.Unmarshal(r.Body, r.Data)
		return err
	}
	if strings.Contains(cType, "json") {
		err := json.Unmarshal(r.Body, r.Data)
		return err
	}

	err := errors.New("unmarshal fail, Content-Type Header is not xml or json")

	return err
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
			return string(b), nil
		}
		e = err
	case "xml":
		b, err := xml.MarshalIndent(v, "", "\t")
		if err == nil {
			return string(b), nil
		}
		e = err
	default:
		b, err := json.MarshalIndent(v, "", "\t")
		if err == nil {
			return string(b), nil
		}
		e = err
	}

	return reflect.TypeOf(v).String() + " Marshal fail", e
}