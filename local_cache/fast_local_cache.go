package local_cache

import "sync"

type FastLocalCache struct {
	sync.RWMutex
	path     string
	cfg      *Config
	datafile *DataFile
}

func NewFastLocalCache() FastLocalCaching {
	return nil
}
