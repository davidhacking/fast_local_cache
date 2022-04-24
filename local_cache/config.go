package local_cache

import "time"

type Config struct {
	confFile      string
	caching       Caching
	keyCodec      KVCodec
	valueCodec    KVCodec
	entryCodec    EntryCodec
	mergeDuration time.Duration
	mergeFileNum  int
}

// Option is a function that takes a config struct and modifies it
type Option func(*Config) error

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

func WithMergeFileNum(mergeFileNum int) Option {
	return func(cfg *Config) error {
		cfg.mergeFileNum = mergeFileNum
		return nil
	}
}
