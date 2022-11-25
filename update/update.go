// Package update provides a customizable update button
package update

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Button is an update button
type Button struct {
	app.Compo
	ID        string // ID is the ID of the button
	Class     string // Class is the string-separated list of classes to apply to the button
	body      app.UI // body is the body of the button
	available bool
}

// New makes a new update button with the given id, class, and body
func New(id, class string, body app.UI) *Button {
	return &Button{
		ID:    id,
		Class: class,
		body:  body,
	}
}

// Render returns the UI of the update button
func (b *Button) Render() app.UI {
	if b.available {
		return app.Button().
			ID(b.ID).
			Class(b.Class).
			Body(b.body).
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
