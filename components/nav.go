package components

import "sugarweb.dev/framework/sugar"

func Nav(props sugar.Props) *sugar.Component {
	HeaderC := sugar.NewComponent("Nav",
	[]sugar.State{},
	`
	<nav>
		<ul>
			<li><a href="/">Home</a></li>
			<li><a href="/products">Products</a></li>
		</ul>
	</nav>
	`,
	``,
	sugar.Props{
		Children: nil,
	})

	return HeaderC
}