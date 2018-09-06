package hashmap

import "errors"

// Key is a hash map key type
type Key string

// HashMaper is a common interface for hash map
type HashMaper interface {
  Set(key Key, value interface{}) error
  Get(key Key) (value interface{}, err error)
  Unset(key Key) error
  Count() int
}

// HashFunc is a hash funcion to calculate bucket number
// based on block size and hash key
type HashFunc func(blockSize int, key Key) int

// ErrNotFound returns when there is no such a key
var ErrNotFound = errors.New("hashmap: key not found")
// ErrInvalidIndex returns when hash function provides
// with invalid bucket number
var ErrInvalidIndex = errors.New("hashmap: invalid index")

// NewHashMap creates a hash map
func NewHashMap(blockSize int, fn HashFunc) HashMaper {
  if blockSize <= 0 {
    return nil
  }
  if fn == nil {
    fn = DefaultHashFunc
  }
  return &HashMap{
    buckets: make([][]keyValuePair, blockSize),
    fn: fn,
    size: blockSize,
  }
}

// HashMap is a hash map that implements HashMaper interface
type HashMap struct {
  buckets [][]keyValuePair
  fn HashFunc
  length int
  size int
}

// Set inserts or update value by key
func (m *HashMap) Set(key Key, value interface{}) error {
  index, err := m.getIndex(key)
  if err != nil {
    return err
  }
  b := m.buckets[index]
  for i, kv := range b {
    if kv.key == key {
      kv.val = value
      b[i] = kv
      return nil
    }
  }
  kv := keyValuePair{key, value}
  b = append(b, kv)
  m.buckets[index] = b
  m.length++
  return nil
}

// Get searches for a key in a hash map and returns it's value
// or returns ErrNotFound error if there is no such key
func (m *HashMap) Get(key Key) (value interface{}, err error) {
  index, err := m.getIndex(key)
  if err != nil {
    return nil, err
  }
  b := m.buckets[index]
  for _, kv := range b {
    if kv.key == key {
      return kv.val, nil
    }
  }
  return nil, ErrNotFound
}

// Unset removes key from a hash map
func (m *HashMap) Unset(key Key) error {
  index, err := m.getIndex(key)
  if err != nil {
    return err
  }
  b := m.buckets[index]
  for i, kv := range b {
    if kv.key == key {
      last := len(b) - 1
      b[i], b[last] = b[last], b[i]
      b = b[:last]
      m.buckets[index] = b
      m.length--
      return nil
    }
  }
  return nil
}

// Count returns number of elements in hash map
func (m *HashMap) Count() int {
  return m.length
}

func (m *HashMap) getIndex(key Key) (index int, err error) {
  index = m.fn(m.size, key)
  if index < 0 || index >= m.size {
    return -1, ErrInvalidIndex
  }
  return index, nil
}

type keyValuePair struct {
  key Key
  val interface{}
}

// DefaultHashFunc is a default hash function
// and applied if no hash function provided for hash map
func DefaultHashFunc(blockSize int, key Key) int {
  h := 0
  for i := 0; i < len(key); i++ {
    h = (h * 31 + int(key[i])) % 1000000000
  }
  h %= blockSize
  return h
}
