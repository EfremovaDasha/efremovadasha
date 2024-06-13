//Реализация различных случаев при работе с каналами:

package main

import (
 "fmt"
 "time"
)

func main() {
 // Простой канал
 simpleChan()

 // Закрытие канала
 closeChannel()

 // Буферизованный канал
 bufferChannel()

 // Множественные отправители и получатели
 multiSenderReceiver()

 // Селект с каналами
 selectWithChannels()

 // Канал с таймаутом
 channelWithTimeout()
}

func simpleChan() {
 fmt.Println("Simple Channel:")
 ch := make(chan int)
 ch <- 42
 value := <-ch
 fmt.Println("Received value:", value)
 fmt.Println()
}

func closeChannel() {
 fmt.Println("Closing Channel:")
 ch := make(chan int)
 ch <- 1
 ch <- 2
 ch <- 3
 close(ch)

 for value := range ch {
  fmt.Println("Received value:", value)
 }
 fmt.Println()
}

func bufferChannel() {
 fmt.Println("Buffered Channel:")
 ch := make(chan int, 2)
 ch <- 1
 ch <- 2
 fmt.Println("Channel length:", len(ch))
 fmt.Println("Received value:", <-ch)
 fmt.Println("Received value:", <-ch)
 fmt.Println("Channel length:", len(ch))
 fmt.Println()
}

func multiSenderReceiver() {
 fmt.Println("Multiple Senders and Receivers:")
 ch := make(chan int)
 go sender(ch, 1)
 go sender(ch, 2)
 receiver(ch)
 receiver(ch)
 fmt.Println()
}

func sender(ch chan<- int, value int) {
 ch <- value
}

func receiver(ch <-chan int) {
 fmt.Println("Received value:", <-ch)
}

func selectWithChannels() {
 fmt.Println("Select with Channels:")
 ch1 := make(chan int)
 ch2 := make(chan int)

 go func() {
  time.Sleep(2 * time.Second)
  ch1 <- 1
 }()

 go func() {
  time.Sleep(1 * time.Second)
  ch2 <- 2
 }()

 select {
 case value := <-ch1:
  fmt.Println("Received from ch1:", value)
 case value := <-ch2:
  fmt.Println("Received from ch2:", value)
 }
 fmt.Println()
}

func channelWithTimeout() {
 fmt.Println("Channel with Timeout:")
 ch := make(chan int)

 go func() {
  time.Sleep(2 * time.Second)
  ch <- 42
 }()

 select {
 case value := <-ch:
  fmt.Println("Received value:", value)
 case <-time.After(1 * time.Second):
  fmt.Println("Timeout")
 }
 fmt.Println()
}