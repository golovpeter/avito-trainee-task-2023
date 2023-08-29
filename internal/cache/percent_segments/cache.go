package percent_segments

import "github.com/golovpeter/avito-trainee-task-2023/internal/repository/segments"

type Cache struct {
	elements map[string]segments.Segment
}

func NewCache() *Cache {
	return &Cache{
		elements: make(map[string]segments.Segment),
	}
}

func (c *Cache) Set(slug string, newSegment segments.Segment) {
	c.elements[slug] = newSegment
}

func (c *Cache) Update(segments map[string]segments.Segment) {
	c.elements = segments
}

func (c *Cache) Get() map[string]segments.Segment {
	return c.elements
}