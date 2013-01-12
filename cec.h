#include <stdlib.h>
#include <stdio.h>
#include <libcec/cecc.h>

extern void logCallback(cec_log_message);

void setName(libcec_configuration *conf, char *name);

int CecLogMessage(void *UNUSED, const cec_log_message message);

void setupCallbacks(libcec_configuration *conf);
