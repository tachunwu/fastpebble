package storage

import "github.com/tachunwu/fastpebble/pkg/storage/pebble"

type Storage interface {
	pebble.DB
}
