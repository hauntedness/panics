package panics

import (
	"log"
	"runtime"
	"strconv"
)

func TODO(err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Println(file + ":" + strconv.Itoa(line) + ": " + "TODO something to be done over here")
	if err != nil {
		panic(err)
	}
}

func Panic(err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Println(file + ":" + strconv.Itoa(line) + ": " + "panic!!!")
	panic(err)
}

func Ignore(err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Println(file + ":" + strconv.Itoa(line) + ": " + "an error is ignored over here")
	if err != nil {
		log.Println(err)
	}
}
