package halper

import "github.com/gofiber/fiber/v2/log"

func PanicIfError(err error) {
	if err != nil {
		log.Info(err.Error())
		panic(err)
	}
}
