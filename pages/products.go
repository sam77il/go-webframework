package pages

import (
	"sugarweb.dev/framework/components"
	"sugarweb.dev/framework/layouts"
	"sugarweb.dev/framework/sugar"
)

func ProductsPage() *sugar.Page {
	ProductsPage := sugar.Page{
		Metadata: sugar.Metadata{
			Title: "Products Page",
		},
		Layout: layouts.RootLayout(),
		Content: []*sugar.Component{
			components.Header(
				"ProductsHeader",
				sugar.Props{
					Children: components.Nav(sugar.Props{}),
				},
			),
		},
	}
	return &ProductsPage
}