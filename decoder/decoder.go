package decoder

import (
	"encoding/json"
	"github.com/schidstorm/duluatool/structure"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
)

func Run(options *Options) error {
	buffer, err := ioutil.ReadFile(options.InputFilePath)
	if err != nil {
		return err
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
		slotDir := path.Join(options.OutputDirectory, "slot", slot.Name)
		handlerDir := path.Join(slotDir, "handler")
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
	slotDir := path.Join(options.OutputDirectory, "slot", slot.Name)
	writeJsonMetadata(path.Join(slotDir, "meta.json"), &slotMeta)
}

func writeHandlerFiles(options *Options, encoded *structure.Encoded) {
	for i, handler := range encoded.Handlers {
		slot, ok := encoded.Slots[handler.Filter.SlotKey]
		if !ok {
			logrus.Warningf("Slot key %s defined in handler with index %d is not defined", handler.Filter.SlotKey, i)
			continue
		}

		writeCodeFile(options, slot, handler)
		writeHandlerMetaFile(options, slot, handler)
	}
}

func writeCodeFile(options *Options, slot structure.Slot, handler structure.Handler) {
	slotDir := path.Join(options.OutputDirectory, "slot", slot.Name)
	handlerDir := path.Join(slotDir, "handler")
	err := ioutil.WriteFile(path.Join(handlerDir, handler.Key+".lua"), []byte(handler.Code), os.ModePerm)
	if err != nil {
		logrus.Warn(err)
	}
}

func writeHandlerMetaFile(options *Options, slot structure.Slot, handler structure.Handler) {
	slotDir := path.Join(options.OutputDirectory, "slot", slot.Name)
	handlerDir := path.Join(slotDir, "handler")
	writeJsonMetadata(path.Join(handlerDir, handler.Key+".json"), &handler)
}

func writeJsonMetadata(path string, data interface{}) {
	handlerMetadata, err := json.Marshal(data)
	err = ioutil.WriteFile(path, handlerMetadata, os.ModePerm)
	if err != nil {
		logrus.Warn(err)
	}
}
