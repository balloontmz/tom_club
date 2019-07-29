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
		mux.Lock()
		log.Print("加锁后获取到的链接为：", index)
		// 此处应该先判断页码是否超限。如果是则 break
		index++
		indexChan <- index
		mux.Unlock()

		// 采用拼装好的链接进行
		go getGoods(0)
	}()
}

func getGoods(try int) {
	i := <-indexChan
	log.Print("获取到页数为", i)
	t, err := server.GetGoods(i)
	if err != nil {
		log.Print("拉取出错，重新拉取, 重试次数为：", try)
		try++
		if try < 3 { //拉取出错只重新拉取三次
			go getGoods(try)
		}
		return
	}
	goodsChan <- t
	return
}
