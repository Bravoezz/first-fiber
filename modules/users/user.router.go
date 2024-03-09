package users

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
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

		var finalName = fmt.Sprintf("%s%s.%s", strings.Join(arrName[:len(arrName)-1], ""), time.Now().Format("12-05-24"), extFile)
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
		c.Set(fiber.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", fileNmae)) //! "Content-Disposition"
		c.Set(fiber.HeaderContentType, "image/*") //! "Content-Type"
		
		return c.Send(imgBytes)
	})

	type UserDto struct {
		Name     string `validate:"required"`
		Number   string `validate:"required,len=5"`
		Document string `validate:"required"`
		Email    string `validate:"required,email"`
		Password string `validate:"required"`
	}

	router.Post("/register-validate",
		validateMiddleware[UserDto](),
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusOK).JSON("todo bien")
		},
	)

	// var n interface{} = "4"
	// fmt.Println("hola", n.(int))

	type LoginUserDto struct {
		User   string 
		Number string `validate:"required,len=5"`
		Age    int    `validate:"required,gte=0,lte=130"`
		Gay    bool   `validate:"boolean"`
	}

	router.Post("/login-validate",
		validateMiddleware[LoginUserDto](),
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusOK).JSON("usuarios logueado con validaciones")
		},
	)

	// ruta con validacion - feo
	router.Post("/register", func(c *fiber.Ctx) error {
		var userDto UserDto
		c.BodyParser(&userDto)
		// if err != nil {
		// 	fmt.Println("error: ", err.Error())
		// 	return c.SendStatus(fiber.StatusInternalServerError)
		// }

		validate := validator.New()
		if err := validate.Struct(&userDto); err != nil {
			var errorMessage []string

			for _, err := range err.(validator.ValidationErrors) {
				errorMessage = append(errorMessage, fmt.Sprintf("Field '%s' validation failed on tag '%s'", err.Field(), err.Tag()))
			}

			return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"error": errorMessage,
				"res":   false,
			})
		}

		fmt.Printf("la data: %v \n", userDto)
		fmt.Println("uno", userDto.Name)
		fmt.Println("dos", userDto.Number)
		fmt.Println("tres", userDto.Document)
		fmt.Println("cuatro", userDto.Email)
		fmt.Println("cinco", userDto.Password)

		dt := map[string]interface{}{"hola": "algo", "name": userDto.Number}

		return c.JSON(dt)
	})

}

func validateMiddleware[T any]() fiber.Handler {
	validate := validator.New()
	return func(c *fiber.Ctx) error {
		var dto T
		errBdy := c.BodyParser(&dto)
		if errBdy != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
				"res":     false,
				"message": fmt.Sprintf("Error body %s", errBdy.Error()),
			})
		}

		if err := validate.Struct(&dto); err != nil {
			var errorMessage []string
			for _, err := range err.(validator.ValidationErrors) {
				errorMessage = append(errorMessage, fmt.Sprintf(" '%s' no valido", err.Field()))
			}
			return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
				"res":     false,
				"message": errorMessage,
			})
		}
		return c.Next()
	}
}