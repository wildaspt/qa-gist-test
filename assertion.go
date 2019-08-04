package gist

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

func (b *BrowserSteps) iShouldSeeElement(selector, by string) error {
	element, _ := b.GetWebDriver().FindElement(by, selector)

	if element != nil {
		return nil
	} else {
		return fmt.Errorf("Element is not in page")
	}
}

func (b *BrowserSteps) iShouldSeeGistFilename(filename string) error {
	element, err := b.GetWebDriver().FindElement("xpath", "//*[@id='gist-pjax-container']/div[1]/div/div[1]/h1/strong/a")
	if err != nil {
		return err
	}

	filetext, _ := element.Text()
	if filetext == filename {
		return nil
	} else {
		return fmt.Errorf("Gist is not in the page")
	}
}

func (b *BrowserSteps) iShouldSeeListOfGist() error {
	element, _ := b.GetWebDriver().FindElement("xpath", "//*[@id='gist-pjax-container']/div/div[1]/div[2]")
	if element != nil {
		return nil
	} else {
		return fmt.Errorf("Element is not in page")
	}
}

func (b *BrowserSteps) buildAssertionSteps(s *godog.Suite) {
	s.Step(`^I should see gist with filename "([^"]*)"$`, b.iShouldSeeGistFilename)
	s.Step(`^I should see list of gist`, b.iShouldSeeListOfGist)
}
