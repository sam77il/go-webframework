package sugar

import (
	"regexp"
	"strings"
)

var SugarJavaScriptLibrary string = `
window.addEventListener("DOMContentLoaded", () => {
	{{.Scripts}}
});
`

var FastLoadAnchors string = `
	document.addEventListener("click", function(e) {
		if (e.target.tagName === "A") {
			const href = e.target.getAttribute("href");
			if (href && href.startsWith("/") && !href.startsWith("//")) {
				e.preventDefault();
				window.history.pushState({}, '', href);
				loadContent(href)
			}
		}
	})

	window.addEventListener('popstate', () => {
		loadContent(window.location.pathname);
	});

	async function loadContent(path) {
		const res = await fetch(path, {headers: {"Sugar-Fast-Load": "true"}})
		const json = await res.json()
		document.getElementById('app').innerHTML = json.content
		if (json.metadata?.title) {
			document.title = json.metadata.title
		}
	}
`

var StateHandlerJS string = `
	class SugarState {
		constructor(name, initValue, element) {
			this.name = name;
			this.value = initValue;
			this.element = element;
			this.loadComponent();
		}

		loadComponent() {
			let clickableNodes = this.element.querySelectorAll('[sugar-onclick]');

			clickableNodes.forEach((item) => {
				const eventName = item.getAttribute("sugar-onclick");
				
				item.addEventListener("click", function(event) {
					const customEvent = new CustomEvent("sugar:" + eventName, {
                    	detail: { element: item }
                	});
                	document.dispatchEvent(customEvent);
				});
			})
		}

		setState(stateName, cb) {
			let value = cb(this.value);
			this.element.querySelector('[sugar-state="' + stateName + '"]').innerHTML = value;
			this.value = value
		}
	}
`

// Funktion zur Zusammenstellung des benötigten JavaScript
func ComposeJavaScript(content string) string {
	var jsModules []string

	if hasAnchors(content) {
		jsModules = append(jsModules, FastLoadAnchors)
	}

	jsModules = append(jsModules, StateHandlerJS)

	return strings.Join(jsModules, "\n")
}

func hasAnchors(content string) bool {
	// Einfache Regex zur Erkennung von <a href="/..."> Tags
	re := regexp.MustCompile(`<a\s+[^>]*href=['"]\/[^'"]*['"][^>]*>`)
	return re.MatchString(content)
}