package uploadrpchandler

import (
	"Distributed-cloud-storage/service/upload/config"
	"Distributed-cloud-storage/service/upload/proto/upload"
	"context"
)

// 用于实现UploadServicehandler的接口
type Upload struct {
}

// 用于获取上传的入口地址
func UploadEntry(ctx context.Context, resp *upload.RespEntry, req *upload.RespEntry) error {
	resp.Entry = config.UploadEntry
}
