package lib

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := &Queue{}
	err := q.Enqueue(typeDog, Dog{"Max", "German Shepherd"})
	if err != nil {
		t.Fatalf("Failed to enqueue: %v", err)
	}
	err = q.Enqueue(typeDog, Dog{"Rocky", "Labrador"})
	if err != nil {
		t.Fatalf("Failed to enqueue: %v", err)
	}

	expectedLen := 2
	if q.Length() != expectedLen {
		t.Fatalf("Expected length of queue to be %d, got %d", expectedLen, q.Length())
	}
}

func TestDequeue(t *testing.T) {
	q := &Queue{}
	q.Enqueue(typeDog, Dog{"Max", "German Shepherd"})
	q.Enqueue(typeDog, Dog{"Rocky", "Labrador"})

	dog, err := q.Dequeue(typeDog)
	if err != nil {
		t.Fatalf("Failed to dequeue: %v", err)
	}
	expectedDog := Dog{"Max", "German Shepherd"}
	if dog != expectedDog {
		t.Fatalf("Expected dog to be %v, got %v", expectedDog, dog)
	}

	expectedLen := 1
	if q.Length() != expectedLen {
		t.Fatalf("Expected length of queue to be %d, got %d", expectedLen, q.Length())
	}
}

func TestDequeueEmpty(t *testing.T) {
	q := &Queue{}
	_, err := q.Dequeue(typeDog)
	if err == nil {
		t.Fatalf("Expected an error when dequeueing from an empty queue")
	}
}
