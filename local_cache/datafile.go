package local_cache

type DataFiles struct {
	metaInfo  *MetaInfo
	datafiles []string
}

type MetaInfo struct {
	verStart Version
	verEnd   Version
	path     string
	size     Size
}

type DataFile struct {
	version Version
	size    Size
	path    string // <big ver>.<small ver>.data
	entries []*Entry
}
