package encoder

import (
	"encoding/json"
	"github.com/schidstorm/duluatool/structure"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
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

	err = ioutil.WriteFile(options.OutputFilePath, buffer, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func generateEncoded(directory string) (*structure.Encoded, error) {
	slotsDir := path.Join(directory, "slot")
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
			encoded.Handlers = append(encoded.Handlers, readSlotHandlers(path.Join(slotDir, "handler"))...)
		}
	}

	sort.Sort(handlerSortInterface{Handlers: encoded.Handlers})

	return encoded, nil
}

func readSlotMeta(directory string) (structure.SlotMeta, error) {
	slotMeta := structure.SlotMeta{}
	slotMetadata, err := ioutil.ReadFile(path.Join(directory, "meta.json"))
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

	for _, handlerFileInfo := range handlersFileInfo {
		if !handlerFileInfo.IsDir() && strings.HasSuffix(handlerFileInfo.Name(), ".json") {
			handler, err := readHandler(path.Join(directory, handlerFileInfo.Name()))
			if err != nil {
				logrus.Warn(err)
			} else {
				handlers = append(handlers, handler)
			}
		}
	}

	return handlers
}

func readHandler(metadataFilePath string) (structure.Handler, error) {
	handler := structure.Handler{}
	codePath := path.Join(path.Dir(metadataFilePath), strings.TrimSuffix(path.Base(metadataFilePath), ".json")+".lua")

	metadataBuffer, err := ioutil.ReadFile(metadataFilePath)
	if err != nil {
		return handler, err
	}

	err = json.Unmarshal(metadataBuffer, &handler)
	if err != nil {
		return handler, err
	}

	codeBuffer, err := ioutil.ReadFile(codePath)
	if err != nil {
		logrus.Warnf("handler %s doesnt have a code file %s", metadataFilePath, codePath)
		return handler, nil
	} else {
		handler.Code = string(codeBuffer)
		return handler, nil
	}
}
