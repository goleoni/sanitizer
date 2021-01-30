package sanitizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// MyCustom estrutura criada para ler as "Tags" como um exemplo de uso
type MyCustom struct {
	Name       string `transform:"firstword"`
	Zipcode    string `transform:"length=8"`
	Occupation string `transform:"initial=developer"`
}

func TestTransform(t *testing.T) {
	sanitizer := NewSanitizer()

	myStruct := MyCustom{Name: "Andre Leoni", Zipcode: "122433200000"}

	sanitizer.Sanitize(&myStruct)

	assert.Equal(t, "Andre", myStruct.Name)
	assert.Equal(t, "12243320", myStruct.Zipcode)
	assert.Equal(t, "developer", myStruct.Occupation)
}
