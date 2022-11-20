package install

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Button is an install button
type Button struct {
	app.Compo
	ID        string // ID is the ID of the button
	Class     string // Class is the string-separated list of classes to apply to the button
	Text      string // Text is the text of the button
	available bool
}

// New makes a new install button with the given id, class, and text
func New(id, class, text string) *Button {
	return &Button{
		ID:    id,
		Class: class,
		Text:  text,
	}
}

// Render returns the UI of the install button
func (b *Button) Render() app.UI {
	if b.available {
		return app.Button().
			ID(b.ID).
			Class(b.Class).
			Text(b.Text).
			OnClick(b.OnClick)
	}
	return app.Text("")
}

// OnMount is called when the install button is mounted
func (b *Button) OnMount(ctx app.Context) {
	b.available = ctx.IsAppInstallable()
}

// OnClick is called when someone clicks on the install button
func (b *Button) OnClick(ctx app.Context, e app.Event) {
	ctx.ShowAppInstallPrompt()
}

// OnAppInstallChange is called when the installability of the app changes
func (b *Button) OnAppInstallChange(ctx app.Context) {
	b.available = ctx.IsAppInstallable()
}
