package sugar

import (
	"bytes"
	"errors"
	"html/template"
	"strings"
)

const coreLayout string = `
<!DOCTYPE html>
<html>
	<head>
		{{.Metadata}}
	</head>
	<body>
		<div id="app">
			{{.Layout}}
		</div>

		<script>{{.SugarJS}}</script>
	</body>
</html>
`

type Metadata struct {
	Title string
}

type Layout struct {
	Metadata Metadata
	Content []*Component
}

type Page struct {
	Metadata Metadata
	Layout  Layout
	Content []*Component
}

func (l Layout) LoadMetadata(metadata Metadata) string {
	var resultMetadata string

	if metadata.Title != "" {
		resultMetadata += "<title>" + metadata.Title + "</title>"
	}
	return resultMetadata
}

func (l Layout) LoadPage(page Page) (string, error) {
	SugarJavaScriptLibrary = l.LoadJavaScript(page.Content)

	loadedCoreLayout, err := l.LoadCoreLayout(page.Metadata)
	if err != nil {
		return "", errors.New(strings.ToUpper("Couldn't load page #1"))
	}
	finalPage, err := l.LoadLayout(loadedCoreLayout, page)
	if err != nil {
		return "", errors.New(strings.ToUpper("Couldn't load page #2"))
	}
	return finalPage, nil
}

func (l Layout) GetContentFromComponents(components []*Component, withScript bool) map[string]string {
	var componentContent string
	var componentScript string
	for _, value := range components {
		componentContent += value.Markup
		if withScript {
			componentScript += value.Script
		}
	}

	return map[string]string{
		"Content": componentContent,
		"Script": componentScript,
	}
}

func (l Layout) LoadJavaScript(pageContent []*Component) string {
	pageData := l.GetContentFromComponents(pageContent, true)
	layoutData := l.GetContentFromComponents(l.Content, true)

	fullMarkup := layoutData["Content"] + pageData["Content"]
	fullPage := ComposeJavaScript(fullMarkup) + pageData["Script"] + layoutData["Script"]

	dataCore := struct{
		Scripts template.HTML
	}{
		Scripts: template.HTML(fullPage),
	}
	pageScript, err := l.parseTemplate("javascript", SugarJavaScriptLibrary, dataCore)
	if err != nil {
		return ""
	}
	return pageScript
}

func (l Layout) LoadCoreLayout(pageMetadata Metadata) (string, error) {
	chosenMetadata := Metadata{}

	if pageMetadata.Title != "" {
		chosenMetadata.Title = pageMetadata.Title
	} else if l.Metadata.Title != "" {
		chosenMetadata.Title = l.Metadata.Title
	}
	layoutData := l.GetContentFromComponents(l.Content, false)

	dataCore := struct{
		Metadata template.HTML
		Layout template.HTML
		SugarJS template.JS
	}{
		Metadata: template.HTML(l.LoadMetadata(chosenMetadata)),
		Layout: template.HTML(layoutData["Content"]),
		SugarJS:     template.JS(SugarJavaScriptLibrary),
	}
	coreShi, err := l.parseTemplate("coreLayout", coreLayout, dataCore)
	if err != nil {
		return "", errors.New(strings.ToUpper("Couldn't load page #3"))
	}

	return coreShi, nil
}

func (l Layout) LoadLayout(loadedCoreLayout string, page Page) (string, error) {
	pageData := l.GetContentFromComponents(page.Content, false)
	dataLayout := struct{
		PageContent template.HTML
	}{
		PageContent: template.HTML(pageData["Content"]),
	}
	finalPage, err := l.parseTemplate("final", loadedCoreLayout, dataLayout)
	if err != nil {
		return "", errors.New(strings.ToUpper("Couldn't load page #4"))
	}
	return finalPage, nil
}

func (l Layout) FastLoadPage(page Page) (string, error) {
	var pageComponents string
	for _, value := range page.Content {
		pageComponents += value.Markup
	}
	dataLayout := struct{
		PageContent template.HTML
	}{
		PageContent: template.HTML(pageComponents),
	}
	layoutData := l.GetContentFromComponents(l.Content, false)

	layoutPageOnly, err := l.parseTemplate("layoutPageOnly", layoutData["Content"], dataLayout)
	if err != nil {
		return "", errors.New(strings.ToUpper("Couldn't load page #5"))
	}
	return layoutPageOnly, nil
}

func (l Layout) parseTemplate(templateName string, markup string, data any) (string, error) {
	tmpl, err := template.New(templateName).Parse(markup)
	if err != nil {
		return "", errors.New(strings.ToUpper("Couldn't load page #6"))
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", errors.New(strings.ToUpper("Couldn't load page #7"))
	}
	return buf.String(), nil
}