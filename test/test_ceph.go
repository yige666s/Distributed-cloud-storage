package main

import (
	"Distributed-cloud-storage/store/ceph"
	"fmt"

	"gopkg.in/amz.v3/s3"
)

func main() {
	bucket, _ := ceph.GetCephBucket("testbucket1")
	// 创建一个新的bucket
	err := bucket.PutBucket(s3.PublicRead)
	fmt.Printf("create bucket err :%v\n", err)
	//查询这个vucket下面指定的object keys
	res, _ := bucket.List("", "", "", 100)
	fmt.Printf("objects keys %+v\n", res)
	// 新上传一个对象
	err = bucket.Put("/testupload/a.txt", []byte("just for test"), "octet-stream", s3.PublicRead)
	fmt.Printf("objects keys %+v\n", err)
	//查询这个vucket下面指定的object keys
	res, _ = bucket.List("", "", "", 100)
	fmt.Printf("objects keys %+v\n", res)
}
