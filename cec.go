package cec 

/*
#cgo pkg-config: libcec
#include <stdlib.h> 
#include <stdio.h>
#include <libcec/cecc.h>

void setName(libcec_configuration conf, char *name) {
	snprintf(conf.strDeviceName, 13, "%s", name);
}

*/
import "C"
import "errors"

type CECConfiguration struct {
	DeviceName string
}

type CECAdapter struct {
	Path string
	Comm string
}

func Init(config CECConfiguration) error {

	var conf C.libcec_configuration
	
	C.setName(conf, C.CString(config.DeviceName))

	result := C.cec_initialise(&conf)
	if result < 1 {
		return errors.New("Failed to init CEC")
	}
	return nil
}

func GetFirstAdapter() (CECAdapter, error) {

	var adapter CECAdapter

	var deviceList [1]C.cec_adapter 
	devicesFound := C.cec_find_adapters(&deviceList[0], 1, nil)

	if devicesFound < 1 {
		return adapter, errors.New("No Device Found")
	}

	device := deviceList[0]
	adapter.Path = C.GoStringN(&device.path[0], 1024)
	adapter.Comm = C.GoStringN(&device.comm[0], 1024)

	return adapter, nil
}

func PingAdapters() bool {
	return C.cec_ping_adapters() == 1
}

func VolumeUp() {
	C.cec_volume_up(1)
}

func VolumeDown() {
	C.cec_volume_down(1)
}

func MuteAudio() {
	C.cec_mute_audio(1)
}

func InitVideoStandalone() {
	C.cec_init_video_standalone()
}
