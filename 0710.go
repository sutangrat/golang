func clientSide() {
	client, err := rpc.Dial("tcp","127.0.0.1:8191")
	if err != nil {
		fmt.Println(err)
		return
	}
	var reply float64
	
}