package fastpebble

import "github.com/tachunwu/fastpebble/pkg/entity"

type Service interface {
	Batch(batch *entity.BatchRequest) (*[]entity.KeyValue, error)
	Scan(sr entity.ScanRequest) (*[]entity.KeyValue, error)
	Set(kv *entity.KeyValue) (*entity.KeyValue, error)
	Get(key string) (*entity.KeyValue, error)
	Delete(key string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Batch(batch *entity.BatchRequest) (*[]entity.KeyValue, error) {
	return s.repository.Batch(batch)
}
func (s *service) Scan(sr entity.ScanRequest) (*[]entity.KeyValue, error) {
	return s.repository.Scan(sr)
}
func (s *service) Set(kv *entity.KeyValue) (*entity.KeyValue, error) {
	return s.repository.Set(kv)
}
func (s *service) Get(key string) (*entity.KeyValue, error) {
	return s.repository.Get(key)
}
func (s *service) Delete(key string) error {
	return s.repository.Delete(key)
}
