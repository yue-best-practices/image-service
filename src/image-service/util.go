package main

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"github.com/satori/go.uuid"
	"strings"
)

func CheckExistAndCreate(dir string)error{
	_,err:=os.Stat(dir)
	if err!=nil{
		if os.IsNotExist(err){
			err:=os.MkdirAll(dir,os.ModePerm)
			if err!=nil{
				return err
			}
			return nil
		}
		return err
	}
	return nil
}


func CheckFileIsExist(fileName string)bool{
	_,err:=os.Stat(fileName)
	if err!=nil{
		if os.IsNotExist(err){
			return false
		}
		fmt.Printf("File check exist error:%v\n",err)
		return false
	}
	return true
}


func GenerateImagePath()string{
	return time.Now().Format("2006/01/02")
}

func GenerateImageName()string{
	return strings.Replace(fmt.Sprintf("%s",uuid.Must(uuid.NewV4())),"-","",-1)
}

func CheckIsImage(name string)bool{
	str:=strings.Split(name,".")
	allowedExtension:=[]string{"jpg","png","gif","jpeg"}
	if len(str) >=2 && ArrayIsIncludeItem(allowedExtension,strings.ToLower(str[len(str)-1])){
		return true
	}
	return false
}

func ArrayIsIncludeItem(array []string,item string)bool{
	for _,v:=range array{
		if v==item{
			return true
		}
	}

	return false
}

func WriteHttpResponse(w http.ResponseWriter,code int,msg string,data interface{},errors error){
	v:=make(map[string]interface{})
	v["code"]=code
	v["msg"]=msg
	if data!=nil{
		v["data"]=data
	}

	if errors!=nil{
		v["errors"]=errors.Error()
	}
	b,_:=json.Marshal(v)
	w.Header().Set("Content-Type","application/json")
	w.Write(b)
}