package lib

import (
	"reflect"
)

type Queue struct {
	collection Collection
}

func (q *Queue) Enqueue(t reflect.Type, item interface{}) error {
	return q.collection.Append(t, item)
}

func (q *Queue) Dequeue(t reflect.Type) (interface{}, error) {
	result, err := q.collection.GetFirst(t)
	if err == nil {
		err = q.collection.RemoveLast()
		if err == nil {
			return result, nil
		}
		return nil, err
	}
	return nil, err
}

func (q *Queue) Length() int {
	return q.collection.Length()
}
