package folder1

import "runtime"

func Version() string {
	return runtime.Version()
}
