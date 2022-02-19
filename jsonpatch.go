package jsonpatch

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
	Test(path string, value interface{}) Patcher    // Add a test operation
	Remove(path string) Patcher                     // Add a remove operation
	Add(path string, value interface{}) Patcher     // Add an add operation
	Replace(path string, value interface{}) Patcher // Add a replace operation
	Move(from, to string) Patcher                   // Add a move operation
	Copy(from, to string) Patcher                   // Add a copy operation
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

func (p *Patch) Test(path string, value interface{}) *Patch {
	*p = append(*p, Item{
		Op:    OperationTest,
		Path:  path,
		Value: value,
	})
	return p
}

func (p *Patch) Remove(path string) *Patch {
	*p = append(*p, Item{
		Op:   OperationRemove,
		Path: path,
	})
	return p
}

func (p *Patch) Add(path string, value interface{}) *Patch {
	*p = append(*p, Item{
		Op:    OperationAdd,
		Path:  path,
		Value: value,
	})
	return p
}

func (p *Patch) Replace(path string, value interface{}) *Patch {
	*p = append(*p, Item{
		Op:    OperationReplace,
		Path:  path,
		Value: value,
	})
	return p
}

func (p *Patch) Move(from, to string) *Patch {
	*p = append(*p, Item{
		Op:   OperationMove,
		From: from,
		Path: to,
	})
	return p
}

func (p *Patch) Copy(from, to string) *Patch {
	*p = append(*p, Item{
		Op:   OperationCopy,
		From: from,
		Path: to,
	})
	return p
}
