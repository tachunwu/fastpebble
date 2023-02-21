package entity

const (
	Set = iota
	Delete
)

type KeyValue struct {
	Key   string `json:"key"`
	Value []byte `json:"value,omitempty"`
}

type Value struct {
	Value []byte `json:"value,omitempty"`
}

type ScanRequest struct {
	Key     string `json:"key"`
	Count   int64  `json:"count"`
	Reverse bool   `json:"reverse"`
}

type BatchRequest struct {
	Operations []Operation `json:"operations"`
}

type Operation struct {
	Type int `json:"type"`
	KeyValue
}
