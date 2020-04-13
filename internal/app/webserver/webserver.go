package webserver

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func SendClient(w *websocket.Conn) {
	for {
		mt, message, err := w.ReadMessage()
		if err != nil {
			logrus.Error("read:", err)
			break
		}
		logrus.Infof("recv: %s", message)
		err = w.WriteMessage(mt, message)
		if err != nil {
			logrus.Error("write:", err)
			break
		}
	}
}
