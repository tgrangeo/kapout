package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/contrib/websocket"
)

type NetService struct {
	quizService *QuizService

	host *websocket.Conn
	tick int
}

func Net(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
	}
}

func (c *NetService) OnDisconnect(con *websocket.Conn) {
}

func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {
	str := string(msg)
	parts := strings.Split(str, ":")
	cmd := parts[0]
	argument := parts[1]
	switch cmd {
	case "host":
		{
			fmt.Println("host quizz", argument)
			c.host = con
			c.tick = 100
			go func(){
				for{
					c.tick--
					c.host.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(c.tick)))
					time.Sleep(time.Second)
				}
			}()
			break
		}
	case "join":
		{
			fmt.Println("join code", argument)
			if c.host != nil {
				err := c.host.WriteMessage(websocket.TextMessage, []byte("A player has joined"))
				if err != nil {
					fmt.Println("Error writing message to host:", err)
				}
			} else {
				fmt.Println("No host is connected")
			}
			break
		}
	}
}
