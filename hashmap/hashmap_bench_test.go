package hashmap

import "math/rand"
import "strconv"
import "testing"

const NumberOfOperations = 1000000
const KeyRange = 10000

func randomKey() Key {
  return Key(strconv.Itoa(rand.Intn(KeyRange)))
}

func benchHashMapSize(size int, b *testing.B) {
  m := NewHashMap(size, nil)
  for i := 0; i < b.N; i++ {
    m.Set(Key(i), i)
  }
  for i := 0; i < b.N; i++ {
    m.Get(Key(i))
  }
  for i := 0; i < b.N; i++ {
    m.Unset(Key(i))
  }
}

func BenchmarkHashMap(b *testing.B) {
  sizes := []int{16, 64, 128, 1024, 16384}
  for _, size := range sizes {
    prefix := strconv.Itoa(size) + "-"
    m := NewHashMap(size, nil)
    b.Run(prefix + "SET", func(b *testing.B) {
      b.N = NumberOfOperations
      for i := 0; i < b.N; i++ {
        m.Set(randomKey(), i)
      }
    })
    b.Run(prefix + "GET", func(b *testing.B) {
      b.N = NumberOfOperations
      for i := 0; i < b.N; i++ {
        m.Get(randomKey())
      }
    })
    b.Run(prefix + "UNSET", func(b *testing.B) {
      b.N = NumberOfOperations
      for i := 0; i < b.N; i++ {
        m.Unset(randomKey())
      }
    })
  }
}

func BenchmarkNativeMap(b *testing.B) {
  m := map[Key]interface{}{}
  b.Run("SET", func(b *testing.B) {
    b.N = NumberOfOperations
    for i := 0; i < b.N; i++ {
      m[randomKey()] = i
    }
  })
  b.Run("GET", func(b *testing.B) {
    b.N = NumberOfOperations
    for i := 0; i < b.N; i++ {
      _ = m[randomKey()]
    }
  })
  b.Run("UNSET", func(b *testing.B) {
    b.N = NumberOfOperations
    for i := 0; i < b.N; i++ {
      delete(m, randomKey())
    }
  })
}

func BenchmarkDefaultHashFunc(b *testing.B) {
  sizes := []int{16, 64, 128, 1024, 16384}
  for _, size := range sizes {
    prefix := strconv.Itoa(size)
    b.Run(prefix, func(b *testing.B) {
      b.N = NumberOfOperations
      for i := 0; i < b.N; i++ {
        DefaultHashFunc(size, randomKey())
      }
    })
  }
}
