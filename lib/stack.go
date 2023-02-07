package lib

import (
	"reflect"
)

type Stack struct {
	collection Collection
}

func (s *Stack) Push(t reflect.Type, item interface{}) error {
	return s.collection.Append(t, item)
}

func (s *Stack) Pop(t reflect.Type) (interface{}, error) {
	result, err := s.collection.GetLast(t)
	if err == nil {
		err = s.collection.RemoveLast()
		if err == nil {
			return result, nil
		}
		return nil, err
	}
	return nil, err
}

func (s *Stack) Length() int {
	return s.collection.Length()
}
