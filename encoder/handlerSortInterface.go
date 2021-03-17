package encoder

import (
	"github.com/schidstorm/duluatool/structure"
	"strconv"
)

type handlerSortInterface struct {
	Handlers []structure.Handler
}

func (h handlerSortInterface) Len() int {
	return len(h.Handlers)
}

func (h handlerSortInterface) Less(i, j int) bool {
	a := h.Handlers[i]
	b := h.Handlers[j]
	aInt, _ := strconv.ParseInt(a.Key, 10, 32)
	bInt, _ := strconv.ParseInt(b.Key, 10, 32)
	return aInt < bInt
}

func (h handlerSortInterface) Swap(i, j int) {
	tmp := h.Handlers[i]
	h.Handlers[i] = h.Handlers[j]
	h.Handlers[j] = tmp
}
