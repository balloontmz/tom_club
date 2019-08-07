package worker

import (
	"time"

	"github.com/labstack/gommon/log"
)

//Run 导出的方法
func Run() {
	go func (){
		for {
			log.Print("初始化进入或者执行定时任务完毕")
			// 每天零点执行一次日志生成命令
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)

			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())

			t := time.NewTimer(next.Sub(now))

			<-t.C

			// 将日志写入文件，定时执行
			runPullGoods()
			runUpdateGoods()
		}
	}()
}
