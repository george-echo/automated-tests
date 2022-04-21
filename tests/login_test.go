package tests

import (
	"testing"

	"github.com/george-echo/automated-tests/actions"
)

func TestLogin(t *testing.T) {
	ctx := actions.Setup(t)
	p := ctx.Login("sammysnail@echo.co.uk", "password")
	p.MustElementR("h1", "Medicine")
	p.MustScreenshot("MedicinePage.png")
}
