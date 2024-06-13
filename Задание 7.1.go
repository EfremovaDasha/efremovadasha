// Клиент
package main

import (
 "bufio"
 "fmt"
 "net"
 "os"
 "strings"
)

func main() {
 conn, err := net.Dial("tcp", ":8080")
 if err != nil {
  fmt.Println("Ошибка при подключении к серверу:", err)
  return
 }
 defer conn.Close()

 reader := bufio.NewReader(os.Stdin)
 fmt.Print("Введите свой ник: ")
 nick, _ := reader.ReadString('\n')
 nick = strings.TrimSpace(nick)

 _, err = fmt.Fprintf(conn, "%s\n", nick)
 if err != nil {
  fmt.Println("Ошибка при отправке ника:", err)
  return
 }

 response, err := bufio.NewReader(conn).ReadString('\n')
 if err != nil {
  fmt.Println("Ошибка при получении ответа от сервера:", err)
  return
 }

 if strings.TrimSpace(response) != "" {
  fmt.Println(response)
  return
 }

 for {
  fmt.Print("Введите ник получателя и сообщение: ")
  message, _ := reader.ReadString('\n')
  message = strings.TrimSpace(message)

  _, err = fmt.Fprintf(conn, "%s\n", message)
  if err != nil {
   fmt.Println("Ошибка при отправке сообщения:", err)
   return
  }

  response, err := bufio.NewReader(conn).ReadString('\n')
  if err != nil {
   fmt.Println("Ошибка при получении ответа от сервера:", err)
   return
  }

  fmt.Println(response)
 }
}