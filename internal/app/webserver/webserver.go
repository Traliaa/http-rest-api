package webserver

import (
	"github.com/Traliaa/http-rest-api/internal/app/model"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func SendClient(w *websocket.Conn) {

	for {
		m := model.SmartDevice{}

		err := w.ReadJSON(&m)
		if err != nil {
			logrus.Error("Error reading json.", err)
		}

		logrus.Infof("recv: %#v\n", m)

		if err = w.WriteJSON(m); err != nil {
			logrus.Error("write:", err)
		}
	}
}
