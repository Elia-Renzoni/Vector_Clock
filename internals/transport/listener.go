package transport

type Listener struct {
	message       <-chan string
	goroutinePool <-chan int
}

func (l *Listener) StartListening() {

}

func (l *Listener) HandleMessage() {

}
