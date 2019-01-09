package footnote

import (
	"strconv"

	"github.com/vito/booklit"
)

func init() {
	booklit.RegisterPlugin("footnote", NewPlugin)
}

func NewPlugin(section *booklit.Section) booklit.Plugin {
	return &Plugin{section: section}
}

type Plugin struct {
	section *booklit.Section
	notes   []booklit.Content
}

func (plugin *Plugin) Footnote(note booklit.Content) booklit.Content {
	plugin.notes = append(plugin.notes, note)
	return booklit.Styled{
		Style:   booklit.Style("footnote"),
		Content: booklit.String(strconv.Itoa(len(plugin.notes) - 1)),
	}
}

func (plugin *Plugin) Footnotemark() booklit.Content {
	body := booklit.Styled{
		Style:   booklit.Style("footnotes"),
		Content: booklit.Sequence(plugin.notes),
	}
	plugin.notes = nil
	return body
}
