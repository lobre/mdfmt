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
        want: "## This is my heading",
    }, {
        name: "paragraph",
        input: " This is a paragraph\nthat goes back to line",
        want: "This is a paragraph\nthat goes back to line",
    }, {
        name: "two paragraphs",
        input: "This is a first paragraph\n\n\nThis is a second paragraph",
        want: "This is a first paragraph\n\nThis is a second paragraph",
    }, {
        name: "hard line break",
        input: "I want a hard break with double spaces  \nhere",
        want: "I want a hard break with double spaces  \nhere",
    }, {
        name: "simple list",
        input: "This is a list:\n- first item\n- second item",
        want: "This is a list:\n\n- first item\n- second item",
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
            t.Errorf("[%d] %s:\nINPUT:\n%v\n\nGOT:\n%v\n\nWANT:\n%v\n\n", i, test.name, reveal(test.input), reveal(got), reveal(test.want))
        }
    }
}

func reveal(s string) string {
    s = strings.Replace(s, " ", "·", -1)
    s = strings.Replace(s, "\n", "$\n", -1)
    s += "$"
    return s
}
