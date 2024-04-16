package engine

type Store interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	Delete(key string)
	Clear()
}

type MapStore struct {
	data map[string]interface{}
}

func NewMapStore() *MapStore {
	return &MapStore{
		data: make(map[string]interface{}),
	}
}

func (s *MapStore) Set(key string, value interface{}) {
	s.data[key] = value
}

func (s *MapStore) Get(key string) (interface{}, bool) {
	value, ok := s.data[key]
	return value, ok
}

func (s *MapStore) Delete(key string) {
	delete(s.data, key)
}

func (s *MapStore) Clear() {
	s.data = make(map[string]interface{})
}
