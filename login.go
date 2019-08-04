package gist

import "github.com/DATA-DOG/godog"

var username string

func (b *BrowserSteps) iLoginWithEmailPassword(email, pass string) error {

	err1 := b.iWriteTo(email, "login_field", "id")
	if err1 != nil {
		return err1
	}

	err2 := b.iWriteTo(pass, "password", "id")
	if err2 != nil {
		return err2
	}

	err3 := b.iClick("//*[@id='login']/form/div[3]/input[4]", "xpath")
	if err3 != nil {
		return err3
	}

	return nil
}

func (b *BrowserSteps) iGetTheUserName() error {
	err1 := b.iClick("/html/body/div[1]/header/div[8]/details/summary/span", "xpath")
	if err1 != nil {
		return err1
	}

	element, err := b.GetWebDriver().FindElement("xpath", "/html/body/div[1]/header/div[8]/details/details-menu/div[1]/a/strong")
	if err != nil {
		return err
	}

	username, _ = element.Text()
	return nil
}

func (b *BrowserSteps) buildLoginSteps(s *godog.Suite) {
	s.Step(`^I login with "([^"]*)" as email and "([^"]*)" as password$`, b.iLoginWithEmailPassword)
	s.Step(`^I get the username$`, b.iGetTheUserName)
}
