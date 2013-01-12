#include <stdlib.h>
#include <stdio.h>
#include <libcec/cecc.h>

ICECCallbacks        g_callbacks;
libcec_configuration g_config;

int CecLogMessage(void *UNUSED, const cec_log_message message)
{

	printf("%s\n", message.message);

	return 0;
}

main()
{
	snprintf(g_config.strDeviceName, 13, "CECTester");
	g_config.clientVersion       = CEC_DEFAULT_SETTING_CEC_VERSION;
	g_config.bActivateSource     = 4;
	g_callbacks.CBCecLogMessage = &CecLogMessage;
	g_config.callbacks           = &g_callbacks;
	
	int result = cec_initialise(&g_config);
	if (result < 1) {
		printf("Failed to init CEC\n");
		return;
	}

	cec_adapter *deviceList = malloc(sizeof(cec_adapter));	
	int8_t iDevicesFound = cec_find_adapters(deviceList, 1, NULL);

	if (iDevicesFound <= 0) {
		printf("No devices found!\n");
		return;
	}

	int opened = cec_open((*deviceList).comm, 1000);
	if (opened < 1) {
		printf("Failed to open device\n");
		return;
	}

	

	return 0;
}
