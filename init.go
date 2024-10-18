package tools

import (
	"os"
	"path"
)

func SetWorkSpace() (workspace string, err error) {
	workspace = path.Dir(os.Args[0])
	err = os.Chdir(workspace)
	return
}
