package worker

//Run 导出的方法
func Run() {
	go func() {
		runPullGoods()
		runUpdateGoods()
	}()
}
