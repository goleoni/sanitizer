package sanitizer

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

const valueSeparator = "="
const wordSeparator = " "

// Sanitizer is the base struct for the sanitize calls
type Sanitizer struct{}

// NewSanitizer returns a new sanitizer struct
func NewSanitizer() *Sanitizer {
	return &Sanitizer{}
}

// Sanitize is the function that gets an struct and sanitize accourding with the tags of struct
//   For example:
//     type MyCustom struct {
//       Name       string `transform:"firstword"`
// 	     Zipcode    string `transform:"length=8"`
// 	     Occupation string `transform:"initial=developer"`
//     }
func (Sanitizer) Sanitize(s interface{}) {
	var reflectValue reflect.Value

	reflectValue = reflect.ValueOf(s)
	reflectElem := reflectValue.Elem()
	curKind := reflectElem.Kind()

	if curKind == reflect.Struct {
		typ := reflectElem.Type()

		numberOfFields := typ.NumField()

		for i := 0; i < numberOfFields; i++ {
			fld := typ.Field(i)
			fieldVal := reflectElem.FieldByName(fld.Name)

			if fieldVal.IsValid() {
				if fieldVal.CanSet() {
					tagValidate := fld.Tag.Get("transform")
					splittedValidateString := strings.Split(tagValidate, valueSeparator)

					var value string
					key := splittedValidateString[0]

					if len(splittedValidateString) > 2 {
						log.Fatal("3 sanitizer args not supported yet")
					} else if len(splittedValidateString) > 1 {
						value = splittedValidateString[1]
					}

					if fieldVal.Kind() == reflect.String {
						stringVal := fieldVal.String()

						if key == "firstword" {
							separatedWord := strings.Split(stringVal, wordSeparator)
							fieldVal.SetString(separatedWord[0])
						}

						if key == "length" {
							i, err := strconv.Atoi(value)
							if err != nil {
								log.Fatal("length should be a string:", err)
							}

							fieldVal.SetString(firstCharactersString(stringVal, i))
						}

						if key == "initial" {
							if stringVal == "" {
								fieldVal.SetString(value)
							}
						}
					}
				}
			}
		}
	} else {
		log.Fatal("not implemented sanitizer for", curKind)
	}
}

func firstCharactersString(s string, n int) string {
	i := 0

	for j := range s {
		if i == n {
			return s[:j]
		}

		i++
	}
	return s
}
