## Cache package
This is a test package

### Here, I practiced with concepts such as:
* structures
* empty interface
* maps
* working with a single object (passing by reference)
* modules
* tests


### Here is the translated text for the method descriptions:
* `c.Get` - adds a value to the cache
* `c.Set` - retrieves a value by key
* `c.Delete` - deletes a value by key
* `c.Clear` - clears the entire cache

### Usage Example

```go
package main

import (
	"fmt"
	cache "github.com/lysyi-dev/cache"
)

func main() {
	c := cache.New()

	// Add a value to the cache
	c.Set("some_key", "some_value")

	// Retrieve a value by key
	if val, ok := c.Get("some_key"); ok {
		fmt.Println("Found:", val)
	} else {
		fmt.Println("Not ok")
	}

	// Delete a value by key
	c.Delete("some_key")
	if val, ok := c.Get("some_key"); ok {
		fmt.Println("Found:", val)
	} else {
		fmt.Println("Not ok")
	}

	// Clear the entire cache
	c.Set("some_key_2", "some_value_2")
	c.Clear()
	if _, ok := c.Get("some_key_2"); !ok {
		fmt.Println("Cache cleared")
	}
}
```