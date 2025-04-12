package sugar

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"reflect"
	"regexp"
	"strings"
)

type Props struct {
	Children *Component
	Data map[string]any
	State State
}

type Component struct {
	Markup string
	Props  Props
	Script string
}

type State struct {
	Name string
	InitValue any
}

func PageContentComponent() *Component {
	return &Component{
		Markup: "{{.PageContent}}",
		Props:  Props{},
		Script: "",
	}
}


func NewComponent(componentName string, states []State, markup string, script string, props Props) *Component {
	trimmedMarkup := strings.TrimSpace(markup)
	re := regexp.MustCompile(`^<([a-zA-Z][a-zA-Z0-9]*)([^>]*)>`)
	markup = re.ReplaceAllStringFunc(trimmedMarkup, func(match string) string {
        // Extrahiere Tag-Name und bestehende Attribute
        parts := re.FindStringSubmatch(match)
        if len(parts) >= 3 {
            tagName := parts[1]
            existingAttrs := parts[2]
            
            // Füge das neue Attribut hinzu
            return "<" + tagName + existingAttrs + " sugar-component=\"" + componentName + "\">"
        }
        return match
    })

	var buf bytes.Buffer

	data := map[string]any{}
		
	for _, v := range states {
		data["State_" + v.Name] = v.InitValue
		data["StateName_" + v.Name] = v.Name
		StateHandlerJS = strings.TrimSpace(StateHandlerJS)
		
		var initValueJS string
		
		if reflect.TypeOf(v.InitValue).Kind() == reflect.String {
			initValueString, _ := v.InitValue.(string)
			initValueJS = fmt.Sprintf("\"%s\"", initValueString)
		} else {
			initValueJS = fmt.Sprintf("%v", v.InitValue)
		}
			

		StateHandlerJS += `
	let ` + componentName + ` = new SugarState("` + componentName + `", ` + initValueJS + `, document.querySelector('[sugar-component="` + componentName + `"]'));
		`
	}
    
	if props.Children != nil && props.Children.Markup != "" {
		data["Children"] = template.HTML(props.Children.Markup)
	}

	fmt.Println(props.Data)

	// tmpl, err := template.New("component").Funcs().Parse(markup)
	tmpl, err := template.New("component").Parse(markup)
	if err != nil {
		log.Fatal("Error loading Component #1")
	}
		
	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Fatal("Error loading Component #2")
	}
	markup = buf.String()

	return &Component{
		Markup: markup,
		Props:  props,
		Script: script,
	}
}