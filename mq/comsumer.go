package mq

import "log"

var done chan bool

// 开始监听队列，获取消息
func StartConsume(qName, cName string, callback func(msg []byte) bool) {
	// 1. 通过channel.Consumer获得消息信道
	msgs, err := channel.Consume(qName, cName, true, false, false, false, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// 2.循环获取队列消息
	done = make(chan bool)
	go func() {
		for msg := range msgs {
			//	// 3. 调用callback方法处理消息
			processSuc := callback(msg.Body)
			if !processSuc {
				// TODO；将任务写入死信队列
			}
		}
	}()
	<-done // 无缓冲channel在没有新消息过来则会一直阻塞
	// 关闭rabbit channel
	channel.Close()
}
