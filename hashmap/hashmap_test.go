package hashmap

import "testing"

import "github.com/stretchr/testify/require"

func TestHashMap(t *testing.T) {
  a := require.New(t)

  m := NewHashMap(1024, nil)
  a.NotNil(m)
  a.Equal(0, m.Count())

  // Add
  a.Nil(m.Set("foo", "red"))
  a.Equal(1, m.Count())
  val, err := m.Get("foo")
  a.Nil(err)
  a.Equal("red", val)

  // Update
  a.Nil(m.Set("foo", "black"))
  a.Equal(1, m.Count())
  val, err = m.Get("foo")
  a.Nil(err)
  a.Equal("black", val)

  // Add
  a.Nil(m.Set("bar", "black"))
  a.Nil(m.Set("baz", "yellow"))
  a.Equal(3, m.Count())
  val, err = m.Get("foo")
  a.Nil(err)
  a.Equal("black", val)
  val, err = m.Get("bar")
  a.Nil(err)
  a.Equal("black", val)
  val, err = m.Get("baz")
  a.Nil(err)
  a.Equal("yellow", val)

  // Unset
  a.Nil(m.Unset("bar"))
  a.Equal(2, m.Count())
  val, err = m.Get("foo")
  a.Nil(err)
  a.Equal("black", val)
  val, err = m.Get("bar")
  a.NotNil(err)
  a.Nil(val)
  val, err = m.Get("baz")
  a.Nil(err)
  a.Equal("yellow", val)

  // Unset nonexistent
  a.Nil(m.Unset("qux"))
  a.Equal(2, m.Count())
}

func TestInvalidSize(t *testing.T) {
  a := require.New(t)

  a.Nil(NewHashMap(0, nil))
  a.Nil(NewHashMap(-1, nil))
  a.Nil(NewHashMap(-1024, nil))
}

func TestWithHashFunc(t *testing.T) {
  a := require.New(t)

  m := NewHashMap(16, func(blockSize int, key Key) int {
    return 0
  })
  a.NotNil(m)
  a.Nil(m.Set("foo", "bar"))
  val, err := m.Get("foo")
  a.Nil(err)
  a.Equal("bar", val)
  a.Nil(m.Unset("foo"))
}

func TestInvalidHashFunc(t *testing.T) {
  a := require.New(t)

  m := NewHashMap(1024, func(blockSize int, key Key) int {
    // Not valid, cuase maximum index is `blockSize - 1`
    return blockSize
  })
  a.NotNil(m)
  a.NotNil(m.Set("foo", "bar"))
  val, err := m.Get("foo")
  a.NotNil(err)
  a.Nil(val)
  a.NotNil(m.Unset("foo"))
}
