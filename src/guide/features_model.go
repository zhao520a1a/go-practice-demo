package main

import (
	"log"
)

type options struct {
	cache  bool
	logger *log.Logger
}

type Option interface {
	apply(*options)
}

type cacheOption bool

func (c cacheOption) apply(opts *options) {
	opts.cache = bool(c)
}

func WithCache(c bool) Option {
	return cacheOption(c)
}

type loggerOption struct {
	Log *log.Logger
}

func (l loggerOption) apply(opts *options) {
	opts.logger = l.Log
}

func WithLogger(log *log.Logger) Option {
	return loggerOption{Log: log}
}


func Open(
	opts ...Option,
)  {
	options := options{
		cache: true ,
		logger: &log.Logger{},
	}

	for _, o := range opts {
		o.apply(&options)
	}




}