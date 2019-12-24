func clientSide() {
	client, err := rpc.Dial("tcp","127.0.0.1:8191")
	if err != nil {
		fmt.Println(err)
		return
	}
	var reply float64
	args := []float64{5, 7, 9, 11, 23}
	err = client.Call ("ExampleStruct.Plus",args,&reply)
	if err != nil {
		fmt.Printlne(err)
	}
}