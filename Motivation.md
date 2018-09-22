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

### Example: Using a specifically typed function to implement a generic interface

Sometimes, you want to define an interface or function type without specific type information, but you don't want to require implementers of that interface to be generic. You define a type such as:

```go
type EndpointHandler func(event interface{}) (interface{}, error)
```

This represents some handler function that takes a context and event data, and returns some value and an error. An EndpointHandler function would be suitable for using as a handler for an API endpoint; the function takes endpoint input, and returns some value and an optional error.

The problem is, when writing an endpoint handler, it's easier to use concrete types:

```go
type handlerEvent struct {
	ID string
	Name string
}

type handlerOutput struct {
	Email string
}

func MyHandler(event handlerEvent) (handlerOutput, error) {
	// event.ID and event.Name can be accessed directly
	// the output is a clearly defined type
}
```

...but now, `MyHandler` cannot be used where an `EndpointHandler` is expected, because the function signature is different. To remedy this, the `go-mirror` package (TODO, not implemented yet) exports the function `FunctionWrap`, which takes two functions as input. The first should be a function with the desired signature, and the second is the function to wrap.

```go
targetSignature := func(event interface{}) (interface{}, error) { return nil, nil }
wrappedFunc := mirror.FunctionWrap(targetSignature, MyHandler)
// wrappedFunc can be used as an EndpointHandler
```

This allows defining functions with concrete inputs and outputs, while still being able to implement a generic interface.