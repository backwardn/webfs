package wrds

import (
	"context"
	"fmt"
	"testing"

	"github.com/brendoncarroll/webfs/pkg/stores"
	"github.com/brendoncarroll/webfs/pkg/webref"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ctx = context.TODO()

func TestIndexOf(t *testing.T) {
	tree := Tree{
		Entries: []*TreeEntry{
			{Key: []byte("bbb")},
			{Key: []byte("ccc")},
			{Key: []byte("ddd")},
		},
	}
	var i int
	i = tree.indexPut([]byte("aaa"))
	assert.Equal(t, 0, i)
	i = tree.indexPut([]byte("eee"))
	assert.Equal(t, 3, i)
}

func TestPutGet(t *testing.T) {
	const N = 1000
	s := &webref.BasicStore{stores.NewMemStore(4096)}

	tree := NewTree()
	var err error
	for i := 0; i < N; i++ {
		key := []byte(fmt.Sprintf("key%03d", i))
		ref := &webref.Ref{}
		tree, err = tree.Put(ctx, s, key, ref)
		require.Nil(t, err)
	}
	t.Log("tree root has", len(tree.Entries), "entries")
	t.Log("tree entries", tree.Entries)

	for i := 0; i < N; i++ {
		key := []byte(fmt.Sprintf("key%03d", i))
		ent, err := tree.Get(ctx, s, key)
		require.Nil(t, err)
		if assert.NotNil(t, ent) {
			assert.Equal(t, key, ent.Key)
		} else {
			t.Log("missing", string(key))
		}
	}
}

func TestMinGt(t *testing.T) {
	const N = 1000
	s := &webref.BasicStore{stores.NewMemStore(4096)}

	tree := NewTree()
	var err error
	for i := 0; i < N; i++ {
		key := []byte(fmt.Sprintf("key%03d", i))
		ref := &webref.Ref{}
		tree, err = tree.Put(ctx, s, key, ref)
		require.Nil(t, err)
	}

	for i := 0; i < N-1; i++ {
		keyStr := fmt.Sprintf("key%03d", i)
		next := []byte(fmt.Sprintf("key%03d", i+1))
		ent, err := tree.MinGt(ctx, s, []byte(keyStr+".000"))
		require.Nil(t, err)
		require.Equal(t, ent.Key, next)
	}

	ent, err := tree.MinGt(ctx, s, []byte("key999"))
	require.Nil(t, err)
	assert.Nil(t, ent)
}

func TestDelete(t *testing.T) {
	s := webref.NewMemStore(4096)

	tree := NewTree()

	const N = 1000
	var err error
	for i := 0; i < N; i++ {
		key := []byte(fmt.Sprintf("key%03d", i))
		ref := &webref.Ref{}
		tree, err = tree.Put(ctx, s, key, ref)
		require.Nil(t, err)
	}

	deleteKeys := []string{
		"key004",
		"key000",
		"key849",
		"key999",
	}

	for _, key := range deleteKeys {
		tree2, err := tree.Delete(ctx, s, []byte(key))
		require.Nil(t, err)
		ent, err := tree2.Get(ctx, s, []byte(key))
		require.Nil(t, err)
		require.Nil(t, ent)
	}
}

func TestReplace(t *testing.T) {
	s := webref.NewMemStore(4096)
	tree := NewTree()

	const N = 10
	key := []byte("key")
	for i := 0; i < N; i++ {
		u := fmt.Sprint("url", i)
		ref := &webref.Ref{Ref: &webref.Ref_Url{u}}
		tree2, err := tree.Put(ctx, s, []byte(key), ref)
		require.Nil(t, err)

		assert.Len(t, tree2.Entries, 1)

		tree = tree2
	}
}
