package constants

import "os"

/*
Absolute path of the 'service-key.json' file used to initiate Firebase SDK.
*/
var ServiceKeyPath = os.Getenv("PROJECT_ROOT") + "/service-key.json"
