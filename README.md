# Go JSON Patch

This package implements types to reflect the json patch spec [rfc6902](https://datatracker.ietf.org/doc/html/rfc6902/). Additionally the patch type has helper methods to add items to the patch slice for the different operations.

## Synopsis

The patch type is a slice of patch items. Calling the methods on the patch type will add items to the patch slice.

```go
package main

import (
 "fmt"

 "github.com/bluebrown/jsonpatch"
)

func main() {
    patch := jsonpatch.Patch{}
    patch.Test("/a/b/c", "foo")
    patch.Remove("/a/b/c")
    patch.Add("/a/b/c", []string{"foo", "bar"})
    patch.Replace("/a/b/c", 42)
    patch.Move("/a/b/c", "/a/b/d")
    patch.Copy("/a/b/d", "/a/b/e")
    fmt.Println(patch.Encode().String())
    // [
    //  { "op": "test", "path": "/a/b/c", "value": "foo" },
    //  { "op": "remove", "path": "/a/b/c" },
    //  { "op": "add", "path": "/a/b/c", "value": [ "foo", "bar" ] },
    //  { "op": "replace", "path": "/a/b/c", "value": 42 },
    //  { "op": "move", "from": "/a/b/c", "path": "/a/b/d" },
    //  { "op": "copy", "from": "/a/b/d", "path": "/a/b/e" }
    // ]
}
```

*Note, the output in this example is formatted for readability. Normally, the output is not formatted.*
