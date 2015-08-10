package webchat

import(
	"fmt"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

var conns []*websocket.Conn

func WebChat(c *echo.Context)(err error){
	ws := c.Socket()
	msg := ""

	conns = append(conns,ws)

	for{
		if err = websocket.Message.Receive(ws,&msg);err != nil{
			return
		}
		fmt.Println(msg)

		for _,conn := range conns {
			//待解决Conn关闭重连问题
			serve_msg := fmt.Sprintf("%s : %s",c.Request().RemoteAddr,msg)
			if err = websocket.Message.Send(conn,serve_msg);err != nil{
				return
			}
		}
	}
	return nil
}
