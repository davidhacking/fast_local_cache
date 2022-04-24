package local_cache

type Entry struct {
	key     []byte
	delFlag bool
	value   []byte
	ts      int64
	crc     uint32
}

func (e *Entry) IsIntegrate() bool {
	// TODO
	return true
}
