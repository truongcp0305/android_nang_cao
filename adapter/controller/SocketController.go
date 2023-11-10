package controller

import (
	"android-service/adapter/incoming"
	"android-service/usecase/service"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type SocketController struct {
	socketService service.SocketService
	room          service.Room
}

func NewSocketController(s service.SocketService, room service.Room) SocketController {
	return SocketController{
		socketService: s,
		room:          room,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (sc *SocketController) EchoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("get message")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Lỗi khi nâng cấp kết nối:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Lỗi khi đọc thông điệp:", err)
			return
		}

		log.Printf("Nhận thông điệp: %s", p)

		// Đáp ứng lại với thông điệp đã nhận
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("Lỗi khi gửi thông điệp:", err)
			return
		}
	}
}

func (sc *SocketController) RoomHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Lỗi khi nâng cấp kết nối:", err)
		return
	}
	defer conn.Close()
	sc.room.Join(conn)
	defer sc.room.Leave(conn)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Lỗi khi đọc thông điệp:", err)
			return
		}
		log.Printf("Nhận thông điệp: %s", p)
		sc.room.Broadcast(p, conn)
	}
}

func (sc *SocketController) Status(c echo.Context) error {
	var params incoming.StatusIncoming
	c.Bind(&params)
	stt := params.GetModel()
	if params.Id == "" {
		return c.JSON(http.StatusBadRequest, errors.New("Invalid param id"))
	}
	res := sc.socketService.GetStatus(*stt)
	return c.JSON(http.StatusOK, res)
}

func (sc *SocketController) Join(c echo.Context) error {
	id := c.Param("id")
	level := c.Param("level")
	if id == "" || level == "" {
		return c.JSON(http.StatusBadRequest, errors.New("Invalid param id"))
	}
	res := sc.socketService.Join(id, level)
	return c.JSON(http.StatusOK, res)
}

func (sc *SocketController) Leave(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, errors.New("Invalid param id"))
	}
	sc.socketService.Leave(id)
	return c.JSON(http.StatusOK, "leave sucsess")
}

// func (sc *SocketController) JoinRoom(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Print("Lỗi khi nâng cấp kết nối:", err)
// 		return
// 	}

// 	sc.room.Join(conn)
// 	for {
// 		_, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("Lỗi khi đọc thông điệp:", err)
// 			return
// 		}

// 		log.Printf("Nhận thông điệp: %s", p)
// 		sc.room.Broadcast(p, conn)
// 	}
// }
