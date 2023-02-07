package lib

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := Stack{}
	dog1 := Dog{Name: "dog1", Breed: "breed1"}
	dog2 := Dog{Name: "dog2", Breed: "breed2"}
	dog3 := Dog{Name: "dog3", Breed: "breed3"}

	if err := stack.Push(typeDog, dog1); err != nil {
		t.Errorf("Error pushing item to stack: %v", err)
	}
	if err := stack.Push(typeDog, dog2); err != nil {
		t.Errorf("Error pushing item to stack: %v", err)
	}
	if err := stack.Push(typeDog, dog3); err != nil {
		t.Errorf("Error pushing item to stack: %v", err)
	}

	if stack.Length() != 3 {
		t.Errorf("Expected length of stack to be 3, got %d", stack.Length())
	}
}

func TestPop(t *testing.T) {
	stack := Stack{collection: Collection{}}
	dog1 := Dog{Name: "dog1", Breed: "breed1"}
	dog2 := Dog{Name: "dog2", Breed: "breed2"}
	dog3 := Dog{Name: "dog3", Breed: "breed3"}

	// Push three dogs to the stack
	err := stack.Push(typeDog, dog1)
	if err != nil {
		t.Fatalf("Error pushing to the stack: %v", err)
	}
	err = stack.Push(typeDog, dog2)
	if err != nil {
		t.Fatalf("Error pushing to the stack: %v", err)
	}
	err = stack.Push(typeDog, dog3)
	if err != nil {
		t.Fatalf("Error pushing to the stack: %v", err)
	}

	// Pop the third dog from the stack
	result, err := stack.Pop(typeDog)
	if err != nil {
		t.Fatalf("Error popping from the stack: %v", err)
	}
	if result.(Dog).Name != "dog3" {
		t.Fatalf("Expected the last dog to be popped, but got %v", result.(Dog))
	}

	// Pop the second dog from the stack
	result, err = stack.Pop(typeDog)
	if err != nil {
		t.Fatalf("Error popping from the stack: %v", err)
	}
	if result.(Dog).Name != "dog2" {
		t.Fatalf("Expected the second dog to be popped, but got %v", result.(Dog))
	}

	// Pop the first dog from the stack
	result, err = stack.Pop(typeDog)
	if err != nil {
		t.Fatalf("Error popping from the stack: %v", err)
	}
	if result.(Dog).Name != "dog1" {
		t.Fatalf("Expected the first dog to be popped, but got %v", result.(Dog))
	}

	// Try to pop an item from an empty stack
	_, err = stack.Pop(typeDog)
	if err == nil {
		t.Fatalf("Expected an error when trying to pop an item from an empty stack, but got nil")
	}
}
