package a

import (
	"time"
)

func f() {
	time.Sleep(1) // want "prefer resource.Retry\\(\\) or \\(resource\\.StateChangeConf\\)\\.WaitForState\\(\\)"
}
