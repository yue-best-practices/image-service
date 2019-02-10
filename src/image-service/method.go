package main

import (
	"net/http"
	"io"
	"os"
	"fmt"
	"strings"
	"encoding/base64"
	"github.com/gernest/alien"
	"io/ioutil"
)

func upload(w http.ResponseWriter,r *http.Request){
	reader,err:=r.MultipartReader()
	if err!=nil{
		WriteHttpResponse(w,400,"Request Error",nil,err)
		return
	}

	imageEncode:=""

	for {
		part,err:=reader.NextPart()
		if err==io.EOF{
			break
		}
		if part.FormName()=="image" && part.FileName()!="" && CheckIsImage(part.FileName()){
			imagePath:=GenerateImagePath()
			err:=CheckExistAndCreate(fmt.Sprintf("%s/%s",IMAGE_PATH,imagePath))
			if err!=nil{
				fmt.Printf("create dir failed:%v\n",err)
				break
			}
			strs:=strings.Split(part.FileName(),".")
			imageName:=fmt.Sprintf("%s.%s",GenerateImageName(),strs[len(strs)-1])
			full:=fmt.Sprintf("%s/%s/%s",IMAGE_PATH,imagePath,imageName)
			dst,err:=os.Create(full)

			if err!=nil{
				fmt.Printf("create tmp file error:%v\n",err)
			}
			defer dst.Close()
			io.Copy(dst,part)
			imageEncode=base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s/%s",imagePath,imageName)))
		}
	}

	if imageEncode==""{
		WriteHttpResponse(w,400,"Invalid Request(image is missing)",nil,nil)
		return
	}
	WriteHttpResponse(w,200,"Success",map[string]string{
		"image":imageEncode,
	},nil)
}


func image(w http.ResponseWriter, r *http.Request){
	p:=alien.GetParams(r)
	imgEncode:=p.Get("imgEncode")
	imgPath,err:=base64.StdEncoding.DecodeString(imgEncode)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sts:=strings.Split(string(imgPath),".")
	extension:=strings.ToLower(sts[len(sts)-1])
	file,err:=ioutil.ReadFile(fmt.Sprintf("%s/%s",IMAGE_PATH,imgPath))
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type",fmt.Sprintf("image/%s",extension))
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}