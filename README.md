## Go

#### Initialize Go Applicaiton
- go mod init booking-app

- it will be wait for goroutine task completion
- go is basically concurrency will carate goroutine as green theread its not create OS lavel thread like as queue
- Abstraction of an actual thread
- we can run hundreads of thoudands of millions of goroutines wihout creating OS lavel thread and wihout affecting the performance
- Build-in Functionality for goroutines to talk with one another
- 
- if we dont use loop then wheen complete the main thread task then applicaiton will exit
-  so main therad will not wait for goroutine
-  for this we need to fix for goroutine task completion then main thread will exit
-  package "sync" provided basis synchronization fucntionality

- call function for sending ticket using concurrency as goroutine with go keyword
- A goroutine is a thread that is running in the background
- A goroutine is a light-weight thread managed by the Go runtime
- need to be wait for goroutine task completion in main thread with use wg.Wait() and add wg.Add(1)
- wg.Add(1) // Sets the number of goroutines to wait for (increases the counter by the provided number)
- go sendTicket(userTickets, firstName, lastName, email)
- time.Sleep(10 * time.Second) // when run this it will be block the code and wait for 10 seconds so we need to concurrency 
- wg.Done() // Decrements the waitGroup counter by 1 So this is called by the goroutine to indicate that the goroutine is done/completed

- 
