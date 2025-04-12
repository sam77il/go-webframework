package layouts

import (
	"sugarweb.dev/framework/components"
	"sugarweb.dev/framework/sugar"
)

func RootLayout() sugar.Layout {
	RootLayout := sugar.Layout{
	Metadata: sugar.Metadata{
		Title: "Main Layout",
	},
	Content: []*sugar.Component{
			components.Header(
				"LayoutHeader",
				sugar.Props{
					Children: components.Nav(sugar.Props{}),
				},
			),
			sugar.PageContentComponent(),
		},
	}
	return RootLayout
}