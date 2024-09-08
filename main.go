package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const DOCS = "./docs"

type Response struct {
	Code int
	Msg  string
	Data any
}

func (resp Response) Bytes() []byte {
	data, err := json.Marshal(resp)
	if err != nil {
		return nil
	}
	return data
}

func main() {
	http.HandleFunc("/getMenu", GetMenu)
	http.HandleFunc("/getArticle", GetArticle)
	log.Fatalln(http.ListenAndServe(":9000", nil))
}

func GetMenu(resp http.ResponseWriter, request *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	resp.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	resp.Header().Set("content-type", "application/json")             //返回数据格式是json

	infos, err := getDirectoryInfo(DOCS)
	if err != nil {
		response := Response{
			Code: 0,
			Msg:  "Error",
			Data: err,
		}
		resp.Write(response.Bytes())
		return
	}
	response := Response{
		Code: 1,
		Msg:  "OK",
		Data: infos,
	}
	resp.Write(response.Bytes())
}

// FileInfo 结构体用于存储文件/目录信息
type FileInfo struct {
	Title    string     `json:"title"`
	Path     string     `json:"path"`
	Index    int        `json:"index"`
	Type     string     `json:"type"`
	Size     int64      `json:"size,omitempty"`  // 仅在文件时有效
	Children []FileInfo `json:"child,omitempty"` // 仅在目录时有效
}

// 获取目录及其子目录信息，但不包括根目录本身
func getDirectoryInfo(path string) ([]FileInfo, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var fileInfos []FileInfo

	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())
		childInfo, err := getFileInfo(childPath)
		if err != nil {
			return nil, err
		}
		fileInfos = append(fileInfos, childInfo)
	}

	return fileInfos, nil
}

// 获取文件或目录的信息
func getFileInfo(path string) (FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, err
	}
	filename := strings.Split(filepath.Base(path), ".")
	index, _ := strconv.Atoi(filename[0])
	title := filename[1]
	fileInfo := FileInfo{
		Index: index,
		Title: title,
		Path:  path,
		Type:  "Directory",
	}
	fmt.Println("path: ", fileInfo.Path)
	if !info.IsDir() {
		fileInfo.Type = "File"
		fileInfo.Size = info.Size()
		return fileInfo, nil
	}

	fileInfo.Children = []FileInfo{}
	entries, err := os.ReadDir(path)
	if err != nil {
		return fileInfo, err
	}

	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())
		childInfo, err := getFileInfo(childPath)
		if err != nil {
			return fileInfo, err
		}
		fileInfo.Children = append(fileInfo.Children, childInfo)
	}
	return fileInfo, nil
}
