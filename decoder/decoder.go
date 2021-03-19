package decoder

import (
	"encoding/json"
	"github.com/atotto/clipboard"
	"github.com/schidstorm/duluatool/constants"
	"github.com/schidstorm/duluatool/structure"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

func Run(options *Options) error {
	var buffer []byte
	var err error
	if options.InputClipboard {
		cpString, err := clipboard.ReadAll()
		if err != nil {
			return err
		}
		buffer = []byte(cpString)
	} else {
		buffer, err = ioutil.ReadFile(options.InputFilePath)
		if err != nil {
			return err
		}
	}

	encoded, err := parseEncoded(buffer)
	prepareSlots(options, encoded)
	writeHandlerFiles(options, encoded)

	return nil
}

func parseEncoded(buffer []byte) (*structure.Encoded, error) {
	encoded := &structure.Encoded{}
	err := json.Unmarshal(buffer, encoded)
	return encoded, err
}

func prepareSlots(options *Options, encoded *structure.Encoded) {
	for slotKey, slot := range encoded.Slots {
		slotDir := path.Join(options.OutputDirectory, constants.Current.SlotDirectoryName, slot.Name)
		handlerDir := path.Join(slotDir, constants.Current.HandlerDirectoryName)
		err := os.MkdirAll(handlerDir, os.ModePerm)
		writeSlotMetaFile(options, slotKey, slot)
		if err != nil {
			logrus.Warn(err)
			continue
		}
	}
}

func writeSlotMetaFile(options *Options, slotKey string, slot structure.Slot) {
	slotMeta := structure.SlotMeta{
		SlotKey: slotKey,
		Slot:    slot,
	}
	slotDir := path.Join(options.OutputDirectory, constants.Current.SlotDirectoryName, slot.Name)
	writeJsonMetadata(path.Join(slotDir, constants.Current.SlotMetaFileName), &slotMeta)
}

func writeHandlerFiles(options *Options, encoded *structure.Encoded) {
	for i, handler := range encoded.Handlers {
		slot, ok := encoded.Slots[handler.Filter.SlotKey]
		if !ok {
			logrus.Warningf("Slot key %s, defined in handler with index %d, is not defined", handler.Filter.SlotKey, i)
			continue
		}

		handlerDir := createHandlerDir(options, slot, handler)
		writeCodeFile(handler, handlerDir)
		writeHandlerMetaFile(handler, handlerDir)
	}
}

func createHandlerDir(options *Options, slot structure.Slot, handler structure.Handler) string {
	handlerFolderName := handler.Key + "_" + handler.Filter.Signature
	if runtime.GOOS == "windows" {
		for _, c := range constants.Current.WindowsInvalidFilenameCharacters {
			handlerFolderName = strings.ReplaceAll(handlerFolderName, c, constants.Current.WindowsInvalidCharacterReplacement)
		}
	}

	slotDir := path.Join(options.OutputDirectory, constants.Current.SlotDirectoryName, slot.Name)
	handlerDir := path.Join(slotDir, constants.Current.HandlerDirectoryName, handlerFolderName)

	err := os.MkdirAll(handlerDir, os.ModePerm)
	if err != nil {
		logrus.Warn(err)
	}

	return handlerDir
}

func writeCodeFile(handler structure.Handler, handlerDir string) {
	err := ioutil.WriteFile(path.Join(handlerDir, constants.Current.HandlerCodeFileName), []byte(handler.Code), os.ModePerm)
	if err != nil {
		logrus.Warn(err)
	}
}

func writeHandlerMetaFile(handler structure.Handler, handlerDir string) {
	handler.Code = ""
	writeJsonMetadata(path.Join(handlerDir, constants.Current.HandlerMetaFileName), &handler)
}

func writeJsonMetadata(path string, data interface{}) {
	handlerMetadata, err := json.Marshal(data)
	err = ioutil.WriteFile(path, handlerMetadata, os.ModePerm)
	if err != nil {
		logrus.Warn(err)
	}
}
