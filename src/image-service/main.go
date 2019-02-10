package main

import "log"

func main() {
	err:=CheckConf()
	if err!=nil{
		log.Fatalf("env conf check failed:%v\n",err)
	}

	StartServer()
}
