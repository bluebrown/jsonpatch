# Go JSON Patch

This package implements types to reflect the json patch spec [rfc6902](https://datatracker.ietf.org/doc/html/rfc6902/). Additionally the patch type has helper methods to add items to the patch slice for the different operations.

## Synopsis

The patch type is a slice of patch items. Calling the methods on the patch type will add items to the patch slice.

```go
package main

import (
 "encoding/json"
 "fmt"

 "github.com/bluebrown/jsonpatch"
)

func main() {
  patch := jsonpatch.New()
  patch.Test("/a/b/c", "foo")
  patch.Remove("/a/b/c")
  patch.Add("/a/b/c", []string{"foo", "bar"})
  patch.Replace("/a/b/c", 42)
  patch.Move("/a/b/c", "/a/b/d")
  patch.Copy("/a/b/d", "/a/b/e")
  b, _ := json.MarshalIndent(patch, "", " ")
  fmt.Println(string(b))
}
```

<details>
<summary>Result</summary>

```json
[
  {
    "op": "test",
    "path": "/a/b/c",
    "value": "foo"
  },
  {
    "op": "remove",
    "path": "/a/b/c"
  },
  {
    "op": "add",
    "path": "/a/b/c",
    "value": [
      "foo",
      "bar"
    ]
  },
  {
    "op": "replace",
    "path": "/a/b/c",
    "value": 42
  },
  {
    "op": "move",
    "from": "/a/b/c",
    "path": "/a/b/d"
  },
  {
    "op": "copy",
    "from": "/a/b/d",
    "path": "/a/b/e"
  }
]
```

</details>

## Chainable methods

All methods are chainable.

```go
jsonpatch.New().Add("/a/b/c", []string{"foo", "bar"}).Delete("/b/c/e")
```

## PatchBuilder Interface

The patch type implements the PatcherBuilder interface.

```go
type PatchBuilder interface {
  Test(path string, value interface{}) PatchBuilder    // Add a test operation
  Remove(path string) PatchBuilder                     // Add a remove operation
  Add(path string, value interface{}) PatchBuilder     // Add an add operation
  Replace(path string, value interface{}) PatchBuilder // Add a replace operation
  Move(from, to string) PatchBuilder                   // Add a move operation
  Copy(from, to string) PatchBuilder                   // Add a copy operation
}
```
