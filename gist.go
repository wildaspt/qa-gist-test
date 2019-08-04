package gist

import (
	"time"

	"github.com/DATA-DOG/godog"
)

func (b *BrowserSteps) iWriteDescription(text string) error {

	err1 := b.iWriteTo(text, "//*[@id='gists']/input", "xpath")
	if err1 != nil {
		return err1
	}

	return nil
}

func (b *BrowserSteps) iWriteFilename(text string) error {

	err1 := b.iWriteTo(text, "//*[@id='gists']/div[2]/div/div[1]/div[1]/input[2]", "xpath")
	if err1 != nil {
		return err1
	}

	return nil
}

func (b *BrowserSteps) iWriteContent(text string) error {

	element, err := b.GetWebDriver().FindElement("css selector", ".CodeMirror")
	if err != nil {
		return err
	}
	b.GetWebDriver().ExecuteScript("arguments[0].CodeMirror.setValue(\""+text+"\");", []interface{}{element})
	return nil
}

func (b *BrowserSteps) iGoToMyGists() error {

	err1 := b.iClick("//*[@id='user-links']/details/summary/span", "xpath")
	if err1 != nil {
		return err1
	}
	err2 := b.iClick("//*[@id='user-links']/details/details-menu/a[1]", "xpath")
	if err2 != nil {
		return err2
	}
	return nil
}

func (b *BrowserSteps) iSelectOneRecentGist() error {

	err1 := b.iClick("//*[@id='gist-pjax-container']/div[1]/div/div/ul/li[1]/div/a", "xpath")
	if err1 != nil {
		return err1
	}
	return nil
}

func (b *BrowserSteps) iClickEditGist() error {
	element, err := b.GetWebDriver().FindElement("css selector", "a[aria-label='Edit this Gist']")
	if err != nil {
		return err
	}

	return element.Click()
}

func (b *BrowserSteps) iClickDeleteGist() error {
	element, err := b.GetWebDriver().FindElement("xpath", "//*[@id='gist-pjax-container']/div[1]/div/div[1]/ul/li[2]/form/button")
	if err != nil {
		return err
	}

	return element.Click()

}

func (b *BrowserSteps) iConfirmTheDelete() error {
	u := time.Second
	time.Sleep(u * time.Duration(5))
	return b.GetWebDriver().AcceptAlert()

}

func (b *BrowserSteps) buildGistSteps(s *godog.Suite) {
	s.Step(`^I write the description as "([^"]*)"$`, b.iWriteDescription)
	s.Step(`^I write the filename as "([^"]*)"$`, b.iWriteFilename)
	s.Step(`^I write the content as "([^"]*)"$`, b.iWriteContent)
	s.Step(`^I go to my gists$`, b.iGoToMyGists)
	s.Step(`^I select one recent gist`, b.iSelectOneRecentGist)
	s.Step(`^I click edit gist`, b.iClickEditGist)
	s.Step(`^I click delete gist`, b.iClickDeleteGist)
	s.Step(`^I confirm the delete`, b.iConfirmTheDelete)
}
