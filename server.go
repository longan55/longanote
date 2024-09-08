package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// 文章
type Article struct {
	Title   string //标题
	Content string //内容
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 解析表单数据
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, "Failed to parse form", http.StatusBadRequest)
	// 	return
	// }

	// 获取表单参数
	path := r.PostFormValue("path")

	fmt.Println("path: ", path)
	path = "./" + path

	fileData, err := os.ReadFile(path)
	if err != nil {
		log.Printf("文件\"%s\"读取失败: %v", path, err)
		return
	}
	article := Article{
		Title:   filepath.Base(path),
		Content: string(fileData),
	}
	marshal, err := json.Marshal(article)
	if err != nil {
		log.Println("json :", err)
	}

	i, err := w.Write(marshal)
	if err != nil {
		log.Println("Write ", i, " err:", err)
	}
}

var filenameCache = make(map[string]map[string]string)
