package local_cache

import "time"

type Config struct {
	confFile        string
	caching         Caching
	keyCodec        KVCodec
	valueCodec      KVCodec
	entryCodec      EntryCodec
	monitor         Monitor
	dataFileReader  DataFileReader
	dataFileWriter  DataFileWriter
	mergeDuration   time.Duration
	mergeFileNum    int
	fileIO          FileIO
	skipBrokenEntry bool
}

func DefaultConfig() *Config {
	return &Config{
		caching: NewBigCaching(),
	}
}

// Option is a function that takes a config struct and modifies it
type Option func(cfg *Config) error

func WithCache(caching Caching) Option {
	return func(cfg *Config) error {
		cfg.caching = caching
		return nil
	}
}

func WithKeyCodec(keyCodec KVCodec) Option {
	return func(cfg *Config) error {
		cfg.keyCodec = keyCodec
		return nil
	}
}

func WithValueCodec(valueCodec KVCodec) Option {
	return func(cfg *Config) error {
		cfg.valueCodec = valueCodec
		return nil
	}
}

func WithEntryCodec(entryCodec EntryCodec) Option {
	return func(cfg *Config) error {
		cfg.entryCodec = entryCodec
		return nil
	}
}

func WithMergeDuration(mergeDuration time.Duration) Option {
	return func(cfg *Config) error {
		cfg.mergeDuration = mergeDuration
		return nil
	}
}

func WithMonitor(monitor Monitor) Option {
	return func(cfg *Config) error {
		cfg.monitor = monitor
		return nil
	}
}

func WithDataFileReader(dataFileReader DataFileReader) Option {
	return func(cfg *Config) error {
		cfg.dataFileReader = dataFileReader
		return nil
	}
}

func WithDataFileWriter(dataFileWriter DataFileWriter) Option {
	return func(cfg *Config) error {
		cfg.dataFileWriter = dataFileWriter
		return nil
	}
}

func WithFileIO(fileIO FileIO) Option {
	return func(cfg *Config) error {
		cfg.fileIO = fileIO
		return nil
	}
}

func WithSkipBrokenEntry(skipBrokenEntry bool) Option {
	return func(cfg *Config) error {
		cfg.skipBrokenEntry = skipBrokenEntry
		return nil
	}
}
