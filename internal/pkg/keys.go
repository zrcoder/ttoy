package pkg

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	copy, paste, home, quit, tab key.Binding
}

func (k *KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.copy, k.paste, k.tab, k.home, k.quit}
}

func (k *KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.copy, k.paste, k.tab, k.home, k.quit}}
}

func getCommonKeys() *KeyMap {
	copy := key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "copy"),
	)
	paste := key.NewBinding(
		key.WithKeys("ctrl+v"),
		key.WithHelp("ctrl+v", "paste"),
	)
	tab := key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch text area"),
	)
	home := key.NewBinding(
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "back to home"),
	)
	quit := key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "quit"),
	)
	return &KeyMap{
		copy:  copy,
		paste: paste,
		tab:   tab,
		home:  home,
		quit:  quit,
	}
}
