# google-json-style-response

[![Build Status](https://travis-ci.org/ory-platform/google-json-style-response.svg)](https://travis-ci.org/ory-platform/google-json-style-response)

Use one command to style your JSON responses like Google does.
For more information, check out [Google's JSON Style Guide](https://google-styleguide.googlecode.com/svn/trunk/jsoncstyleguide.xml).

**WARNING:** This package is still in development, API might receive changes and features like pagination are still missing.

## Usage

```
go get github.com/ory-platform/google-json-style-response/response
```

### Send a success response

```go
import "github.com/ory-platform/google-json-style-response/response"

// ...
func(w http.ResponseWriter, r *http.Request) {
    data := struct{
        A string
    } {
        A: "b",
    }

    responder := New("1.0")
    o := responder.Success(data)
    err := responder.Write(w, o)
    if err != nil {
        // ...
    }
}
// ...
```

### Send an error response

```go
import "github.com/ory-platform/google-json-style-response/response"

// ...
func(w http.ResponseWriter, r *http.Request) {
    responder := New("1.0")
    o := responder.Error(500, "Internal Server Error")
    err := responder.Write(w, o)
}
// ...
```

### Send multiple errors

```go
import "github.com/ory-platform/google-json-style-response/response"

// ...
func(w http.ResponseWriter, r *http.Request) {
    responder := New("1.0")
    o := responder.Error(500, "Internal Server Error")
    responder.AddError(ErrorItem{
        Message: "baz",
    })
    err := responder.Write(w, o)
}
// ...
```
