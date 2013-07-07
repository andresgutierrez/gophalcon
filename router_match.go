package gooky

type RouterMatch struct {
	wasMatched bool
	matchedRoute *Route
	params map[string]string
}

func (this *RouterMatch) WasMatched() bool {
	return this.wasMatched
}

func (this *RouterMatch) SetWasMatched(wasMatched bool) *RouterMatch {
	this.wasMatched = wasMatched
	return this
}

func (this *RouterMatch) GetMatchedRoute() *Route {
	return this.matchedRoute
}

func (this *RouterMatch) SetMatchedRoute(matchedRoute *Route) *RouterMatch {
	this.matchedRoute = matchedRoute
	return this
}

func (this *RouterMatch) SetParams(params map[string]string) *RouterMatch {
	this.params = params
	return this
}

func (this *RouterMatch) GetParams() (map[string]string) {
	return this.params
}
