package applib

import (
	"os"
	"path/filepath"
)

// ProName project name
var ProName string

func init() {
	// 默认可执行文件名作为文件名，执行go build .
	ProName = filepath.Base(os.Args[0])
}

// SetProName set pro name
func SetProName(proName string) {
	ProName = proName
}
