func main()  {
	go serverSide()
	go clientSide()

	var input string
	fmt.Scan(&input)
}	
