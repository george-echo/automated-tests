package actions

import (
	"testing"

	"github.com/go-rod/rod"
	"github.com/ysmood/got"
)

type Ctx struct {
	g got.G

	browser *rod.Browser
}

// Setup for tests
var Setup = func() func(t *testing.T) Ctx {
	browser := rod.New().MustConnect()

	return func(t *testing.T) Ctx {
		t.Parallel() // run each test concurrently

		return Ctx{got.New(t), browser}
	}
}()

func (ctx Ctx) page(url string) *rod.Page {
	page := ctx.browser.MustIncognito().MustPage(url)
	ctx.g.Cleanup(page.MustClose)
	return page
}
