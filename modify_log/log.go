package modify_log

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
