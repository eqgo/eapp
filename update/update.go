package update

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Button is an update button
type Button struct {
	app.Compo
	ID        string // ID is the ID of the button
	Class     string // Class is the string-separated list of classes to apply to the button
	Text      string // Text is the text of the button
	available bool
}

// New makes a new update button with the given id, class, and text
func New(id, class, text string) *Button {
	return &Button{
		ID:    id,
		Class: class,
		Text:  text,
	}
}

// Render returns the UI of the update button
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

// OnClick is called when someone clicks on the update button
func (b *Button) OnClick(ctx app.Context, e app.Event) {
	ctx.Reload()
}

// OnAppUpdate is called when the app is ready for an update
func (b *Button) OnAppUpdate(ctx app.Context) {
	b.available = ctx.AppUpdateAvailable()
}
