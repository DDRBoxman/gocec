package cec 

/*
#cgo pkg-config: libcec
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

func Init(config CECConfiguration) error {

	var conf C.libcec_configuration
	
	C.setName(conf, C.CString(config.DeviceName))

	result := C.cec_initialise(&conf)
	if result < 1 {
		return errors.New("Failed to init CEC")
	}
	return nil
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
