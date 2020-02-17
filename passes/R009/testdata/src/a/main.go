package a

import (
	"log"
)

func f() {
	// Comment ignored

	//lintignore:R009
	panic("test")

	panic("test") //lintignore:R009

	//lintignore:R009
	log.Panic("test")

	log.Panic("test") //lintignore:R009

	//lintignore:R009
	log.Panicf("test")

	log.Panicf("test") //lintignore:R009

	//lintignore:R009
	log.Panicln("test")

	log.Panicln("test") //lintignore:R009

	// Failing

	panic("test") // want "avoid panic\\(\\) usage"

	log.Panic("test") // want "avoid log.Panic\\(\\) usage"

	log.Panicf("%s", "test") // want "avoid log.Panicf\\(\\) usage"

	log.Panicln("test") // want "avoid log.Panicln\\(\\) usage"
}
