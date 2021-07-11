package gphoto

// Binding sourced from: https://github.com/aqiank/go-gphoto2

// #cgo linux CFLAGS: -I/opt/shot-capture/include
// #include <gphoto2/gphoto2.h>
// #include <string.h>
import "C"

type CameraList struct {
	Ref   *C.CameraList
	Count int
}

func NewList() (*CameraList, error) {
	list := &CameraList{
		Count: 0,
	}
	ret := C.gp_list_new(&list.Ref)
	if ret != PORT_RESULT_OK {
		return nil, AsPortResult(ret).Error()
	}
	list.Count = int(ret)
	return list, nil
}

func (list *CameraList) Reset() {
	C.gp_list_reset(list.c())
}

func (list *CameraList) Free() {
	C.gp_list_free(list.c())
}

func (list *CameraList) c() *C.CameraList {
	return (*C.CameraList)(list.Ref)
}

func (list *CameraList) Name(index int) string {
	var name *C.char
	C.gp_list_get_name(list.c(), C.int(index), &name)
	return C.GoString(name)
}

func (list *CameraList) Value(index int) string {
	var value *C.char
	C.gp_list_get_name(list.c(), C.int(index), &value)
	return C.GoString(value)
}