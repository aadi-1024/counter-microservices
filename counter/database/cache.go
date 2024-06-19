package database

import (
	"log"
	"strconv"
	"time"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type Cache struct {
	client *memcache.Client
}

func NewCache() (*Cache, error) {
	cache := &Cache{}

	client := memcache.New("memcached:11211")
	cache.client = client
	var err error

	for i := 0; i < 5; i++ {
		err = client.Ping()
		if err == nil {
			log.Println("connected to memcached")
			break
		}
		log.Println("ping to cache failed, trying again")
		time.Sleep(5 * time.Second)
	}

	return cache, err
}

func (c *Cache) Get(uid int) (int, error) {
	item, err := c.client.Get(strconv.Itoa(uid))
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(item.Value))
}

func (c *Cache) Set(uid, value int) error {
	item := &memcache.Item{
		Key: fmt.Sprint(uid),
		Value: []byte(fmt.Sprint(value)),
		Expiration: 0,
	}
	return c.client.Set(item)
}