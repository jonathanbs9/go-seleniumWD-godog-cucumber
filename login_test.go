package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/jonathanbs9/go-seleniumWD-godog-cucumber/support"
	"github.com/tebeka/selenium"
)

var (
	Driver selenium.WebDriver
)

func queAccediALaPaginaPrincipal() error {
	Driver.Get("https://amanti.websitewelcome.com:2096/")
	log.Println("Accedo pagina principal -> Driver.Get() ")
	return nil
}

func hagoElLoginConY(email, contraseña string) error {
	campoEmail, err := Driver.FindElement(selenium.ByID, "user")

	if err != nil {
		fmt.Println("Error  |  ", err.Error())
	}
	campoEmail.SendKeys(email)

	campoContraseña, err := Driver.FindElement(selenium.ByID, "pass")
	if err != nil {
		fmt.Println("Error  |  ", err.Error())
	}
	campoContraseña.SendKeys(contraseña)

	botonIniciarSesion, err := Driver.FindElement(selenium.ByID, "login_submit")
	if err != nil {
		fmt.Println("Error  |  ", err.Error())
	}
	botonIniciarSesion.Click()

	return nil
}

func estoyAutenticadoConXito() (err error) {

	//topRight, err := Driver.FindElement(selenium.ByClassName, "username")
	topRight, err := Driver.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div[2]/span")
	if err != nil {
		log.Println("Error  |  Al buscar el element: ", err.Error())
	}

	email, err := topRight.Text()
	if err != nil {
		log.Println("Error  |  Al buscar el element: ", err.Error())
	}
	log.Println("Text del div => " + email)

	if email != "jonathans@avalith.net" {
		return fmt.Errorf("Error  al validar autenticacion de usuario! ")
	}

	return nil
}

func deboVerElSiguienteMensaje(mensaje string) error {

	//divAlerta, err := Driver.FindElement(selenium.ByID, "login-status-message")
	divAlerta, err := Driver.FindElement(selenium.ByXPATH, "//*[@id='login-status-message']")
	if err != nil {
		return fmt.Errorf("Error al validar mensaje ! ")
	}
	log.Println(divAlerta)
	alerta, err := divAlerta.Text()
	if err != nil {
		log.Println("Error al devolver el mensaje : ", err.Error())
	}
	log.Printf(alerta)

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		Driver = support.WDInit()

		ctx.Step(`^que accedi a la pagina principal$`, queAccediALaPaginaPrincipal)
		ctx.Step(`^hago el login con "([^"]*)" y "([^"]*)"$`, hagoElLoginConY)
		Driver.SetImplicitWaitTimeout(time.Second * 10)
		ctx.Step(`^estoy autenticado con éxito$`, estoyAutenticadoConXito)
		Driver.SetImplicitWaitTimeout(time.Second * 10)
		ctx.Step(`^debo ver el siguiente mensaje "([^"]*)"$`, deboVerElSiguienteMensaje)
	})

	ctx.AfterScenario(func(s *godog.Scenario, err error) {
		//sc := s.(*gherkin.RuleTypeScenario.Name())
		rgex := regexp.MustCompile("[^0-9a-zA-Z]+")
		fileName := strings.ToLower(rgex.ReplaceAllString(s.Name, "_"))
		screenshot, _ := Driver.Screenshot()

		support.SaveImage(screenshot, fileName)
		Driver.SetPageLoadTimeout(10 * time.Second)
		//Driver.Quit()
	})

}

/*s.AfterScenario(func(i interface{}, e error) {
	sc := i.(*gherkin.Scenario)
	rgex := regexp.MustCompile("[^0-9a-zA-Z]+")
	fileName := strings.ToLower(rgex.ReplaceAllString(sc.Name, "_"))
	screenshot, _ := Driver.Screenshot()

	support.SaveImage(screenshot, fileName)
})*/

//func InitializeScenario2(ctx *godog.ScenarioContext) {
//ctx.Step(`^hago el login con "([^"]*)" y "([^"]*)"$`, hagoElLoginConY)
//ctx.Step(`^estoy autenticado con éxito$`, estoyAutenticadoConXito)
//}
