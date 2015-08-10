package user
import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type User struct{
	Name string `json:name`
	Email string `json:email`
}

var users []*User

func Register(c *echo.Context)error{

	name := c.Form("name")
	email := c.Form("email")

	user := &User{ name,email, }

	users = append(users,user)

	log.Println("users:",users)
	/**
	if err := c.Bind(users);err != nil{
		log.Println("bind:"+err.Error())
		return err
	}**/
	return c.JSON(http.StatusOK,users)
}