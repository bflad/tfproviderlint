package a

import (
	"time"
)

func fcommentignore() {
	time.Sleep(1) //lintignore:R018

	//lintignore:R018
	time.Sleep(1)
}
