func serverSide() {
	rpc.Register (new(ExamleStruct))
	ln, _ := net.Listen("tcp","8191")
}