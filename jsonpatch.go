package jsonpatch

import (
	"bytes"
	"encoding/json"
)

/*
JSON Patch
Spec: https://datatracker.ietf.org/doc/html/rfc6902/
Example:
[
	{ "op": "test", "path": "/a/b/c", "value": "foo" },
	{ "op": "remove", "path": "/a/b/c" },
	{ "op": "add", "path": "/a/b/c", "value": [ "foo", "bar" ] },
	{ "op": "replace", "path": "/a/b/c", "value": 42 },
	{ "op": "move", "from": "/a/b/c", "path": "/a/b/d" },
	{ "op": "copy", "from": "/a/b/d", "path": "/a/b/e" }
]
*/

type Patcher interface {
	Test(path string, value interface{})
	Remove(path string)
	Add(path string, value interface{})
	Replace(path string, value interface{})
	Move(from, to string)
	Copy(from, to string)
	Encode() *bytes.Buffer
}

type Operation string

const (
	OperationTest    Operation = "test"
	OperationRemove  Operation = "remove"
	OperationAdd     Operation = "add"
	OperationReplace Operation = "replace"
	OperationMove    Operation = "move"
	OperationCopy    Operation = "copy"
)

type Item struct {
	Op    Operation   `json:"op"`
	From  string      `json:"from,omitempty"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}

type Patch []Item

func New() *Patch {
	return &Patch{}
}

func (p *Patch) Test(path string, value interface{}) {
	*p = append(*p, Item{
		Op:    OperationTest,
		Path:  path,
		Value: value,
	})
}

func (p *Patch) Remove(path string) {
	*p = append(*p, Item{
		Op:   OperationRemove,
		Path: path,
	})
}

func (p *Patch) Add(path string, value interface{}) {
	*p = append(*p, Item{
		Op:    OperationAdd,
		Path:  path,
		Value: value,
	})
}

func (p *Patch) Replace(path string, value interface{}) {
	*p = append(*p, Item{
		Op:    OperationReplace,
		Path:  path,
		Value: value,
	})
}

func (p *Patch) Move(from, to string) {
	*p = append(*p, Item{
		Op:   OperationMove,
		From: from,
		Path: to,
	})
}

func (p *Patch) Copy(from, to string) {
	*p = append(*p, Item{
		Op:   OperationCopy,
		From: from,
		Path: to,
	})
}

func (p *Patch) Encode() *bytes.Buffer {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)
	return b
}
