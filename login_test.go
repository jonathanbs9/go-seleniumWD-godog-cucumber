package main

import (
	"fmt"

	"github.com/cucumber/godog"
	"github.com/jonathanbs9/go-seleniumWD-godog-cucumber/support"
	"github.com/tebeka/selenium"
)

var Driver selenium.WebDriver

func queAccedoALaPaginaPrincipal() error {
	Driver.Get("https://avalith.net/webmail")
	return nil
}

func hagoLoginCon(email, contraseña string) error {
	campoEmail, err := Driver.FindElement(selenium.ByID, "user")

	if err != nil {
		fmt.Errorf("Error  |  ", err.Error())
	}
	campoEmail.SendKeys(email)

	campoContraseña, err := Driver.FindElement(selenium.ByID, "pass")
	if err != nil {
		fmt.Errorf("Error  |  ", err.Error())
	}
	campoContraseña.SendKeys(contraseña)

	botonIniciarSesion, err := Driver.FindElement(selenium.ByID, "login_submit")
	if err != nil {
		fmt.Errorf("Error  |  ", err.Error())
	}
	botonIniciarSesion.Click()

	return nil
}

func autenticacionExitosa() (err error) {
	spanEmail, err := Driver.FindElement(selenium.ByClassName, "username")
	if err != nil {
		fmt.Errorf("Error  |  ", err.Error())
	}

	email, _ := spanEmail.Text()

	if email != "jonathans@avalith.net" {
		return fmt.Errorf("Error  al validar autenticacion de usuario! ")
	}

	return nil
}

func devuelvoElSiguienteMensaje(mensaje string) error {
	divAlerta, err := Driver.FindElement(selenium.ByClassName, "error-notice")
	if err != nil {
		return fmt.Errorf("Error al validar mensaje ! ")
	}

	alerta, _ := divAlerta.Text()
	if alerta != "El inicio de sesión no es válido." {
		return fmt.Errorf("Error  al validar texto del mensaje")
	}

	return nil
}

func FeatureContext(s *godog.ScenarioContext) {
	s.Step(`^Que accedo a la pagina principal$`, queAccedoALaPaginaPrincipal)
	s.Step(`^hago login con "([^"]*)" y "([^"]*)"$`, hagoLoginCon)
	s.Step(`^Estoy autenticado exitosamente$`, autenticacionExitosa)
	s.Step(`^Debo ver el siguiente mensaje "([^"]*)"$`, devuelvoElSiguienteMensaje)

	s.BeforeScenario(func(interface{}) {
		Driver = support.WDInit()
	})

	s.AfterScenario(func(i interface{}, e error) {
		Driver.Quit()
	})
}
