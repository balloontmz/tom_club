package conf

import (
	"io"
	"os"
	"time"

	"github.com/labstack/gommon/log"
)

//InitLog 初始化日志器的记录选项，初步测试成功
func InitLog() {
	// 重启程序时重新设置日志存放位置
	setLogFile()

	go func() {

		for {
			// 每天零点执行一次日志生成命令
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)

			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())

			t := time.NewTimer(next.Sub(now))

			<-t.C

			// 将日志写入文件，定时执行
			setLogFile()
		}
	}()
}

// 设置日志格式
func setLogFile() {
	f, _ := os.OpenFile("log/echo"+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755) // 追加或者新建文件

	w := io.MultiWriter(f)

	log.SetOutput(w)
	log.SetHeader("${time_rfc3339} ${level} ${prefix} ${short_file} ${line}")
}
