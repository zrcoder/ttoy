package pkg

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/zrcoder/rdor/pkg/style"
)

type Handler interface {
	tea.Model
	Name() string
	SetParent(tea.Model)
}

type Base struct {
	keyMap        *KeyMap
	name          string
	parent        tea.Model
	err           error
	keysHelp      help.Model
	width         int
	height        int
	Inputs        []textarea.Model
	focus         int
	ConvertAction func() error
}

func New(name string, editors int) *Base {
	b := &Base{name: name,
		keyMap:   getCommonKeys(),
		keysHelp: help.New(),
		Inputs:   make([]textarea.Model, editors),
	}
	for i := range b.Inputs {
		b.Inputs[i] = newTextarea()
	}
	b.Inputs[0].Focus()
	return b
}

func (b *Base) Name() string {
	return b.name
}

func (b *Base) FilterValue() string {
	return b.name
}

func (b *Base) SetParent(parent tea.Model) {
	b.parent = parent
}

func (b *Base) SetError(err error) {
	b.err = err
}

func (b *Base) Init() tea.Cmd {
	return textarea.Blink
}

func (b *Base) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		b.err = nil
		switch {
		case key.Matches(msg, b.keyMap.quit):
			return b, tea.Quit
		case key.Matches(msg, b.keyMap.home):
			return b.parent, nil
		case key.Matches(msg, b.keyMap.tab):
			b.Inputs[b.focus].Blur()
			b.focus = (b.focus + 1) % len(b.Inputs)
			cmd := b.Inputs[b.focus].Focus()
			cmds = append(cmds, cmd)
		case key.Matches(msg, b.keyMap.copy):
		case key.Matches(msg, b.keyMap.paste):
		default:
		}
		b.err = b.ConvertAction()
	case tea.WindowSizeMsg:
		b.height = msg.Height
		b.width = msg.Width
	}
	b.sizeInputs()
	for i := range b.Inputs {
		tmp, cmd := b.Inputs[i].Update(msg)
		b.Inputs[i] = tmp
		cmds = append(cmds, cmd)
	}
	return b, tea.Batch(cmds...)
}

func (b *Base) View() string {
	return lg.NewStyle().Padding(1, 3).Render(
		lg.JoinVertical(lg.Left,
			style.Title.Render(b.name),
			"",
			b.mainView(),
			"",
			b.keysHelp.View(b.keyMap),
		),
	)
}

func newTextarea() textarea.Model {
	t := textarea.New()
	t.Prompt = ""
	t.Placeholder = "Type something"
	t.ShowLineNumbers = false
	t.Cursor.Style = cursorStyle
	t.FocusedStyle.Placeholder = focusedPlaceholderStyle
	t.BlurredStyle.Placeholder = placeholderStyle
	t.FocusedStyle.CursorLine = cursorLineStyle
	t.FocusedStyle.Base = focusedBorderStyle
	t.BlurredStyle.Base = blurredBorderStyle
	t.FocusedStyle.EndOfBuffer = endOfBufferStyle
	t.BlurredStyle.EndOfBuffer = endOfBufferStyle
	t.KeyMap.DeleteWordBackward.SetEnabled(false)
	t.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	t.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))
	t.Blur()
	return t
}

func (b *Base) sizeInputs() {
	for i := range b.Inputs {
		b.Inputs[i].SetWidth((b.width - sepWidth) / len(b.Inputs))
		b.Inputs[i].SetHeight(b.height - errorHeight - helpHeight)
	}
}

func (b *Base) mainView() string {
	views := make([]string, len(b.Inputs))
	for i := range b.Inputs {
		views[i] = b.Inputs[i].View()
	}
	mainView := lg.JoinHorizontal(lg.Top, views...)
	if b.err != nil {
		return lg.JoinVertical(lg.Left,
			mainView,
			"",
			style.Error.Render(b.err.Error()),
		)
	}
	return mainView
}
