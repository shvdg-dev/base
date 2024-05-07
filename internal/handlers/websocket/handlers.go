package websocket

import (
	"context"
	"net/http"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	connection, err := websocket.Accept(writer, request, nil)
	if err != nil {
		//...
	}
	defer func(connection *websocket.Conn) {
		err := connection.CloseNow()
		if err != nil {
			//...
		}
	}(connection)

	ctx, cancel := context.WithTimeout(request.Context(), time.Second*10)
	defer cancel()

	var value interface{}
	err = wsjson.Read(ctx, connection, &value)
	if err != nil {
		//...
	}
	err = connection.Close(websocket.StatusNormalClosure, "")
	if err != nil {
		//...
	}
}
