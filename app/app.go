package app

import "github.com/l3x/sizeof-struct/internal/log"

var appLog log.Logger

const DefaultHttpPort = ":7777"

// Represents simple zero-cost message that can be used
// as signal between goroutines.
type sig struct{}
