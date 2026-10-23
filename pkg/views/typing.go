package views

import (
	"context"
	"fmt"
	"html"
	"io"
	"strings"

	"github.com/a-h/templ"
)

// typedText renders one fading span per character. It is written by hand
// because a templ for-loop puts whitespace between sibling spans, which
// spaced every letter apart.
func typedText(s string) templ.Component {
	return templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
		var b strings.Builder
		for i, c := range []rune(s) {
			fmt.Fprintf(&b, `<span class="type-ch" aria-hidden="true" style="animation-delay: %dms">%s</span>`,
				i*typeMs, html.EscapeString(string(c)))
		}
		_, err := io.WriteString(w, b.String())
		return err
	})
}
