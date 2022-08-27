#### gpc-function-logger
is a very simple Go library for proper logging in GCP Cloud Functions by severity levels.

Available methods (without/with message formatting):
- `Debug()/Debugf()`
- `Info()/Infof()`
- `Notice()/Noticef()`
- `Warn()/Warnf()`
- `Error()/Errorf()`
- `Critical()/Criticalf()`
- `Alert()/Alertf()`


Usage example:
```go
package main

import (
	"os"
	"runtime"
	
	logger "github.com/hotafrika/gcp-function-logger"
)

func main() {
	log := logger.NewLogger(os.Stdout)

	// Info level
	log.Info("Hello Cloud!")

	// Info level with formatting
	log.Infof("My architecture is %s", runtime.GOARCH)
}
```
