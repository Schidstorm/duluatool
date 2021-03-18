package constants

type Constants struct {
	SlotDirectoryName    string
	SlotMetaFileName     string
	HandlerDirectoryName string
}

var Current = Constants{
	SlotDirectoryName:    "slot",
	SlotMetaFileName:     "meta.json",
	HandlerDirectoryName: "handler",
}
