package jsonpatch_test

import (
	"encoding/json"
	"testing"

	"github.com/bluebrown/jsonpatch"
)

func TestOps(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		name     string
		item     jsonpatch.Item
		expected string
	}{
		{
			name: "Test",
			item: jsonpatch.Item{
				Op:    jsonpatch.OperationTest,
				Path:  "/a/b/c",
				Value: "foo",
			},
			expected: `[{"op":"test","path":"/a/b/c","value":"foo"}]`,
		},
		{
			name: "Remove",
			item: jsonpatch.Item{
				Op:   jsonpatch.OperationRemove,
				Path: "/a/b/c",
			},
			expected: `[{"op":"remove","path":"/a/b/c"}]`,
		},
		{
			name: "Add",
			item: jsonpatch.Item{
				Op:    jsonpatch.OperationAdd,
				Path:  "/a/b/c",
				Value: []string{"foo", "bar"},
			},
			expected: `[{"op":"add","path":"/a/b/c","value":["foo","bar"]}]`,
		},
		{
			name: "Replace",
			item: jsonpatch.Item{
				Op:    jsonpatch.OperationReplace,
				Path:  "/a/b/c",
				Value: 42,
			},
			expected: `[{"op":"replace","path":"/a/b/c","value":42}]`,
		},
		{
			name: "Move",
			item: jsonpatch.Item{
				Op:   jsonpatch.OperationMove,
				From: "/a/b/c",
				Path: "/a/b/d",
			},
			expected: `[{"op":"move","from":"/a/b/c","path":"/a/b/d"}]`,
		},
		{
			name: "Copy",
			item: jsonpatch.Item{
				Op:   jsonpatch.OperationCopy,
				From: "/a/b/d",
				Path: "/a/b/e",
			},
			expected: `[{"op":"copy","from":"/a/b/d","path":"/a/b/e"}]`,
		},
	}
	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			patch := jsonpatch.New()

			switch tc.item.Op {
			case jsonpatch.OperationTest:
				patch.Test(tc.item.Path, tc.item.Value)
			case jsonpatch.OperationRemove:
				patch.Remove(tc.item.Path)
			case jsonpatch.OperationAdd:
				patch.Add(tc.item.Path, tc.item.Value)
			case jsonpatch.OperationReplace:
				patch.Replace(tc.item.Path, tc.item.Value)
			case jsonpatch.OperationMove:
				patch.Move(tc.item.From, tc.item.Path)
			case jsonpatch.OperationCopy:
				patch.Copy(tc.item.From, tc.item.Path)
			}

			result, _ := json.Marshal(patch)
			encoded := string(result)
			if encoded != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, encoded)
			}

		})
	}
}

func TestChain(t *testing.T) {
	patch := (jsonpatch.New().
		Test("/a/b/c", "foo").
		Add("/a/b/e", []string{"foo", "bar"}).
		Replace("/a/b/f", 42).
		Move("/a/b/g", "/a/b/h").
		Copy("/a/b/i", "/a/b/j"))
	expected := `[{"op":"test","path":"/a/b/c","value":"foo"},{"op":"add","path":"/a/b/e","value":["foo","bar"]},{"op":"replace","path":"/a/b/f","value":42},{"op":"move","from":"/a/b/g","path":"/a/b/h"},{"op":"copy","from":"/a/b/i","path":"/a/b/j"}]`
	result, _ := json.Marshal(patch)
	encoded := string(result)
	if encoded != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
