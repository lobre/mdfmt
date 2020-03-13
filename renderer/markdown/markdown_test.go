package markdown

import (
	"bytes"
    "strings"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
	"testing"
)

var tests = []struct{
    name string
    input string
    want string
} {
    {
        name: "heading",
        input: "##   This is my heading",
        want: "## This is my heading\n",
    }, {
        name: "paragraph",
        input: " This is a paragraph\nthat goes back to line",
        want: "This is a paragraph\nthat goes back to line\n",
    }, {
        name: "hard line break",
        input: "I want a hard break with double spaces  \nhere",
        want: "I want a hard break with double spaces  \nhere\n",
    },
}

func TestMarkdown(t *testing.T) {
    r := renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(NewRenderer(), 1100))) 
	md := goldmark.New(
		goldmark.WithRenderer(r),
	)

    for i, test := range tests {
        var buf bytes.Buffer
        if err := md.Convert([]byte(test.input), &buf); err != nil {
            t.Fatal(err)
        }

        got := buf.String()
        if test.want != got {
			t.Errorf("[%d] %s: got\n\t%v\nexpected\n\t%v", i, test.name, reveal(got), reveal(test.want))
        }
    }
}

func reveal(s string) string {
    return strings.Replace(s, " ", "·", -1)
}
