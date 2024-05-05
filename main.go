package main

import (
	"Distributed-cloud-storage/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	http.HandleFunc("/user/signup", handler.SignUpHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/info", handler.UserInfoHandler)
	http.HandleFunc("/file/fastupload", handler.TryFaseUploadHandler)
	// 分块上传接口
	http.HandleFunc("/file/mpupload/init", handler.InitialMutilpartUploadHandler)
	http.HandleFunc("/file/mpupload/uppart", handler.UploadHandler)
	http.HandleFunc("/file/mpupload/complete", handler.CompleteUploadHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err : %s", err)
	}
}
