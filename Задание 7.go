// Сервер
package main

import (
 "bufio"
 "fmt"
 "net"
 "strings"
 "sync"
)

type Client struct {
 Conn     net.Conn
 Nickname string
}

func main() {
 clients := make(map[string]*Client)
 mutex := sync.Mutex{}

 listener, err := net.Listen("tcp", ":8080")
 if err != nil {
  fmt.Println("Ошибка при запуске сервера:", err)
  return
 }
 defer listener.Close()

 fmt.Println("Сервер запущен и ожидает подключений...")

 for {
  conn, err := listener.Accept()
  if err != nil {
   fmt.Println("Ошибка при принятии подключения:", err)
   continue
  }

  go handleConnection(conn, clients, &mutex)
 }
}

func handleConnection(conn net.Conn, clients map[string]*Client, mutex *sync.Mutex) {
 defer conn.Close()

 reader := bufio.NewReader(conn)
 nick, err := reader.ReadString('\n')
 if err != nil {
  fmt.Println("Ошибка при чтении ника:", err)
  return
 }
 nick = strings.TrimSpace(nick)

 mutex.Lock()
 _, ok := clients[nick]
 if ok {
  fmt.Fprintf(conn, "Ошибка: Этот ник уже используется.\n")
  return
 }
 client := &Client{
  Conn:     conn,
  Nickname: nick,
 }
 clients[nick] = client
 mutex.Unlock()

 fmt.Printf("Новое подключение: %s\n", nick)

 for {
  message, err := reader.ReadString('\n')
  if err != nil {
   fmt.Printf("Клиент %s отключился: %v\n", nick, err)
   mutex.Lock()
   delete(clients, nick)
   mutex.Unlock()
   return
  }

  message = strings.TrimSpace(message)
  parts := strings.Split(message, " ")
  if len(parts) < 2 {
   fmt.Fprintf(conn, "Некорректный формат сообщения.\n")
   continue
  }

  targetNick := parts[0]
  targetMessage := strings.Join(parts[1:], " ")

  mutex.Lock()
  targetClient, ok := clients[targetNick]
  mutex.Unlock()

  if !ok {
   fmt.Fprintf(conn, "Ошибка: Пользователь %s не найден.\n", targetNick)
   continue
  }

  fmt.Fprintf(targetClient.Conn, "%s: %s\n", nick, targetMessage)
 }
}

