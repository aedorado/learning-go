// 6. Panics are not fatal
// they are fatal only when we panic upto the go runtime
// in which case, the go runtime kills this application

fmt.Println("1 Start")
defer tryToRecover()
panic("cannot continue the program")
fmt.Println("3 Post Panic")


func tryToRecover() {
	// recover returns nil if the application is not panicing
	// 	   returns the error that has caused the panic
	if err := recover(); err != nil {
		fmt.Println("Error: ", err)
	}
}