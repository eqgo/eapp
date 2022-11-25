// Package theme provides a customizable theme toggle button.
// The theme toggle button sets the data-color-scheme attribute of the root html element and the button.
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
	lightModeBody app.UI // lightModeBody is the body of the button when in light mode
	darkModeBody  app.UI // darkModeBody is the body of the button when in dark mode
	darkMode      bool
}

// New makes a new theme button with the given id, classes, light mode body, and dark mode body
func New(id, class string, lightModeBody, darkModeBody app.UI) *Button {
	return &Button{
		ID:            id,
		Class:         class,
		lightModeBody: lightModeBody,
		darkModeBody:  darkModeBody,
	}
}

// Render returns the UI of the theme button
func (b *Button) Render() app.UI {
	if b.darkMode {
		return app.Button().
			ID(b.ID).
			Class(b.Class).
			Body(b.darkModeBody).
			DataSet("color-scheme", "dark").
			OnClick(func(ctx app.Context, e app.Event) {
				b.SwitchToLightMode()
			})
	}
	return app.Button().
		ID(b.ID).
		Class(b.Class).
		Body(b.lightModeBody).
		DataSet("color-scheme", "light").
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
	app.Window().Get("document").Get("documentElement").Get("dataset").Set("colorScheme", "light")
	b.Save()
}

// SwitchToDarkMode switches the app to dark mode
func (b *Button) SwitchToDarkMode() {
	b.SetState(true)
	app.Window().Get("document").Get("documentElement").Get("dataset").Set("colorScheme", "dark")
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
