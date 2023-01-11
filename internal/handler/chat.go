package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 0,
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Fatal("error connect to websocket")
		return
	}
	for {
		messageType, data, err := connection.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(messageType, data)
		err = connection.WriteMessage(messageType, data)
		if err != nil {
			return
		}
	}
}

func (h *Handler) chat(c *gin.Context) {
	websocketHandler(c.Writer, c.Request)
}
