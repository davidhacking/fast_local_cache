package local_cache

type DataFile struct {
	metaInfo *MetaInfo
	entries  []*Entries
}

type MetaInfo struct {
	verStart Version
	verEnd   Version
	path     string
	size     Size
}

type Entries struct {
	version Version
	size    Size
	entries []*Entry
}
