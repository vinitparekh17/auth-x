package server

import (
	"fmt"
	"os"
)

func Init() {
	r := NewRouter()
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
