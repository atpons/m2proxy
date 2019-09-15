package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"

	"github.com/atpons/m2proxy/pkg/util"

	"github.com/atpons/m2proxy/pkg/server"
	"github.com/atpons/m2proxy/pkg/storage"
)

func init() {
	if v, _ := strconv.Atoi(os.Getenv("DEBUG")); v > 0 {
		util.Debug = v
	}
}

func main() {
	if util.Debug > 0 || os.Getenv("PPROF") == "1" {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}
	st := storage.NewLocalStorage("sample.db")
	s := server.NewServer(":11211", &st)
	s.ListenAndServe()
}
