package local_cache

import "io"

type FastLocalCaching interface {
	SyncDB(remoteVersion Version) (deltaData []byte, err error)
	Get(key interface{}) (value interface{}, err error)
	Put(key interface{}, value interface{}) error
	Del(key interface{}) error
	BatchGet(keys []interface{}) (values map[interface{}]interface{}, err error)
	BatchPut(values map[interface{}]interface{}) error
	BatchDel(keys []interface{}) error
	Merge() error
	Init(Path string, opts ...Option) error
}

type Caching interface {
	Get(key []byte) (value []byte, err error)
	Put(key, value []byte) error
	Del(key []byte) error
}

type EntryCodec interface {
	Encode(entry *Entry) ([]byte, error)
	Decode(data []byte) (*Entry, error)
}

type KVCodec interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte) (interface{}, error)
}

type DataFileReader interface {
	ReadDataFile(reader io.Reader) (*DataFile, error)
}

type DataFileWriter interface {
	WriteDataFile(dataFile *DataFile, writer io.Writer) error
}

type Monitor interface {
	Mon(name string, n ...int)
}
