package app

import "sync"

// report storage object
type reportStore struct {
	sync.RWMutex
	reports map[string][]byte
}

// IsExist checks if the report is in the storage
func (s *reportStore) IsExist(key string) ([]byte, bool) {
	s.RLock()
	defer s.RUnlock()
	data, ok := s.reports[key]
	return data, ok
}
