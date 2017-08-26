package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/klauern/shellchecker/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
