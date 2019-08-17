package ctl

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// NewWs 创建登录管理控制器
func NewWs() *Ws {
	return &Ws{}
}

// Login 登录管理
// @Name Login
// @Description 登录管理接口
type Ws struct{}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Ping webSocket请求Ping 返回pong
func (w *Ws) Ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	for {
		// 读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			// 客户端关闭连接时也会进入
			fmt.Println(err)
			break
		}
		// msg := &data{}
		// json.Unmarshal(message, msg)
		// fmt.Println(msg)
		fmt.Println(mt)
		fmt.Println(message)
		fmt.Println(string(message))

		// 如果客户端发送ping就返回pong,否则数据原封不动返还给客户端
		if string(message) == "ping" {
			message = []byte("pong - " + time.Now().Format("15:04:05"))
		}
		// 写入ws数据 二进制返回
		err = ws.WriteMessage(mt, message)
		// 返回JSON字符串，借助gin的gin.H实现
		// v := gin.H{"message": msg}
		// err = ws.WriteJSON(v)
		if err != nil {
			break
		}
	}
	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()

	//for {
	//	select {
	//	case t := <-ticker.C:
	//		err := ws.WriteMessage(websocket.TextMessage, []byte(t.String()))
	//		fmt.Println("writing time: ", t.String())
	//		if err != nil {
	//			log.Println("write:", err)
	//			return
	//		}
	//	}
	//}

}
