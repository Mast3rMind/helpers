package helpers

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

func StructToBSONMap(st interface{}) (m map[string]interface{}) {

	s := reflect.ValueOf(st).Elem()
	typeOfT := s.Type()

	m = make(map[string]interface{})

	for i := 0; i < s.NumField(); i++ {

		field := s.Field(i)
		typeField := typeOfT.Field(i)

		fieldName := typeField.Tag.Get("bson")

		if fieldName == "" {

			fieldName = typeField.Name
		}

		m[fieldName] = field.Interface()
	}

	return
}

func IsNil(v interface{}) bool {
	return reflect.ValueOf(v).IsNil()
}

func DecodeJSON(r io.Reader, t interface{}) (err error) {

	err = json.NewDecoder(r).Decode(t)
	return
}

func SHA1(data []byte) string {

	hash := sha1.New()
	hash.Write(data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// From http://devpy.wordpress.com/2013/10/24/create-random-string-in-golang/
func RandomString(n int) string {

	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func RandomInt(a, b int) int {

	var bytes = make([]byte, 1)
	rand.Read(bytes)

	per := float32(bytes[0]) / 256.0
	dif := Max(a, b) - Min(a, b)

	return Min(a, b) + int(per*float32(dif))

}

func Max(a, b int) int {

	if a >= b {

		return a
	}

	return b
}

func Min(a, b int) int {

	if a <= b {

		return a
	}

	return b
}
