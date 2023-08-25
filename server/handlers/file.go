package handlers

import (
	"fmt"
	"strings"

	"github.com/goellavish10/img-upload-go/utils"
	"github.com/gofiber/fiber/v2"
)

func FileUpload(c *fiber.Ctx) error {
	fmt.Println("File upload handler")

	file, err := c.FormFile("file")

	if err != nil {
		fmt.Println("Error parsing the file from client: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Internal server error while parsing the file",
			Success: false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse{
		Message: "File parsed and information extracted!",
		Success: true,
		Data: utils.Data{
			Filename: strings.Split(file.Filename, ".")[0],
			FileExt:  strings.Split(file.Filename, ".")[1],
			Filesize: file.Size,
			FileType: file.Header["Content-Type"][0],
		},
	})
}
