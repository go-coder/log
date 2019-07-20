package stderr

import (
	"os"
	"sync"

	"github.com/go-coder/log/pkg/api"
	"github.com/go-coder/log/pkg/impl/common"
)

// New returns an api.EntryWriter, backendend by os.Stderr
func New() api.EntryWriter {
	return &mutexed{}
}

// mutexed is an goroutine-safe implementation of api.EntryWriter, by mutex
type mutexed struct {
	mu sync.Mutex
}

var _ api.EntryWriter = (*mutexed)(nil)

func (o *mutexed) WriteEntry(e *api.Entry) {
	o.mu.Lock()
	common.NewEntryWriteFunc(os.Stderr)(e)
	o.mu.Unlock()
}

// func Stderr() api.EntryWriter {
// 	out := &outter{
// 		entryChan: make(chan *api.Entry),
// 		writeFunc: common.NewEntryWriteFunc(os.Stderr),
// 	}
// 	go func() {
// 		for e := range out.entryChan {
// 			out.writeFunc(e)
// 		}
// 	}()
// 	return out
// }

// type outter struct {
// 	entryChan chan *api.Entry
// 	writeFunc func(*api.Entry)
// }

// var _ api.EntryWriter = (*outter)(nil)

// func (o *outter) WriteEntry(e *api.Entry) {
// 	o.entryChan <- e
// }
