package local_cache

const (
	datafileSuffix = ".data"
	datafileMatch  = "*.data"
)

type Version uint64

var (
	KB              Size = 2 << 10
	MB                   = KB * KB
	GB                   = MB * KB
	DataFileMaxSize      = 4 * MB
)

type Size uint64
