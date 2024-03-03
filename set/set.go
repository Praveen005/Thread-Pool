package set

import "sync"

type Set struct {
	_map *sync.Map
}

// creates and returns a new set
func NewSet() *Set {
	set := new(Set)
	set._map = new(sync.Map)
	return set
}

//properties of set:
// stores each element just once

//adding values to the set

func (s *Set) Add(value interface{}) {
	s._map.Store(value, true)
}

// Removing a value from the set

func (s *Set) Remove(key interface{}) {
	s._map.Delete(key)
}

// check if the key is present or nor
// interface{} satisfies all the data types, cz it has no underlying method that the data type has to satisfy
func (s *Set) Contains(key interface{}) bool {
	_, ok := s._map.Load(key)
	return ok
}

//returns all the values as array
func (s *Set) GetAll() []interface{} {
	values := make([]interface{}, 0)

	s._map.Range(func(key, value interface{}) bool {
		values = append(values, key)
		return true
	})
	return values
}

// Return all the values as an array of string
func (s *Set) GetAllAsString() []string {
	// making an array
	values := make([]string, 0)
	s._map.Range(func(key interface{}, value interface{}) bool {
		values = append(values, key.(string)) //dekh lo type assertion syntax
		return true
	})

	return values
}

// It returns at max n data
func (s *Set) GetAllWithLimit(n int) []interface{} {
	values := make([]interface{}, 0)

	s._map.Range(func(key interface{}, value interface{}) bool {
		values = append(values, key)
		if n--; n > 0 {
			return true // continue iteration
		}
		return false // stop iteration
	})
	return values
}
