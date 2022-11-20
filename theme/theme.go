package theme

import (
	"log"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Button is a button to toggle the theme
type Button struct {
	app.Compo
	ID            string // ID is the ID of the button
	Class         string // Class is the string-separated list of classes to apply to the button
	LightModeText string // LightModeText is the text of the button when in light mode
	DarkModeText  string // DarkModeText is the text of the button when in dark mode
	darkMode      bool
}

// New makes a new theme button with the given id, classes, light mode text, and dark mode text
func New(id, class, lightModeText, darkModeText string) *Button {
	return &Button{
		ID:            id,
		Class:         class,
		LightModeText: lightModeText,
		DarkModeText:  darkModeText,
	}
}

// Render returns the UI of the theme button
func (b *Button) Render() app.UI {
	if b.darkMode {
		return app.Button().
			ID(b.ID).
			Class(b.Class).
			Text(b.DarkModeText).
			DataSet("theme", "dark").
			OnClick(func(ctx app.Context, e app.Event) {
				b.SwitchToLightMode()
			})
	}
	return app.Button().
		ID(b.ID).
		Class(b.Class).
		Text(b.LightModeText).
		DataSet("theme", "light").
		OnClick(func(ctx app.Context, e app.Event) {
			b.SwitchToDarkMode()
		})
}

// OnNav is called when a page with the theme button is navigated to
func (b *Button) OnNav(ctx app.Context) {
	b.Load()
	b.Apply()
}

// SwitchToLightMode switches the app to light mode
func (b *Button) SwitchToLightMode() {
	b.SetState(false)
	app.Window().Get("document").Get("documentElement").Get("dataset").Set("theme", "light")
	b.Save()
}

// SwitchToDarkMode switches the app to dark mode
func (b *Button) SwitchToDarkMode() {
	b.SetState(true)
	app.Window().Get("document").Get("documentElement").Get("dataset").Set("theme", "dark")
	b.Save()
}

// Save saves the theme button's state to local storage
func (b *Button) Save() {
	app.Window().Get("localStorage").Call("setItem", "useDarkMode", b.darkMode)
}

// Load loads the theme button's state from local storage, if it is saved there. Otherwise, it uses the operating system theme setting.
func (b *Button) Load() {
	darkMode := app.Window().Get("localStorage").Call("getItem", "useDarkMode")
	if darkMode.IsNull() {
		b.LoadFromOperatingSystem()
		return
	}
	darkModeBool, err := strconv.ParseBool(darkMode.String())
	if err != nil {
		log.Println(err)
		return
	}
	b.darkMode = darkModeBool
}

// LoadFromOperatingSystem loads the theme button's state from the operating system theme
func (b *Button) LoadFromOperatingSystem() {
	if app.Window().Call("matchMedia", "(prefers-color-scheme: dark)").Get("matches").Bool() {
		b.darkMode = true
	}
}

// Apply applies the theme button's current state to the app
func (b *Button) Apply() {
	if b.darkMode {
		b.SwitchToDarkMode()
		return
	}
	b.SwitchToLightMode()
}

// State returns the current state of the button (whether it is in dark mode)
func (b *Button) State() bool {
	return b.darkMode
}

// SetState sets the current state of the button to the given value (of whether it is in dark mode)
func (b *Button) SetState(darkMode bool) {
	b.darkMode = darkMode
}
