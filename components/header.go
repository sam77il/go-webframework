package components

import (
	"sugarweb.dev/framework/sugar"
)

func Header(uniqueName string, props sugar.Props) *sugar.Component {
	componentData := map[string]any{
		"HeaderItems": map[string]int{
			"Samil": 16,
			"Ben": 16,
			"Furkan": 21,
			"Hoplit": 22,
		},
		"Component": map[string]func(string, sugar.Props) *sugar.Component{
			"HeaderItem": HeaderItem,
		},
	}
	props.Data = componentData
	HeaderC := sugar.NewComponent(uniqueName,
		[]sugar.State{
			{
				Name: "count",
				InitValue: 1,
			},
		}, `
	<header>
		<h1>Current Count: <span sugar-state="{{.StateName_count}}">{{.State_count}}</span></h1>
		<button sugar-onclick="testFunc">Click Me!</button>

		{{range $key, $value := .HeaderItems}}
			{{component $key $value)}}
		{{end}}

		{{.Children}}
	</header>
	`, `
	document.addEventListener("sugar:testFunc", function(e) {
		` + uniqueName + `.setState("count", (curVal) => {
			return curVal += 1;
		});
	});
	`, props)

	return HeaderC
}