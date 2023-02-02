package storage

import (
	"errors"
	"sync"
)

type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func New[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

func (s *SafeMap[K, V]) Insert(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *SafeMap[K, V]) Get(key K) (V, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]
	if !ok {
		return value, errors.New("key not found")
	}

	return value, nil
}

func (s *SafeMap[K, V]) Update(key K, value V) error {
	if !s.Has(key) {
		return errors.New("key not found")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	return nil
}

func (s *SafeMap[K, V]) Delete(key K) error {
	if !s.Has(key) {
		return errors.New("key not found")
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)

	return nil
}

func (s *SafeMap[K, V]) Has(key K) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.data[key]
	return ok
}
