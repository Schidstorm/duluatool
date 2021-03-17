package encoder

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestRun(t *testing.T) {
	options := &Options{
		InputDirectory: "../test/resources/first",
		OutputFilePath: path.Join(os.TempDir(), "first.json"),
	}
	assert.NoError(t, os.Remove(options.OutputFilePath))

	logrus.Info("encode options ", options)
	err := Run(options)
	assert.NoError(t, err)

	assert.FileExists(t, options.OutputFilePath)
	actual, _ := ioutil.ReadFile(options.OutputFilePath)
	expected, _ := ioutil.ReadFile("../test/resources/first.json")

	assert.JSONEq(t, string(expected), string(actual))

}
