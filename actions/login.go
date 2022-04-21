package actions

import (
	"os"
	"time"

	"github.com/go-rod/rod"
)

func (ctx *Ctx) Login(email string, password string) *rod.Page {
	p := ctx.page("https://web.staging.echo.co.uk/login")
	googleLogin(p)
	p.MustElementR("button", "Use your password instead").MustClick()
	e := p.MustElement("input#email")
	time.Sleep(time.Millisecond * 100)
	e.MustInput(email)
	ps := p.MustElement("input#password")
	time.Sleep(time.Millisecond * 100)
	ps.MustInput(password)
	p.MustElementR("button", "Log in").MustClick()
	return p
}

func googleLogin(p *rod.Page) {
	p.MustElement("input#identifierId.whsOnd.zHQkBf").MustInput("george.songhurst@echo.co.uk")
	time.Sleep(time.Millisecond * 100)
	p.MustElement("button.VfPpkd-LgbsSe.VfPpkd-LgbsSe-OWXEXe-k8QpJ.VfPpkd-LgbsSe-OWXEXe-dgl2Hf.nCP5yc.AjY5Oe.DuMIQc.qfvgSe.qIypjc.TrZEUc.lw1w4b").MustClick()
	time.Sleep(time.Second)
	p.MustElement("input.whsOnd.zHQkBf").MustInput(os.Getenv("password"))
	time.Sleep(time.Millisecond * 100)
	p.MustElement("button.VfPpkd-LgbsSe.VfPpkd-LgbsSe-OWXEXe-k8QpJ.VfPpkd-LgbsSe-OWXEXe-dgl2Hf.nCP5yc.AjY5Oe.DuMIQc.qfvgSe.qIypjc.TrZEUc.lw1w4b").MustClick()
}
