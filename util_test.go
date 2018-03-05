package golog

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hiro4bbh/go-assert"
)

func TestIsTerminal(t *testing.T) {
	var buf bytes.Buffer
	goassert.New(t, false).Equal(IsTerminal(&buf))
	f := goassert.New(t).SucceedNew(ioutil.TempFile("", "is_terminal")).(*os.File)
	goassert.New(t, false).Equal(IsTerminal(f))
	// We cannot test the following case?
	//goassert.New(t, true).Equal(IsTerminal(os.Stdout))
}
