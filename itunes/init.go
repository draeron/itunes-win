// +build windows

package itunes

import (
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// Init initializes COM library and returns ITunes reference. This must be called
// at the beginning.
func Init() (*ITunes, error) {
	if err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED); err != nil {
		return nil, err
	}
	obj, err := oleutil.CreateObject("iTunes.Application")
	if err != nil {
		ole.CoUninitialize()
		return nil, err
	}
	iTunes, err := obj.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		ole.CoUninitialize()
		return nil, err
	}
	return &ITunes{
		COM: COM{obj: iTunes},
	}, nil
}

// Exit does uninitialization of COM library. This must be called at the end.
func (i *ITunes) Exit() {
	ole.CoUninitialize()
}
