package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"quiz.com/quiz/internal/entity"
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

type ConnectPacket struct {
	Code string `json:"code"`
	Name string `json:"Name"`
}

type HostGamePacket struct {
	QuizId string `json:"quizId"`
}

type QuestionShowPacket struct {
	Question entity.QuizQuestion `json:"question"`
}

func (c *NetService) packetIdToPacket(packetId uint8) any {
	switch packetId {
	case 0:
		{
			return &ConnectPacket{}
		}
	case 1:
		{
			return &HostGamePacket{}
		}
	}
	return nil
}

func (c *NetService) packetToPacketId(packet any) (uint8, error) {
	switch packet.(type) {
	case QuestionShowPacket:
		{
			return 2, nil
		}
	}
	return 0, errors.New("invalid packet type")
}

func (c *NetService) OnDisconnect(con *websocket.Conn) {
}

func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {
	if len(msg) < 2 {
		return
	}
	packetId := msg[0]
	packetData := msg[1:]
	packet := c.packetIdToPacket(packetId)
	if packet == nil {
		return
	}
	err := json.Unmarshal(packetData, &packet)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(packetId)
	fmt.Println(string(packetData))

	switch packet := packet.(type) {
	case ConnectPacket:
		{
			fmt.Println(packet.Name)
		}
	case HostGamePacket:
		{
			fmt.Println(packet.QuizId)
		}
	}

}

// 	str := string(msg)
// 	parts := strings.Split(str, ":")
// 	cmd := parts[0]
// 	argument := parts[1]
// 	switch cmd {
// 	case "host":
// 		{
// 			fmt.Println("host quizz", argument)
// 			c.host = con
// 			c.tick = 100
// 			go func(){
// 				for{
// 					c.tick--
// 					c.host.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(c.tick)))
// 					time.Sleep(time.Second)
// 				}
// 			}()
// 			break
// 		}
// 	case "join":
// 		{
// 			fmt.Println("join code", argument)
// 			if c.host != nil {
// 				err := c.host.WriteMessage(websocket.TextMessage, []byte("A player has joined"))
// 				if err != nil {
// 					fmt.Println("Error writing message to host:", err)
// 				}
// 			} else {
// 				fmt.Println("No host is connected")
// 			}
// 			break
// 		}
// 	}
// }

func (c *NetService) PacketToBytes(packet any) ([]byte, error) {
	packetId, err := c.packetToPacketId(packet)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(packet)
	if err != nil {
		return nil, err
	}
	final := append([]byte{packetId}, bytes...)
	return final, nil
}

func (c *NetService) SendPacket(connection *websocket.Conn, packet any) error {
	bytes, err := c.PacketToBytes(packet)
	if err != nil {
		return err
	}
	return connection.WriteMessage(websocket.BinaryMessage, bytes)
}
