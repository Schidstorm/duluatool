package constants

type Constants struct {
	SlotDirectoryName                  string
	SlotMetaFileName                   string
	HandlerDirectoryName               string
	WindowsInvalidFilenameCharacters   []string
	HandlerCodeFileName                string
	HandlerMetaFileName                string
	WindowsInvalidCharacterReplacement string
}

var Current = Constants{
	SlotDirectoryName:    "slot",
	SlotMetaFileName:     "meta.json",
	HandlerDirectoryName: "handler",
	WindowsInvalidFilenameCharacters: []string{
		">", "<", ":", "\"", "/", "\\", "|", "?", "*",
	},
	HandlerCodeFileName:                "code.lua",
	HandlerMetaFileName:                "meta.json",
	WindowsInvalidCharacterReplacement: "_",
}
