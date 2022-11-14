package main

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"github.com/playwright-community/playwright-go"
	"log"
)

type webFeature struct {
	pw *playwright.Playwright
}

type godogsCtxKey struct{}

func (w *webFeature) iCanSeeTheLogo() error {
	return godog.ErrPending
}

func (w *webFeature) iClickPackagesOnTheMenu() error {
	return godog.ErrPending
}

func (w *webFeature) iEnterPackageNameInTheSearch(arg1 string) error {
	return godog.ErrPending
}

func (w *webFeature) iNavigateToByClickingPackagesOnMenu(arg1 string) error {
	return godog.ErrPending
}

func (w *webFeature) iPressSearchButton() error {
	return godog.ErrPending
}

func (w *webFeature) iShouldBeRedirectedTo(arg1 string) error {
	return godog.ErrPending
}

func (w *webFeature) iShouldSeeASearchPageWithPackage(arg1 string) error {
	return godog.ErrPending
}

func (w *webFeature) iVisit(ctx context.Context, url string) error {
	browser, ok := ctx.Value(godogsCtxKey{}).(playwright.Browser)
	if !ok {
		return errors.New("failed to get browser instance")
	}
	page, err := browser.NewPage()
	if err != nil {
		return err
	}
	_, err = page.Goto(url)
	if err != nil {
		return err
	}

	checkLogo, err := page.QuerySelector(".js-headerLogo.Header-logo")
	if err != nil {
		return err
	}
	if checkLogo == nil {
		return errors.New("failed to find logo")
	}
	return nil
}

func (w *webFeature) reset(ctx context.Context) (context.Context, error) {
	option := playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
		Channel:  playwright.String("chrome"),
	}

	browser, err := w.pw.Chromium.Launch(option)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, godogsCtxKey{}, browser), nil
}

func (w *webFeature) close() error {
	return w.pw.Stop()
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	feature := &webFeature{pw: pw}
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		ctx, err := feature.reset(ctx)
		if err != nil {
			return ctx, err
		}

		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		err = feature.close()
		if err != nil {
			return ctx, err
		}

		return ctx, nil
	})

	ctx.Step(`^I can see the logo$`, feature.iCanSeeTheLogo)
	ctx.Step(`^I click packages on the menu$`, feature.iClickPackagesOnTheMenu)
	ctx.Step(`^I enter "([^"]*)" package name in the search$`, feature.iEnterPackageNameInTheSearch)
	ctx.Step(`^I navigate to "([^"]*)" by clicking packages on menu$`, feature.iNavigateToByClickingPackagesOnMenu)
	ctx.Step(`^I press search button$`, feature.iPressSearchButton)
	ctx.Step(`^I should be redirected to "([^"]*)"$`, feature.iShouldBeRedirectedTo)
	ctx.Step(`^I should see a search page with "([^"]*)" package$`, feature.iShouldSeeASearchPageWithPackage)
	ctx.Step(`^I visit "([^"]*)"$`, feature.iVisit)
}
