package log

import (
	// _ "github.com/wangfmD/rvs/log"
	"errors"
	"log"
	"testing"
)

// TestT1 ...
func TestT1(t *testing.T) {
	log.Println("ddd")
}

func ExampleNew() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		log.Println(err)
	}
}
