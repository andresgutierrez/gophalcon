
package phalcon

import (
	"strings"
	//"fmt"
)

type Route struct {
    id int
    pattern string
    compiledPattern string
    paths map[string]int
}

func (this *Route) CompilePattern() (string, map[string]int) {

	var intermediate, parenthesesCount, marker int
	var bracketCount, foundPattern int
	//var looking_placeholder bool

	if (strings.Contains(this.pattern, "{")) {

		var notValid bool
		var numberMatches, i, j, regexpLength int
		var variable, regexp string
		var matches = make(map[string]int)

		route_str := ""

		for i = 0; i < len(this.pattern); i++ {

			ch := this.pattern[i]

			if (parenthesesCount == 0) {
				if (ch == '{') {
					if (bracketCount == 0) {
						marker = i + 1;
						intermediate = 0;
						notValid = false;
					}
					bracketCount++;
				} else {
					if (ch == '}') {
						bracketCount--;
						if (intermediate > 0) {
							if (bracketCount == 0) {

								numberMatches++;
								variable = ""
								regexp = ""
								item := this.pattern[marker:i];
								length := len(item);
								marker = i;

								for j = 0; j < length; j++ {
									ch = item[j]
									if (j == 0 && !((ch >= 'a' && ch <='z') || (ch >= 'A' && ch <='Z'))){
										notValid = true
										break
									}
									if ((ch >= 'a' && ch <='z') || (ch >= 'A' && ch <='Z') || (ch >= '0' && ch <='9') || ch == '-' || ch == '_' || ch ==  ':') {
										if (ch == ':') {
											regexpLength = length - j - 1;
											variable = this.pattern[marker:j]
											regexp = this.pattern[j:i+length]
											break;
										}
									} else {
										notValid = true
										break;
									}
								}

								if (!notValid) {
									{
										if (len(variable) > 0) {
											if (regexpLength > 0) {

												foundPattern = 0;
												for k := 0; k < regexpLength; k++ {
													if (foundPattern == 0) {
														if (regexp[k] == '(') {
															foundPattern = 1;
														}
													} else {
														if (regexp[k] == ')') {
															foundPattern = 2;
															break;
														}
													}
												}

												if (foundPattern != 2) {
													route_str += "(" + regexp + ")"
												} else {
													route_str += regexp
												}
												matches[variable] = numberMatches
											}
										} else {
											route_str += "([^/]*)"
											matches[item] = numberMatches
										}
									}
								} else {
									route_str += "{" + item + "}"
								}
								continue;
							}
						}
					}
				}

			}

			if (bracketCount == 0) {
				if (ch == '(') {
					parenthesesCount++;
				} else {
					if (ch == ')') {
						parenthesesCount--;
						if (parenthesesCount == 0) {
							numberMatches++;
						}
					}
				}
			}

			if (bracketCount > 0) {
				intermediate++;
			} else {
				route_str += string(ch)
			}
		}

		return "^" + route_str + "$", matches
	}
	return this.pattern, nil
}

func (this *Route) SetPattern(pattern string) *Route {
    this.pattern = pattern
    return this
}

func (this *Route) SetId(id int) *Route {
    this.id = id
    return this
}

func (this *Route) GetPattern() string {
    return this.pattern
}

func (this *Route) GetPaths() map[string]int {
    return this.paths
}

func (this *Route) GetCompiledPattern() string {
	if (this.compiledPattern == "") {
		this.compiledPattern, this.paths = this.CompilePattern()
	}
    return this.compiledPattern
}

func (this *Route) GetId() int {
    return this.id
}