package version

import (
	"fmt"
)

var (
	hash           string
	buildtimestamp string
)

// Version returns the build information like git hash and compile date as string
func Version() string {
	return fmt.Sprintf("Version: %s from %s", hash, buildtimestamp)
}
