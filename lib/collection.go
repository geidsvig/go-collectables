package lib

import (
	"errors"
	"reflect"
)

type Collection struct {
	slice []interface{}
}

func (c *Collection) Append(t reflect.Type, val interface{}) error {
	if t != reflect.TypeOf(val) {
		return errors.New("type missmatch")
	}
	c.slice = append(c.slice, val)
	return nil
}

func (c *Collection) Prepend(t reflect.Type, val interface{}) error {
	if t != reflect.TypeOf(val) {
		return errors.New("type missmatch")
	}
	c.slice = append([]interface{}{val}, c.slice...)
	return nil
}

func (c *Collection) GetFirst(t reflect.Type) (interface{}, error) {
	if len(c.slice) == 0 {
		return nil, errors.New("no element for index")
	}
	result := c.slice[0]
	if reflect.TypeOf(result) != t {
		return nil, errors.New("type missmatch")
	}
	return result, nil
}

func (c *Collection) RemoveFirst() error {
	if len(c.slice) == 0 {
		return errors.New("no element for index")
	}
	c.slice = c.slice[1:]
	return nil
}

func (c *Collection) GetLast(t reflect.Type) (interface{}, error) {
	if len(c.slice) == 0 {
		return nil, errors.New("no element for index")
	}
	result := c.slice[len(c.slice)-1]
	if reflect.TypeOf(result) != t {
		return nil, errors.New("type missmatch")
	}
	return result, nil
}

func (c *Collection) RemoveLast() error {
	if len(c.slice) > 0 {
		c.slice = c.slice[:len(c.slice)-1]
		return nil
	}
	return errors.New("no element for index")
}

func (c *Collection) GetAt(t reflect.Type, index int) (interface{}, error) {
	if index >= 0 && index < len(c.slice) {
		result := c.slice[index]
		if reflect.TypeOf(result) != t {
			return nil, errors.New("type missmatch")
		}
		return result, nil
	}
	return 0, errors.New("no element for index")
}

func (c *Collection) RemoveAt(index int) error {
	if index >= 0 && index < len(c.slice) {
		c.slice = append(c.slice[:index], c.slice[index+1:]...)
		return nil
	}
	return errors.New("no element for index")
}

func (c *Collection) IndexOf(val interface{}) int {
	for i, v := range c.slice {
		if val == v {
			return i
		}
	}
	return -1
}

func (c *Collection) IsEmpty() bool {
	return len(c.slice) == 0
}

func (c *Collection) Length() int {
	return len(c.slice)
}

/*
@param func the modifyer for each of the items in the collection.
@return any collection type you want, be sure to check your types with reflect.TypeOf
*/
func (c *Collection) Map(f func(interface{}) interface{}) *Collection {
	newSlice := make([]interface{}, 0, len(c.slice))
	for _, v := range c.slice {
		newSlice = append(newSlice, f(v))
	}
	return &Collection{slice: newSlice}
}

func (c *Collection) Foreach(f func(interface{})) {
	for _, v := range c.slice {
		f(v)
	}
}

func (c *Collection) Filter(fn func(interface{}) bool) *Collection {
	filtered := &Collection{}
	for _, item := range c.slice {
		if fn(item) {
			filtered.slice = append(filtered.slice, item)
		}
	}
	return filtered
}
