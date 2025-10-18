package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	interval := time.Second * 5
	cases := []struct {
		key  string
		data []byte
	}{
		{
			key:  "http://example.com",
			data: []byte("some test data"),
		},
		{
			key:  "https://example2.com",
			data: []byte("some test data encrypted"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.data)
			cachedData, exists := cache.Get(c.key)
			if !exists {
				t.Errorf("expected to find key")
				return
			}
			if string(cachedData) != string(c.data) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = time.Second * 5
	const waitTime = baseTime + time.Second*5
	cache := NewCache(baseTime)
	cache.Add("www.randomwebsite.com", []byte("mostrandomdataever"))
	_, exists := cache.Get("www.randomwebsite.com")

	if !exists {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, exists = cache.Get("www.randomwebsite.com")

	if exists {
		t.Errorf("expected to not find key")
		return
	}
}
