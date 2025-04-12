package components

import "sugarweb.dev/framework/sugar"

func HeaderItem(uniqueName string, props sugar.Props) *sugar.Component {
	HeaderC := sugar.NewComponent(uniqueName,
		[]sugar.State{}, `
	<div>Name: {{.Name}}, Value: {{.Value}}</div>
	`, ``, props)

	return HeaderC
}