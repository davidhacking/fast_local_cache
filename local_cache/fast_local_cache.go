package local_cache

import (
	"fmt"
	"github.com/spf13/cast"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type FastLocalCache struct {
	sync.RWMutex
	path string
	cfg  *Config
}

func (f *FastLocalCache) SyncDB(remoteVersion Version) (deltaData []byte, err error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastLocalCache) Get(key interface{}) (value interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastLocalCache) Put(key interface{}, value interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (f *FastLocalCache) Del(key interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (f *FastLocalCache) BatchGet(keys []interface{}) (values map[interface{}]interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (f *FastLocalCache) BatchPut(values map[interface{}]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (f *FastLocalCache) BatchDel(keys []interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (f *FastLocalCache) Merge() error {
	//TODO implement me
	panic("implement me")
}

func (f *FastLocalCache) Init(path string, opts ...Option) error {
	f.path = path
	f.cfg = DefaultConfig()
	for _, opt := range opts {
		err := opt(f.cfg)
		if err != nil {
			return fmt.Errorf("init config failed, err=%v", err)
		}
	}
	c := f.cfg
	datafiles, err := listDataFiles(f.path)
	if err != nil {
		return fmt.Errorf("listDataFiles failed, err=%v", err)
	}
	sort.SliceStable(datafiles, func(i, j int) bool {
		aa := strings.Split(datafiles[i], ".")
		if len(aa) != 3 {
			panic(fmt.Errorf("datafile name format invalid, name=%v", datafiles[i]))
		}
		bb := strings.Split(datafiles[j], ".")
		if len(bb) != 3 {
			panic(fmt.Errorf("datafile name format invalid, name=%v", datafiles[j]))
		}
		a0, a1 := cast.ToInt(aa[0]), cast.ToInt(aa[1])
		b0, b1 := cast.ToInt(bb[0]), cast.ToInt(bb[1])
		if a0 != b0 {
			return a0 < b0
		}
		return a1 < b1
	})
	cache := f.cfg.caching
	for _, datafile := range datafiles {
		df, err := readDataFile(datafile, c.fileIO, c.dataFileReader)
		if err != nil {
			return fmt.Errorf("readDataFile failed, err=%v", err)
		}
		err = writeDataFileToCache(df, cache, c.skipBrokenEntry, c.monitor, c.entryCodec)
		if err != nil {
			return fmt.Errorf("writeDataFileToCache failed, err=%v", err)
		}
	}
	return nil
}

func writeDataFileToCache(datafile *DataFile, cache Caching, skipBrokenEntry bool, monitor Monitor,
	entryCodec EntryCodec) error {
	for i, entry := range datafile.entries {
		if entry.IsIntegrate() {
			monitor.Inc("entry_broken")
			if !skipBrokenEntry {
				return fmt.Errorf("read broken entry, path=%v, i=%v", datafile.path, i)
			}
			continue
		}
		err := putEntryToCache(entry, cache, entryCodec, monitor)
		if err != nil {
			return err
		}
	}
	return nil
}

func putEntryToCache(entry *Entry, cache Caching, entryCodec EntryCodec, monitor Monitor) error {
	value, err := cache.Get(entry.key)
	if err != nil {
		return err
	}
	if len(value) > 0 {
		old, err := entryCodec.Decode(value)
		if err != nil {
			return err
		}
		if entry.ts < old.ts {
			monitor.Inc("entry_ver_small")
			return nil
		}
	}
	if entry.delFlag {
		err := cache.Del(entry.key)
		if err != nil {
			return err
		}
		return nil
	}
	bs, err := entryCodec.Encode(entry)
	if err != nil {
		return err
	}
	err = cache.Put(entry.key, bs)
	if err != nil {
		return err
	}
	return nil
}

func listDataFiles(path string) ([]string, error) {
	globPattern := filepath.Join(path, datafileMatch)
	matches, err := filepath.Glob(globPattern)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func readDataFile(path string, fio FileIO, dataFileReader DataFileReader) (*DataFile, error) {
	reader, err := fio.BuildReader(path)
	if err != nil {
		return nil, err
	}
	return dataFileReader.ReadDataFile(reader)
}

func New() FastLocalCaching {
	return &FastLocalCache{}
}
