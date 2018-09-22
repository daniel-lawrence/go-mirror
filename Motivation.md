## Motivation

There are several situations in which someone might want to use more generic code in Go. While it's possible to use the `reflect` package in the standard library for a wide variety of tasks, it's fairly low-level, and not very user-friendly. This package is meant to provide a higher-level interface to reflection in Go, so that some common tasks can be done without requiring deep knowledge of Go's type system.

### Example: Accessing struct fields by name at runtime

Let's say you want to write a function that allows passing in any struct (technically, passing in an `interface{}` that's intended to be a struct). As an example, we want a function that saves a value to a data store. We just want to save whatever is passed in, and we'll assume that it will require an `ID` field as a primary key. Here's how it might be done with the `reflect` package:

```go
func SaveObject(obj interface{}) {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Struct {
		return
	}

	idField := objValue.FieldByName("ID")
	if !idField.CanInterface() {
		return
	}

	id, ok := idField.Interface().(string)
	if !ok {
		return
	}

	// do something with id
}
```

This is very clumsy, and a complete example would be worse.

The `go-mirror` package exports a `StructToMap` function which can make this a lot easier:

```go
func SaveObject(obj interface{}) {
    objFields := mirror.StructToMap(obj)
    id := objFields["ID"].(string)
    // do something with id
}
```

That's better! In the future, other methods can be added as well, such as `StructToStringMap`, which could return a map of all struct fields that are strings. This would completely remove the need for type assertions in many cases.
