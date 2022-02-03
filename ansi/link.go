package ansi

import (
	"io"
	"net/url"
)

// A LinkElement is used to render hyperlinks.
type LinkElement struct {
	Text    string
	BaseURL string
	URL     string
	Child   ElementRenderer // FIXME
}

func (e *LinkElement) Render(w io.Writer, ctx RenderContext) error {
	var textRendered bool
	/*
		if node.LastChild != nil {
			if node.LastChild.Type == bf.Image {
				el := tr.NewElement(node.LastChild)
				err := el.Renderer.Render(w, node.LastChild, tr)
				if err != nil {
					return err
				}
			}
			if len(node.LastChild.Literal) > 0 &&
				string(node.LastChild.Literal) != string(node.LinkData.Destination) {
				textRendered = true
				el := &BaseElement{
					Token: string(node.LastChild.Literal),
					Style: ctx.style[LinkText],
				}
				err := el.Render(w, node.LastChild, tr)
				if err != nil {
					return err
				}
			}
		}
	*/
	if len(e.Text) > 0 && e.Text != e.URL {
		textRendered = true
	}

	u, err := url.Parse(e.URL)
	if err == nil &&
		"#"+u.Fragment != e.URL { // if the URL only consists of an anchor, ignore it
		style := ctx.options.Styles.Link
		if !textRendered {
			style.BlockPrefix = ""
			style.BlockSuffix = ""
		}

		var url = resolveRelativeURL(e.BaseURL, e.URL)
		el := &BaseElement{
			Token:  url,
			Prefix: "\u001b]8;;",
			Suffix: "\u001b\\",
			Style:  style,
		}
		err := el.Render(w, ctx)
		if err != nil {
			return err
		}
	}

	if len(e.Text) > 0 && e.Text != e.URL {
		style := ctx.options.Styles.LinkText

		el := &BaseElement{
			Token: e.Text,
			Style: style,
			Suffix: "\u001b]8;;\u001b\\",
		}
		err := el.Render(w, ctx)
		if err != nil {
			return err
		}
	}


	return nil
}
