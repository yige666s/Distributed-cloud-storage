package main

import (
	"Distributed-cloud-storage/config"
	"Distributed-cloud-storage/db"
	"Distributed-cloud-storage/mq"
	"Distributed-cloud-storage/store/oss"
	"bufio"
	"encoding/json"
	"log"
	"os"
)

// 处理文件转移的真正逻辑
func ProcessTransfer(msg []byte) bool {
	// 1.解析msg
	pubData := mq.TansferData{}
	err := json.Unmarshal(msg, pubData)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	// 2.根据临时存储文件路径，创建文件句柄
	fd, err := os.Open(pubData.CurLocation)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	// 3.通过文件句柄将文件读出并上传到oss
	// options := []aoss.Option{ // 保证下载时为原文件名
	// 	aoss.ContentDisposition("attachment;filename=\"" + fileMeta.FileName + "\"")}
	err = oss.Bucket().PutObject(pubData.DesLocation, bufio.NewReader(fd))
	if err != nil {
		log.Println(err.Error())
		return false
	}

	// 4.更新文件表
	suc := db.UpdateFileLocation(
		pubData.FileHash,
		pubData.DesLocation,
	)
	if !suc {
		log.Println(err.Error())
		return false
	}

	return true
}
func main() {
	log.Println("开始监听转移任务队列")
	mq.StartConsume(config.TransOSSQueueName, "transfer_oss", ProcessTransfer)
}
