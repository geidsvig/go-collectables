package lib

import (
	"reflect"
	"strings"
	"testing"
)

type Dog struct {
	Name  string
	Breed string
}

var (
	typeDog = reflect.TypeOf(Dog{})
	typeInt = reflect.TypeOf(1)
)

func TestAppend(t *testing.T) {
	c := Collection{}
	err := c.Append(typeInt, 1)
	if err != nil {
		t.Errorf("Failed to append int: %s", err)
	}

	dog := Dog{"Luna", "Aussie"}
	err = c.Append(typeDog, dog)
	if err != nil {
		t.Errorf("Failed to append Dog struct: %s", err)
	}
}

func TestPrepend(t *testing.T) {
	c := Collection{}
	err := c.Prepend(typeInt, 1)
	if err != nil {
		t.Errorf("Failed to prepend int: %s", err)
	}

	dog := Dog{"Luna", "Aussie"}
	err = c.Prepend(typeDog, dog)
	if err != nil {
		t.Errorf("Failed to prepend Dog struct: %s", err)
	}
}

func TestGetFirst(t *testing.T) {
	c := Collection{[]interface{}{1, 2}}

	val, err := c.GetFirst(typeInt)
	if err != nil {
		t.Errorf("Failed to get first int: %s", err)
	}
	if val != 1 {
		t.Errorf("Incorrect first int value: expected 1, got %d", val)
	}

	c = Collection{[]interface{}{Dog{"Luna", "Aussie"}}}
	dog, err := c.GetFirst(typeDog)
	if err != nil {
		t.Errorf("Failed to get first Dog struct: %s", err)
	}
	d := Dog{"Luna", "Aussie"}
	if dog != d {
		t.Errorf("Incorrect first Dog struct value: expected {Rufus Labrador}, got %+v", dog)
	}
}

func TestRemoveFirst(t *testing.T) {
	c := Collection{[]interface{}{1, 2}}
	err := c.RemoveFirst()
	if err != nil {
		t.Errorf("Failed to remove first element: %s", err)
	}

	if len(c.slice) != 1 {
		t.Errorf("Incorrect slice length after removing first element: expected 1, got %d", len(c.slice))
	}
}

func TestGetLast(t *testing.T) {
	col := &Collection{slice: []interface{}{1, 2, 3}}
	result, err := col.GetLast(typeInt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.(int) != 3 {
		t.Errorf("Expected 3, but got %v", result)
	}

	d := Dog{"Buddy", "Labrador"}
	col = &Collection{slice: []interface{}{d}}
	result, err = col.GetLast(typeDog)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.(Dog) != d {
		t.Errorf("Expected %v, but got %v", d, result)
	}
}

func TestRemoveLast(t *testing.T) {
	col := &Collection{slice: []interface{}{1, 2, 3}}
	err := col.RemoveLast()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(col.slice) != 2 {
		t.Errorf("Expected slice to have length 2, but got %v", len(col.slice))
	}

	d := Dog{"Buddy", "Labrador"}
	col = &Collection{slice: []interface{}{d}}
	err = col.RemoveLast()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(col.slice) != 0 {
		t.Errorf("Expected slice to have length 0, but got %v", len(col.slice))
	}
}

func TestGetAt(t *testing.T) {
	col := &Collection{slice: []interface{}{1, 2, 3}}
	result, err := col.GetAt(typeInt, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.(int) != 2 {
		t.Errorf("Expected 2, but got %v", result)
	}

	d := Dog{"Buddy", "Labrador"}
	col = &Collection{slice: []interface{}{d}}
	result, err = col.GetAt(typeDog, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.(Dog) != d {
		t.Errorf("Expected %v, but got %v", d, result)
	}
}

func TestRemoveAt(t *testing.T) {
	c := Collection{slice: []interface{}{1, 2, 3}}
	err := c.RemoveAt(2)
	if err != nil {
		t.Error("Failed to remove at index 2:", err)
	}
	if len(c.slice) != 2 {
		t.Error("Failed to remove at index 2, unexpected length:", len(c.slice))
	}
	if c.slice[0] != 1 || c.slice[1] != 2 {
		t.Error("Failed to remove at index 2, unexpected slice:", c.slice)
	}
	err = c.RemoveAt(5)
	if err == nil {
		t.Error("Expected error for removing at index 5, got nil")
	}
}

func TestIndexOf(t *testing.T) {
	c := Collection{slice: []interface{}{1, 2, 3}}
	index := c.IndexOf(2)
	if index != 1 {
		t.Error("Failed to find index of 2, got:", index)
	}
	index = c.IndexOf(5)
	if index != -1 {
		t.Error("Unexpected index of 5, got:", index)
	}
	dog := &Dog{"dog1", "breed1"}
	c = Collection{slice: []interface{}{dog, &Dog{"dog2", "breed2"}}}
	index = c.IndexOf(dog)
	if index != 0 {
		t.Error("Failed to find index of dog, got:", index)
	}
	index = c.IndexOf(&Dog{"dog3", "breed3"})
	if index != -1 {
		t.Error("Unexpected index of dog, got:", index)
	}
}

func TestIsEmpty(t *testing.T) {
	c := Collection{slice: []interface{}{1, 2, 3}}
	if c.IsEmpty() {
		t.Error("Unexpected IsEmpty, got true")
	}
	c = Collection{slice: []interface{}{}}
	if !c.IsEmpty() {
		t.Error("Unexpected IsEmpty, got false")
	}
}

func TestLength(t *testing.T) {
	c := Collection{slice: []interface{}{1, 2, 3}}
	if c.Length() != 3 {
		t.Error("Unexpected length, got:", c.Length())
	}
	c = Collection{slice: []interface{}{}}
	if c.Length() != 0 {
		t.Error("Unexpected length, got:", c.Length())
	}
}

func TestMap(t *testing.T) {
	// Test Map on int type
	col := Collection{slice: []interface{}{1, 2, 3}}
	double := func(x interface{}) interface{} {
		return x.(int) * 2
	}
	doubledCol := col.Map(double)
	expected := &Collection{slice: []interface{}{2, 4, 6}}
	if !reflect.DeepEqual(doubledCol, expected) {
		t.Errorf("Expected %v, but got %v", expected, doubledCol)
	}

	// Test Map on Dog type
	col = Collection{slice: []interface{}{
		Dog{Name: "dog1", Breed: "breed1"},
		Dog{Name: "dog2", Breed: "breed2"},
		Dog{Name: "dog3", Breed: "breed3"},
	}}
	upperName := func(x interface{}) interface{} {
		d := x.(Dog)
		return Dog{
			Name:  strings.ToUpper(d.Name),
			Breed: d.Breed,
		}
	}
	upperNameCol := col.Map(upperName)
	expected = &Collection{slice: []interface{}{
		Dog{Name: "DOG1", Breed: "breed1"},
		Dog{Name: "DOG2", Breed: "breed2"},
		Dog{Name: "DOG3", Breed: "breed3"},
	}}
	if !reflect.DeepEqual(upperNameCol, expected) {
		t.Errorf("Expected %v, but got %v", expected, upperNameCol)
	}
}

func TestForeach(t *testing.T) {
	// Test Foreach on int type
	col := Collection{slice: []interface{}{1, 2, 3}}
	var sum int
	sumInts := func(x interface{}) {
		sum += x.(int)
	}
	col.Foreach(sumInts)
	expected := 6
	if sum != expected {
		t.Errorf("Expected %v, but got %v", expected, sum)
	}

	// Test Foreach on Dog type
	col = Collection{slice: []interface{}{
		Dog{Name: "dog1", Breed: "breed1"},
		Dog{Name: "dog2", Breed: "breed2"},
		Dog{Name: "dog3", Breed: "breed3"},
	}}
	var names string
	concatenateNames := func(x interface{}) {
		names += x.(Dog).Name
	}
	col.Foreach(concatenateNames)
	expectedStr := "dog1dog2dog3"
	if names != expectedStr {
		t.Errorf("Expected %v, but got %v", expectedStr, names)
	}
}

func TestFilter(t *testing.T) {
	// Test Filter with ints
	c := Collection{}
	c.Append(typeInt, 1)
	c.Append(typeInt, 2)
	c.Append(typeInt, 3)
	c.Append(typeInt, 4)

	filteredInts := c.Filter(func(val interface{}) bool {
		return val.(int) >= 3
	})
	if filteredInts.Length() != 2 {
		t.Errorf("Expected filteredInts length to be 2, but got %d", filteredInts.Length())
	}
	if v, err := filteredInts.GetAt(typeInt, 0); err != nil || v.(int) != 3 {
		t.Errorf("Expected first item in filteredInts to be 3, but got %v", v)
	}
	if v, err := filteredInts.GetAt(typeInt, 1); err != nil || v.(int) != 4 {
		t.Errorf("Expected second item in filteredInts to be 4, but got %v", v)
	}

	// Test Filter with Dogs
	c = Collection{}
	dog1 := Dog{"dog1", "breed1"}
	dog2 := Dog{"dog2", "breed2"}
	dog3 := Dog{"dog3", "breed3"}
	c.Append(typeDog, dog1)
	c.Append(typeDog, dog2)
	c.Append(typeDog, dog3)

	filteredDogs := c.Filter(func(val interface{}) bool {
		return val.(Dog).Name == "dog2" || val.(Dog).Breed == "breed3"
	})
	if filteredDogs.Length() != 2 {
		t.Errorf("Expected filteredDogs length to be 2, but got %d", filteredDogs.Length())
	}
	if v, err := filteredDogs.GetAt(typeDog, 0); err != nil || v.(Dog) != dog2 {
		t.Errorf("Expected first item in filteredDogs to be dog2, but got %v", v)
	}
	if v, err := filteredDogs.GetAt(typeDog, 1); err != nil || v.(Dog) != dog3 {
		t.Errorf("Expected second item in filteredDogs to be dog3, but got %v", v)
	}
}
