package mq

import "Distributed-cloud-storage/common"

// 转移队列中消息载体的结构格式
type TansferData struct {
	FileHash      string
	CurLocation   string
	DesLocation   string
	DestStoreType common.StoreType
}
