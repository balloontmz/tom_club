package worker

import (
	"sync"

	"tom_club/app/server"

	"github.com/labstack/gommon/log"
)

var (
	index     int
	indexChan = make(chan int, 20) // 定义一个 worker 全局的 chan，该值相当于并发量
	mux       sync.Mutex
)

func runPullGoods() {
	// 运行该任务前首先将 index 重定义为 0
	index = 0
	// 应该先获取容许的最大值
	go func() {
		for {
			mux.Lock()
			log.Print("加锁后获取到的链接为：", index)
			// 此处应该先判断页码是否超限。如果是则 break
			index++
			indexChan <- index // 同步写入
			mux.Unlock()
		}
	}()
	go func() {
		for i := range indexChan {
			go getGoods(0, i) //读取到值立马运行
		}
	}()
}

func getGoods(try, i int) {
	log.Print("获取到页数为", i)
	t, err := server.GetGoods(i)
	if err != nil {
		log.Print("拉取出错，重新拉取, 重试次数为：", try, "错误信息为：", err)
		try++
		if try < 3 { //拉取出错只重新拉取三次
			go getGoods(try, i)
		}
		return
	}
	goodsChan <- t
	return
}
