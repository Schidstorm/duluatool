package structure

type Encoded struct {
	Slots    map[string]Slot `json:"slots"`
	Handlers []Handler       `json:"handlers"`
	Events   []interface{}   `json:"events"`
	Methods  []interface{}   `json:"methods"`
}

type SlotMeta struct {
	SlotKey string `json:"slotKey"`
	Slot    Slot   `json:"slot"`
}

type Slot struct {
	Name string   `json:"name"`
	Type SlotType `json:"type"`
}

type SlotType struct {
	Events  []interface{} `json:"events"`
	Methods []interface{} `json:"methods"`
}

type Handler struct {
	Code   string        `json:"code"`
	Filter HandlerFilter `json:"filter"`
	Key    string        `json:"key"`
}

type HandlerFilter struct {
	Args      []HandlerFilterArg `json:"args"`
	Signature string             `json:"signature"`
	SlotKey   string             `json:"slotKey"`
}

type HandlerFilterArg struct {
	Value string `json:"value"`
}
