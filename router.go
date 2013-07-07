package gooky

import (
	"strings"
//	"fmt"
	"regexp"
)

type Router struct {
	consecutive int
	routes []*Route
}

func (this *Router) Add(pattern string) *Route {
	route := &Route{}
	route.SetId(this.consecutive)
	this.consecutive++
	route.SetPattern(pattern)
	this.routes = append(this.routes, route)
	return route
}

func (this *Router) Handle(uri string) *RouterMatch {

	match := &RouterMatch {}
	match.SetWasMatched(false)

	for _, route := range this.routes {
		compiledPattern := route.GetCompiledPattern()
		if (strings.Contains(compiledPattern, "(")) {

			re := regexp.MustCompile(compiledPattern)
			res := re.FindStringSubmatch(uri)
			if res == nil {
				continue
			}

			match.SetWasMatched(true)
			match.SetMatchedRoute(route)

			params := make(map[string]string)
			for name, position := range route.GetPaths() {
				params[name] = res[position]
			}
			match.SetParams(params)

		} else {
			if (compiledPattern == uri) {
				match.SetWasMatched(true)
				match.SetMatchedRoute(route)
				break
			}
		}
	}

	return match
}
