package ui

import (
	"chat-ui/core/models"
	"fmt"
	"github.com/jroimartin/gocui"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

var MessageSend chan MessageBox

var UserList = []models.User{
	{Name: "chen yi chao"},
	{Name: "zhou meng"},
}

func SaveMain(g *gocui.Gui, v *gocui.View) error {
	f, err := ioutil.TempFile("", "gocui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	p := make([]byte, 5)
	v.Rewind()
	for {
		n, err := v.Read(p)
		if n > 0 {
			if _, err := f.Write(p[:n]); err != nil {
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveVisualMain(g *gocui.Gui, v *gocui.View) error {
	f, err := ioutil.TempFile("", "gocui_demo_")
	if err != nil {
		return err
	}
	defer f.Close()

	vb := v.ViewBuffer()
	if _, err := io.Copy(f, strings.NewReader(vb)); err != nil {
		return err
	}
	return nil
}

func SendMessage(g *gocui.Gui, v *gocui.View) error {
	sideView, _ := g.View("side")
	_, y := sideView.Cursor()
	user := models.User{
		Name: UserList[y].Name,
	}
	message := v.Buffer()
	if len(strings.TrimSpace(message)) != 0 {
		MessageSend <- MessageBox{
			User:       user,
			MessageStr: v.Buffer(),
			Time:       time.Now(),
		}
	}
	g.Update(func(gui *gocui.Gui) error {
		v.Clear()
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)
		return nil
	})
	return nil
}

func Connect(g *gocui.Gui, v *gocui.View) error {
	MessageSend = make(chan MessageBox, 1)
	messageView, _ := g.View("main")
	g.SetViewOnTop("side")
	g.SetViewOnTop("main")
	g.SetViewOnTop("input")
	g.SetCurrentView("input")
	messageView.Clear()
	go func() {
		for {
			select {
			case sendMessage := <-MessageSend:
				g.Update(func(gui *gocui.Gui) error {
					fmt.Fprintf(messageView, AddMessageBox(sendMessage))
					return nil
				})
				break
			default:
			}
		}
	}()
	return nil
}
