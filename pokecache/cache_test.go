package pokecache

import (
	"testing"
	"time"
)

func TestPokeCache(t *testing.T) {
	cache := Cache{}
	cache.NewCache(10 * time.Second)

	//get not existing
	_, ok := cache.Get("a")
	if ok {
		t.Fail()
	}

	//get existing
	cache.Add("a", CacheEntry{Val: []byte("test")})
	val, ok := cache.Get("a")
	if !ok {
		t.Fail()
	}
	if string(val) != "test" {
		t.Fail()
	}

	//update val
	cache.Add("a", CacheEntry{Val: []byte("hello")})
	val, ok = cache.Get("a")
	if !ok {
		t.Fail()
	}
	if string(val) != "hello" {
		t.Fail()
	}

	//get removed
	time.Sleep(11 * time.Second)
	val, ok = cache.Get("a")
	if ok {
		t.Fail()
	}

}
