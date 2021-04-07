package support

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/tebeka/selenium"
)

var (
	driver           selenium.WebDriver
	chromeDriverPath = `C:\Users\jonat\go\src\github.com\jonathanbs9\go-seleniumWD-godog-cucumber\resources\chromedriver.exe`
	seleniumPath     = `C:\Users\jonat\go\src\github.com\jonathanbs9\go-seleniumWD-godog-cucumber\resources\selenium-server-standalone-3.141.59.jar`
	port             = 4444
)

// WDInit retorna una instancia de WebDriver
func WDInit() selenium.WebDriver {
	var err error

	ops := []selenium.ServiceOption{
		selenium.ChromeDriver(seleniumPath),
	}

	//service, err := selenium.NewSeleniumService(seleniumPath, port, ops...)
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, ops...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
	}
	//Delay service shutdown
	defer service.Stop()

	//log.Println("Service => ", service)

	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	//log.Println("Capabilities => ", caps)

	driver, err := selenium.NewRemote(caps, "")

	if err != nil {
		log.Println("support/base |  Error al instanciar el driver de Selenium : ", err.Error())
	}
	//driver.ResizeWindow("note", 1920, 1080)
	return driver
}

// Función para guarda  screenshot
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
