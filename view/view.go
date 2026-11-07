package view

import (
	"github.com/ELY-Open-Source/CrychicUI/navigation"
	"github.com/egoist/mygo/ui"
)

func tabs(c *ui.Context, activeTab *int, tabs ...string) {
	navigation.UnderlineTabs(c, "Format", tabs, activeTab) //.Margin(18)
}
