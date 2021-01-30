# About Sanitizer

Golang sanitizer is a study case for use of reflections and tags.

# Usage

For example, with the following struct:

```
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
```

# Community

This is a WIP project and will be in continuous improvement.

If anyone should like to collaborate, I appreciate it.