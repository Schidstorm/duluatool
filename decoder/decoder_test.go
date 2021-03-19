package decoder

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
		InputFilePath:   "../test/resources/first.json",
		OutputDirectory: path.Join(os.TempDir(), "project"),
	}
	_ = os.RemoveAll(options.OutputDirectory)

	logrus.Info("decode options ", options)
	err := Run(options)
	assert.NoError(t, err)

	assertFileContent(t, path.Join(options.OutputDirectory, "slot/library/meta.json"), "{\"slotKey\":\"-3\",\"slot\":{\"name\":\"library\",\"type\":{\"events\":[],\"methods\":[]}}}")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/slot6/meta.json"), "{\"slotKey\":\"0\",\"slot\":{\"name\":\"slot6\",\"type\":{\"events\":[],\"methods\":[]}}}")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/system/meta.json"), "{\"slotKey\":\"-2\",\"slot\":{\"name\":\"system\",\"type\":{\"events\":[],\"methods\":[]}}}")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/unit/meta.json"), "{\"slotKey\":\"-1\",\"slot\":{\"name\":\"unit\",\"type\":{\"events\":[],\"methods\":[]}}}")

	assertFileContent(t, path.Join(options.OutputDirectory, "slot/slot6/handler/0_start()/meta.json"), "{\"code\":\"\",\"filter\":{\"args\":[],\"signature\":\"start()\",\"slotKey\":\"0\"},\"key\":\"0\"}")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/slot6/handler/0_start()/code.lua"), "code0")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/system/handler/3_actionStart(action)/meta.json"), "{\"code\":\"\",\"filter\":{\"args\":[{\"value\":\"option2\"}],\"signature\":\"actionStart(action)\",\"slotKey\":\"-2\"},\"key\":\"3\"}")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/system/handler/3_actionStart(action)/code.lua"), "code3")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/system/handler/4_actionStart(action)/meta.json"), "{\"code\":\"\",\"filter\":{\"args\":[{\"value\":\"option1\"}],\"signature\":\"actionStart(action)\",\"slotKey\":\"-2\"},\"key\":\"4\"}")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/system/handler/4_actionStart(action)/code.lua"), "code4")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/system/handler/5_actionStart(action)/meta.json"), "{\"code\":\"\",\"filter\":{\"args\":[{\"value\":\"option3\"}],\"signature\":\"actionStart(action)\",\"slotKey\":\"-2\"},\"key\":\"5\"}")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/system/handler/5_actionStart(action)/code.lua"), "code5")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/unit/handler/1_stop()/meta.json"), "{\"code\":\"\",\"filter\":{\"args\":[],\"signature\":\"stop()\",\"slotKey\":\"-1\"},\"key\":\"1\"}")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/unit/handler/1_stop()/code.lua"), "code1")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/unit/handler/2_tick(timerId)/meta.json"), "{\"code\":\"\",\"filter\":{\"args\":[{\"value\":\"Live\"}],\"signature\":\"tick(timerId)\",\"slotKey\":\"-1\"},\"key\":\"2\"}")
	assertFileContent(t, path.Join(options.OutputDirectory, "slot/unit/handler/2_tick(timerId)/code.lua"), "code2")

}

func assertFileContent(t *testing.T, path string, content string) {
	assert.FileExists(t, path)
	buffer, err := ioutil.ReadFile(path)
	assert.NoError(t, err)
	assert.Equal(t, content, string(buffer))
}
