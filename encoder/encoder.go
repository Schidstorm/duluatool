package encoder

import (
	"encoding/json"
	"github.com/atotto/clipboard"
	"github.com/schidstorm/duluatool/constants"
	"github.com/schidstorm/duluatool/structure"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"sort"
)

func Run(options *Options) error {
	encoded, err := generateEncoded(options.InputDirectory)
	if err != nil {
		return err
	}

	buffer, err := json.Marshal(encoded)
	if err != nil {
		return err
	}

	if options.OutputClipboard {
		err = clipboard.WriteAll(string(buffer))
		if err != nil {
			return err
		}
	} else {
		err = ioutil.WriteFile(options.OutputFilePath, buffer, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateEncoded(directory string) (*structure.Encoded, error) {
	slotsDir := path.Join(directory, constants.Current.SlotDirectoryName)
	slotsFileInfo, err := ioutil.ReadDir(slotsDir)
	if err != nil {
		return nil, err
	}

	encoded := &structure.Encoded{
		Slots:    map[string]structure.Slot{},
		Handlers: []structure.Handler{},
		Events:   []interface{}{},
		Methods:  []interface{}{},
	}
	for _, slotFileInfo := range slotsFileInfo {
		if slotFileInfo.IsDir() {
			slotDir := path.Join(slotsDir, slotFileInfo.Name())
			slotMeta, err := readSlotMeta(slotDir)
			if err != nil {
				logrus.Warn(err)
				continue
			}

			encoded.Slots[slotMeta.SlotKey] = slotMeta.Slot
			encoded.Handlers = append(encoded.Handlers, readSlotHandlers(path.Join(slotDir, constants.Current.HandlerDirectoryName))...)
		}
	}

	sort.Sort(handlerSortInterface{Handlers: encoded.Handlers})

	return encoded, nil
}

func readSlotMeta(directory string) (structure.SlotMeta, error) {
	slotMeta := structure.SlotMeta{}
	slotMetadata, err := ioutil.ReadFile(path.Join(directory, constants.Current.SlotMetaFileName))
	if err != nil {
		return slotMeta, err
	}

	err = json.Unmarshal(slotMetadata, &slotMeta)
	if err != nil {
		return slotMeta, err
	}

	return slotMeta, nil
}

func readSlotHandlers(directory string) []structure.Handler {
	var handlers []structure.Handler
	handlersFileInfo, err := ioutil.ReadDir(directory)
	if err != nil {
		logrus.Warn(err)
		return handlers
	}

	for _, handlerDirInfo := range handlersFileInfo {
		if handlerDirInfo.IsDir() {
			handler, err := readHandler(path.Join(directory, handlerDirInfo.Name()))
			if err != nil {
				logrus.Warn(err)
			} else {
				handlers = append(handlers, handler)
			}
		}
	}

	return handlers
}

func readHandler(handlerDir string) (structure.Handler, error) {
	handler := structure.Handler{}
	codePath := path.Join(handlerDir, constants.Current.HandlerCodeFileName)
	metaPath := path.Join(handlerDir, constants.Current.HandlerMetaFileName)

	metadataBuffer, err := ioutil.ReadFile(metaPath)
	if err != nil {
		return handler, err
	}

	err = json.Unmarshal(metadataBuffer, &handler)
	if err != nil {
		return handler, err
	}

	codeBuffer, err := ioutil.ReadFile(codePath)
	if err != nil {
		logrus.Warnf("handler %s doesnt have a code file %s", handlerDir, codePath)
		return handler, nil
	} else {
		handler.Code = string(codeBuffer)
		return handler, nil
	}
}
