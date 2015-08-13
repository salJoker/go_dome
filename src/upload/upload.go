package upload

import(
	"log"
	"net/http"
	"os"
	"io"
	"github.com/labstack/echo"
)

func UploadFile(c *echo.Context)error{
	req := c.Request()
	req.ParseMultipartForm(16 << 20)
	files := req.MultipartForm.File["files"]
	for _,file:= range files{
		src,err := file.Open()
		if err != nil{
			return err
		}
		defer src.Close()

		log.Println("FileName:",file.Filename)

		dst,err := os.Create(file.Filename)
		if err != nil{
			return err
		}
		defer dst.Close()

		if _,err := io.Copy(dst,src);err != nil {
			return err
		}
	}
	return c.String(http.StatusOK,"Upload Success!")
}