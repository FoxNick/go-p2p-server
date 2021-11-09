package app

import (
	"fmt"
	// "log"
	"net/http"
	// "runtime"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/midoks/go-p2p-server/internal/client"
	"github.com/midoks/go-p2p-server/internal/conf"
	"github.com/midoks/go-p2p-server/internal/hub"
	"github.com/midoks/go-p2p-server/internal/tools"
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

	clientId := client.New(id, ws, true)
	fmt.Println(clientId)

	verNum := tools.GetVersionNum(conf.App.Version)
	clientId.SendMsgVersion(verNum)

	hub.DoRegister(clientId)

	go func() {

		for {
			// msg, err = wsutil.ReadClientMessage(conn, msg[:0])
			// if err != nil {
			// 	log.Infof("read message error: %v", err)
			// 	break
			// }

			mt, message, err := ws.ReadMessage()
			if err != nil {
				// logger.Errorf("read websocket msg: %v", err)
				fmt.Println("err:", err, "id:", mt)
				break
			}

			fmt.Println("go func:", mt, message)
			// for _, m := range message {
			// 	// ping
			// if m.OpCode.IsControl() {
			// 	c.UpdateTs()
			// 	//log.Warnf("receive ping from %s platform %s", id, platform)
			// 	err := wsutil.HandleClientControlMessage(conn, m)
			// 	if err != nil {
			// 		fmt.Println("handle control error: ", err)
			// 	}
			// 	continue
			// }
			// 	data := bytes.TrimSpace(bytes.Replace(m.Payload, newline, space, -1))
			// 	hdr, err := handler.NewHandler(data, c)
			// 	if err != nil {
			// 		// 心跳包
			// 		c.UpdateTs()
			// 		log.Infof("NewHandler " + err.Error())
			// 	} else {
			// 		hdr.Handle()
			// 	}
			// }
		}
	}()
	// for {

	// 	// go func(ws *websocket.Conn) {
	// 	mt, message, err := ws.ReadMessage()
	// 	if err != nil {
	// 		// logger.Errorf("read websocket msg: %v", err)
	// 		fmt.Println("err:", err, "id:", mt)
	// 		break
	// 	}

	// 	fmt.Println(mt, string(message), err, runtime.NumGoroutine())

	// 	// }(ws)
	// 	now := time.Now().Format("2006-01-02 15:04:05")
	// 	log.Printf("recv: %s", message)

	// 	err = ws.WriteMessage(mt, []byte(now))
	// 	if err != nil {
	// 		log.Println("write:", err)
	// 		break
	// 	}

	// }
}
