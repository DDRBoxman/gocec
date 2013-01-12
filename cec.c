#include "_cgo_export.h"
#include "cec.h"

ICECCallbacks g_callbacks;

void setName(libcec_configuration *conf, char *name)
{
	snprintf((*conf).strDeviceName, 13, "%s", name);
}

int CecLogMessage(void *UNUSED, const cec_log_message message)
{
	printf("%s\n", message.message);
	//logCallback(message);
	return 0;
}

void setupCallbacks(libcec_configuration *conf)
{
	g_callbacks.CBCecLogMessage = &CecLogMessage;
	(*conf).callbacks = &g_callbacks;
}
