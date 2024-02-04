package users

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router) {

	router.Post("/user-image", func(c *fiber.Ctx) error {
		st, err := c.FormFile("userimage")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
		}
		fmt.Println("nombre", st.Filename)
		fmt.Println("size", st.Size)
		var arrName = strings.Split(st.Filename, ".")
		fmt.Println("el array", arrName)
		var extFile = arrName[len(arrName)-1] //! filepath.Ext("ide.gfg.") => .org
		fmt.Println("extencion\n", extFile)

		var finalName = fmt.Sprintf("%s%s.%s", strings.Join(arrName[:len(arrName) - 1], ""), time.Now().Format("12-05-24"), extFile)
		fmt.Println("nombre del archivo", finalName)
		err = c.SaveFile(st, fmt.Sprintf("./storage/%s", finalName))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
		}

		return c.Status(fiber.StatusOK).Send([]byte("saved image"))
	})

	// para envair archivos que se puedan descargar
	router.Get("/image", func(c *fiber.Ctx) error {
		var fileNmae = c.Query("name") // logo-genericpng24-07-41.png

		imgBytes, err := os.ReadFile(fmt.Sprintf("./storage/%s", fileNmae))
		if err != nil {
			fmt.Println("error: ", err.Error())
			return c.SendStatus(fiber.StatusBadRequest)
		}

		// Configurar el encabezado para la descarga
		c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileNmae))
		c.Set("Content-Type", "image/*")

		return c.Send(imgBytes)
	})
}
