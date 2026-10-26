// Command dmc-sandbox-go exists for the deps-merger GitHub App spike: its
// dependencies are outdated on purpose so Dependabot opens PRs.
package main

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func title(s string) string { return cases.Title(language.English).String(s) }

func main() {
	out, _ := yaml.Marshal(map[string]string{"id": uuid.NewString(), "name": title("sandbox")})
	fmt.Print(string(out))
}
