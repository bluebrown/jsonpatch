package jsonpatch_test

import (
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

			patch := jsonpatch.Patch{}

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

			encoded := patch.Encode().String()
			if encoded != tc.expected+"\n" {
				t.Errorf("expected %s, got %s", tc.expected, encoded)
			}

		})
	}
}
