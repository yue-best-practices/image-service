package main

import (
	"os"
	"strconv"
	"errors"
)

var (
	APP_PORT = 8030
	IMAGE_PATH = "."
)


func CheckConf()error{
	p:=os.Getenv("APP_PORT")
	if p!=""{
		l,err:=strconv.Atoi(p)
		if err!=nil{
			return err
		}
		APP_PORT=l
	}

	i:=os.Getenv("IMAGE_PATH")
	if i!=""{
		IMAGE_PATH = i
	}

	if IMAGE_PATH==""{
		return errors.New("image path not set")
	}
	return nil
}