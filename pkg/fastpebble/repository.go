package fastpebble

import (
	"github.com/cockroachdb/pebble"
	"github.com/tachunwu/fastpebble/pkg/entity"
	"github.com/tachunwu/fastpebble/pkg/storage"
	pebbleDB "github.com/tachunwu/fastpebble/pkg/storage/pebble"
)

type Repository interface {
	Batch(batch *entity.BatchRequest) (*[]entity.KeyValue, error)
	Scan(sr entity.ScanRequest) (*[]entity.KeyValue, error)
	Set(kv *entity.KeyValue) (*entity.KeyValue, error)
	Get(key string) (*entity.KeyValue, error)
	Delete(key string) error
}

func NewRepository() Repository {
	return &repository{
		s: pebbleDB.NewPebbleDB("./data"),
	}
}

type repository struct {
	s storage.Storage
}

func (r *repository) Batch(batch *entity.BatchRequest) (*[]entity.KeyValue, error) {
	b := r.s.NewBatch()
	kvs := []entity.KeyValue{}

	for _, op := range batch.Operations {
		switch op.Type {
		case entity.Set:
			b.Set([]byte(op.Key), op.Value, nil)
			kvs = append(kvs, op.KeyValue)
		case entity.Delete:
			b.Delete([]byte(op.Key), nil)
			kvs = append(kvs, op.KeyValue)
		}
	}
	if err := b.Commit(pebble.Sync); err != nil {
		return nil, b.Close()
	}
	return &kvs, nil
}
func (r *repository) Scan(sr entity.ScanRequest) (*[]entity.KeyValue, error) {
	iter := r.s.NewIter(nil)
	kvs := []entity.KeyValue{}

	r.s.Scan(iter, []byte(sr.Key), sr.Count, sr.Reverse)
	if !iter.Valid() {
		return nil, iter.Close()
	}

	for iter.First(); iter.Valid(); iter.Next() {
		kv := entity.KeyValue{
			Key:   string(iter.Key()),
			Value: iter.Value(),
		}
		kvs = append(kvs, kv)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return &kvs, nil
}
func (r *repository) Set(kv *entity.KeyValue) (*entity.KeyValue, error) {
	b := r.s.NewBatch()
	b.Set([]byte(kv.Key), kv.Value, pebble.Sync)
	b.Commit(pebble.Sync)
	return &entity.KeyValue{
		Key:   kv.Key,
		Value: kv.Value,
	}, nil
}
func (r *repository) Get(key string) (*entity.KeyValue, error) {
	iter := r.s.NewIter(nil)
	defer iter.Close()
	r.s.Scan(iter, []byte(key), 1, false)
	iter.SeekGE([]byte(key))

	if !iter.Valid() {
		return nil, iter.Close()
	}
	return &entity.KeyValue{
		Key:   string(iter.Key()),
		Value: iter.Value(),
	}, nil
}
func (r *repository) Delete(key string) error {
	b := r.s.NewBatch()
	err := b.Delete([]byte(key), pebble.Sync)
	if err != nil {
		return err
	}
	if err := b.Commit(pebble.Sync); err != nil {
		return err
	}
	return nil
}
