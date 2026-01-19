package pokecache

import (
"testing"
"time"
)

func TestCacheAddGet(t *testing.T) {
cache := NewCache(5 * time.Second)

cache.Add("key", []byte("value"))

val, ok := cache.Get("key")
if !ok {
t.Fatal("expected cache hit")
}

if string(val) != "value" {
t.Fatalf("expected value, got %s", val)
}
}
