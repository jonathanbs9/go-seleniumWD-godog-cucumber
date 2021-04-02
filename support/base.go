package support

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

// WDInit retorna una instancia de WebDriver
func WDInit() selenium.WebDriver {
	var err error
	//caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	caps := selenium.Capabilities{"browserName": "chrome"}

	driver, err = selenium.NewRemote(caps, "")
	if err != nil {
		fmt.Println("support/base |  Error al instanciar el driver de Selenium : ", err.Error())
		log.Println(err.Error())
	}

	driver.SetImplicitWaitTimeout(time.Second * 10)

	driver.ResizeWindow("note", 1280, 800)

	return driver
}

func SaveImage(foto []byte, name string) {
	img, _, _ := image.Decode(bytes.NewReader(foto))

	out, err := os.Create("./log/screenshots/" + name + ".png")
	if err != nil {
		fmt.Println("Error a realizar la captura de escenario! ")
		os.Exit(1)
	}

	err = png.Encode(out, img)
	if err != nil {
		fmt.Println("Error | " + err.Error())
		os.Exit(1)
	}

}
