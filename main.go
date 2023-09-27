/* package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// @title 		Boilerplate API
// @version		1.0
// @description An API in Go using Gin framework
// @host 		localhost:8000

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil{
		log.Println(err)
	}
	reader(ws)
}
func main() {
	// fx.New(bootstrap.Module).Run()
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRoutes() {
	http.HandleFunc("/api", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}
 */

 package main

import (
    "log"
    "net/http"
    "strings"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var todoList []string

func getCmd(input string) string {
    inputArr := strings.Split(input, " ")
    return inputArr[0]
}

func getMessage(input string) string {
    inputArr := strings.Split(input, " ")
    var result string
    for i := 1; i < len(inputArr); i++ {
        result += inputArr[i]
    }
    return result
}

func updateTodoList(input string) {
    tmpList := todoList
    todoList = []string{}
    for _, val := range tmpList {
        if val == input {
            continue
        }
        todoList = append(todoList, val)
    }
}

func main() {

    http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
        // Upgrade upgrades the HTTP server connection to the WebSocket protocol.
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Print("upgrade failed: ", err)
            return
        }
        defer conn.Close()

        // Continuosly read and write message
        for {
            mt, message, err := conn.ReadMessage()
            if err != nil {
                log.Println("read failed:", err)
                break
            }
            input := string(message)
            cmd := getCmd(input)
            msg := getMessage(input)
            if cmd == "add" {
                todoList = append(todoList, msg)
            } else if cmd == "done" {
                updateTodoList(msg)
            }
            output := "Current Todos: \n"
            for _, todo := range todoList {
                output += "\n - " + todo + "\n"
            }
            output += "\n----------------------------------------"
            message = []byte(output)
            err = conn.WriteMessage(mt, message)
            if err != nil {
                log.Println("write failed:", err)
                break
            }
        }
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "websockets.html")
    })

    http.ListenAndServe(":8080", nil)
}