package sugar

import (
	"encoding/json"
	"net/http"
)

type Routes struct {
	AvailableRoutes map[string]string
}

type RouteHandler struct {
	Writer  http.ResponseWriter
	Request *http.Request
	RouteData RouteData
}

type RouteData struct {
	FastLoad bool
}

func (r *RouteHandler) JSON(data any) *RouteHandler {
    r.Writer.Header().Set("Content-Type", "application/json")
    json.NewEncoder(r.Writer).Encode(data)
    return r
}

func (r *RouteHandler) Status(statusCode int) *RouteHandler {
	r.Writer.WriteHeader(statusCode)
	return r
}

func (r *RouteHandler) HTML(page *Page) *RouteHandler {
	if r.RouteData.FastLoad {
		pageData, err := page.Layout.FastLoadPage(*page)
		if err != nil {
			r.JSON(map[string]string{"error": "Couldn't load page #8"})
		}

		pageMetaData := map[string]string{}

		if page.Metadata.Title != "" {
			pageMetaData["title"] = page.Metadata.Title
		} else if page.Layout.Metadata.Title != "" {
			pageMetaData["title"] = page.Layout.Metadata.Title
		}

		pageMap := map[string]any{
			"content": pageData,
			"metadata": pageMetaData,
		}
		pageJson, err := json.Marshal(pageMap)
		if err != nil {
			r.JSON(map[string]string{"error": "Couldn't load page #9"})
		}
		r.Writer.Header().Set("Content-Type", "application/json")
		r.Writer.Write(pageJson)
	} else {
		finalPage, err := page.Layout.LoadPage(*page)
		if err != nil {
			r.JSON(map[string]string{"error": err.Error()})
		}
		
		r.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		r.Writer.Write([]byte(finalPage))
	}
	return r
}

func (routes *Routes) GET(route string, handler func(*RouteHandler)) {
	routes.AvailableRoutes[route] = "GET"

	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fastLoad := false
			if r.Header.Get("Sugar-Fast-Load") == "true" {
				fastLoad = true
			}
			ctx := RouteHandler{
				Writer: w,
				Request: r,
				RouteData: RouteData{
					FastLoad: fastLoad,
				},
			}
			handler(&ctx)
		}
	})
}