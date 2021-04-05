package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

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
	log.Println("Campo email : ", campoEmail, " campo contreaseña: ", campoContraseña)

	botonIniciarSesion, err := Driver.FindElement(selenium.ByID, "login_submit")
	if err != nil {
		fmt.Println("Error  |  ", err.Error())
	}
	botonIniciarSesion.Click()

	return nil
}

func estoyAutenticadoConXito() (err error) {
	spanEmail, err := Driver.FindElement(selenium.ByClassName, "username")
	if err != nil {
		fmt.Println("Error  |  ", err.Error())
	}

	email, _ := spanEmail.Text()

	if email != "jonathans@avalith.net" {
		return fmt.Errorf("Error  al validar autenticacion de usuario! ")
	}

	return nil
}

func deboVerElSiguienteMensaje(mensaje string) error {
	divAlerta, err := Driver.FindElement(selenium.ByClassName, "error-notice")
	if err != nil {
		return fmt.Errorf("Error al validar mensaje ! ")
	}

	alerta, _ := divAlerta.Text()
	if alerta != mensaje {
		return fmt.Errorf("Esperado: %v  | Obtenido: %v", mensaje, alerta)
	}

	return nil
}

func InitializeScenario(s *godog.ScenarioContext) {
	s.BeforeScenario(func(*godog.Scenario) {
		Driver = support.WDInit()
		Driver.Get("https://amanti.websitewelcome.com:2096/")

		s.Step(`^Que accedo a la pagina principal$`, queAccediALaPaginaPrincipal)
		s.Step(`^hago login con "([^"]*)" y "([^"]*)"$`, hagoElLoginConY)
		s.Step(`^Estoy autenticado exitosamente$`, estoyAutenticadoConXito)
		s.Step(`^Debo ver el siguiente mensaje "([^"]*)"$`, deboVerElSiguienteMensaje)
	})

	s.AfterScenario(func(s *godog.Scenario, err error) {
		//sc := s.(*gherkin.RuleTypeScenario.Name())
		rgex := regexp.MustCompile("[^0-9a-zA-Z]+")
		fileName := strings.ToLower(rgex.ReplaceAllString(s.Name, "_"))
		screenshot, _ := Driver.Screenshot()

		support.SaveImage(screenshot, fileName)

	})

}

/*s.AfterScenario(func(i interface{}, e error) {
	sc := i.(*gherkin.Scenario)
	rgex := regexp.MustCompile("[^0-9a-zA-Z]+")
	fileName := strings.ToLower(rgex.ReplaceAllString(sc.Name, "_"))
	screenshot, _ := Driver.Screenshot()

	support.SaveImage(screenshot, fileName)
})*/
