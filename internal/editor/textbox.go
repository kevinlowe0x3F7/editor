package editor

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	gapbuffer "github.com/neelanjan00/gap-buffer"
)

type TextBox struct {
	widget.BaseWidget
	Text    *gapbuffer.GapBuffer
	focused bool
	cursor  *canvas.Rectangle
}

func NewTextBox() *TextBox {
	gb := new(gapbuffer.GapBuffer)
	gb.SetString(" ")

	t := &TextBox{
		Text:   gb,
		cursor: canvas.NewRectangle(theme.Color(theme.ColorNameForegroundOnPrimary)),
	}
	t.cursor.SetMinSize(fyne.NewSize(2, 20))
	t.ExtendBaseWidget(t)
	return t
}

func (t *TextBox) FocusGained() {
	fmt.Println("focus gained")
	t.focused = true
	t.Refresh()
}

func (t *TextBox) FocusLost() {
	t.focused = false
	t.Refresh()
}

func (t *TextBox) Focused() bool {
	return t.focused
}

// Tapped is called when the widget is clicked, giving it focus.
func (t *TextBox) Tapped(_ *fyne.PointEvent) {
	fmt.Println("I got tapped")
	fyne.CurrentApp().Driver().CanvasForObject(t).Focus(t) // This requests focus when clicked
}

// TypedKey handles key events when the widget is focused.
func (t *TextBox) TypedKey(event *fyne.KeyEvent) {
	switch event.Name {
	case fyne.KeyBackspace:
		if len(t.Text.GetString()) > 0 {
			t.Text.Backspace()
		}
	case fyne.KeyLeft:
		t.Text.MoveCursorLeft()
	case fyne.KeyRight:
		t.Text.MoveCursorRight()

	default:
		// Handle other special keys here if needed
	}
	t.Refresh()
}

// TypedRune handles typed characters when the widget is focused.
func (t *TextBox) TypedRune(r rune) {
	fmt.Printf("Key pressed: %c\n", r)
	t.Text.Insert(r)
	t.Refresh()
}

func (t *TextBox) CreateRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(t.Text.GetString(), theme.Color(theme.ColorNameForeground))
	text.Alignment = fyne.TextAlignLeading

	return &textBoxRenderer{
		textBox: t,
		text:    text,
		cursor:  t.cursor,
	}
}

type textBoxRenderer struct {
	textBox *TextBox
	text    *canvas.Text
	cursor  *canvas.Rectangle
}

func (r *textBoxRenderer) Layout(size fyne.Size) {
	r.text.Resize(size)

	// Calculate cursor position based on cursorPos index
	cursorX := r.text.MinSize().Width
	r.cursor.Move(fyne.NewPos(cursorX, 0))
}

func (r *textBoxRenderer) MinSize() fyne.Size {
	return fyne.NewSize(200, 200) // size of window
}

// Refresh refreshes the widget's state.
func (r *textBoxRenderer) Refresh() {
	r.text.Text = r.textBox.Text.GetString()
	if r.textBox.focused {
		r.text.Color = theme.Color(theme.ColorNameForegroundOnPrimary)
		r.cursor.Show()
	} else {
		r.text.Color = theme.Color(theme.ColorNameForeground)
		r.cursor.Hide()
	}
	canvas.Refresh(r.text)
}

// Objects returns the drawable objects in the renderer.
func (r *textBoxRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.text, r.cursor}
}

// Destroy cleans up any resources.
func (r *textBoxRenderer) Destroy() {}
