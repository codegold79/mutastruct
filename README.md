# mutastruct

Use the reflect package to make atypical changes to a struct.

To use, import the mutastruct package.

Call the following function to set a struct field.

```go
func Set(anyStruct any, fieldName string, setting any) error {
    ...
}
```

The first argument must be a pointer to a struct.
The last parameter, `setting`, is the desired value of the field specified by `fieldName` (the second parameter).
The `setting` argument must be the same type as the field you wish to change.
