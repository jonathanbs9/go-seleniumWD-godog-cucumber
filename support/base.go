package support

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

// WDInit retorna una instancia de WebDriver
func WDInit() selenium.WebDriver {
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "Chrome"})
	driver, err = selenium.NewRemote(caps, "")
	if err != nil {
		fmt.Println("Error al instanciar el driver de Selenium : ", err.Error())
	}

	driver.SetImplicitWaitTimeout(time.Second * 10)

	driver.ResizeWindow("note", 1280, 800)

	return driver
}

func retornaMensaje() string {
	return "Hola! "
}

func retornaValor() float64 {
	return 200.00
}
