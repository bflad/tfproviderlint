package a

import (
	t "time"
)

func falias() {
	t.Sleep(1) // want "prefer resource.Retry\\(\\) or \\(resource\\.StateChangeConf\\)\\.WaitForState\\(\\)"
}
