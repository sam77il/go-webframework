package pages

import (
	"sugarweb.dev/framework/components"
	"sugarweb.dev/framework/layouts"
	"sugarweb.dev/framework/sugar"
)

func RootPage() *sugar.Page {
	RootPage := sugar.Page{
		Metadata: sugar.Metadata{
			Title: "Home Page",
		},
		Layout: layouts.RootLayout(),
		Content: []*sugar.Component{
			components.Header("PageHeader", sugar.Props{}),
		},
	}
	return &RootPage
}