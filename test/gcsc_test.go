package test

import (
	"fmt"
	"github.com/gogf/gf/os/gfile"
	"testing"
)

func TestTmpPath(t *testing.T) {
	fmt.Println(gfile.TempDir() + gfile.Separator + "configClient")
}
