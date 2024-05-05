package ceph

import (
	"gopkg.in/amz.v3/aws"
	"gopkg.in/amz.v3/s3"
)

var cephConn *s3.S3

func GetCephConnection() *s3.S3 {
	if cephConn != nil {
		return cephConn
	}
	// 1. 初始化ceph的一些信息
	auth := aws.Auth{
		AccessKey: "", // TODO 搭好集群后设置
		SecretKey: "",
	}

	curRegion := aws.Region{
		Name:                 "default",
		EC2Endpoint:          "http://127.0.0.1:9080",
		S3Endpoint:           "http://127.0.0.1:9080",
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
		// Sign:aws.SignV2(), // 这个字段不用了
	}
	// 2. 创建S3类型连接
	return s3.New(auth, curRegion)
}

// 获取指定的Bucket对象
func GetCephBucket(bucket string) (*s3.Bucket, error) {
	conn := GetCephConnection()
	return conn.Bucket(bucket)
}
