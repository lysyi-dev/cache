package cache

type Item struct {
	Value interface{}
}

type Cache struct {
	items map[string]Item
}

func New() *Cache {
	return &Cache{
		items: make(map[string]Item),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = Item{
		Value: value,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	return item.Value, true
}

func (c *Cache) Delete(key string) {
	delete(c.items, key)
}

func (c *Cache) Clear() {
	c.items = make(map[string]Item)
}
