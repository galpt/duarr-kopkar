package main

import (
	"crypto/tls"
	"runtime"
	"time"
)

const (
	Gigabyte      = 1 << 30
	Megabyte      = 1 << 20
	Kilobyte      = 1 << 10
	timeoutTr     = 2 * time.Hour
	memCacheLimit = 300 << 20 // 300 MB
	hostPortGin   = "7777"
	usrAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"
)

var (
	mem         runtime.MemStats
	HeapAlloc   string
	SysMem      string
	Frees       string
	NumGCMem    string
	timeElapsed string
	latestLog   string

	// CertFilePath = "/etc/letsencrypt/live/net.0ms.dev/fullchain.pem"
	// KeyFilePath  = "/etc/letsencrypt/live/net.0ms.dev/privkey.pem"

	tlsConf = &tls.Config{
		InsecureSkipVerify: true,
	}
)
