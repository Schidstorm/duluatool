package decoder

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestRun(t *testing.T) {
	options := &Options{
		InputFilePath:   "../test/resources/first.json",
		OutputDirectory: path.Join(os.TempDir(), "project"),
	}
	_ = os.RemoveAll(options.OutputDirectory)

	logrus.Info("decode options ", options)
	err := Run(options)
	assert.NoError(t, err)

	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/library/meta.json"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/slot6/meta.json"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/system/meta.json"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/unit/meta.json"))

	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/slot6/handler/0.json"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/slot6/handler/0.lua"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/system/handler/3.json"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/system/handler/3.lua"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/system/handler/4.json"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/system/handler/4.lua"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/system/handler/5.json"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/system/handler/5.lua"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/unit/handler/1.json"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/unit/handler/1.lua"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/unit/handler/2.json"))
	assert.FileExists(t, path.Join(options.OutputDirectory, "slot/unit/handler/2.lua"))
}
