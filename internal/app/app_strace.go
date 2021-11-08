package app

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//websocket
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//websocket实现
func websocketReqMethod(c *gin.Context) {
	// c.Request.ParseForm()
	// id := c.Request.Form.Get("id")
	id := c.Query("id")

	fmt.Println("websocket id:[", id, "]")
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// defer ws.Close()

	for {

		// go func(ws *websocket.Conn) {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			// logger.Errorf("read websocket msg: %v", err)
			fmt.Println("err:", err, "id:", mt)
			// logger.Errorf("read websocket msg: %v", err)
			break
		}

		fmt.Println(mt, string(message), err, runtime.NumGoroutine())

		// }(ws)
		now := time.Now().Format("2006-01-02 15:04:05")
		log.Printf("recv: %s", message)

		err = ws.WriteMessage(mt, []byte(now))
		if err != nil {
			log.Println("write:", err)
			break
		}

	}
}