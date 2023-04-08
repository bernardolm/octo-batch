package debug

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/k0kubun/pp"

	"github.com/bernardolm/octo-batch/config"
)

func Print(key string, value interface{}) {
	if !config.Debugging {
		return
	}

	color.NoColor = false

	fmt.Println()

	key = fmt.Sprintf(" 💬  %-*s", 50, key)

	color.New().
		Add(color.FgBlack).
		Add(color.Bold).
		Add(color.Italic).
		Add(color.Attribute(46)).
		Println(key)

	pp.Println(value)
}
