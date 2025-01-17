package main

import (
	"fmt"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"github.com/whimthen/temp/macdriver-gui/widgets"
	"github.com/whimthen/temp/macdriver-gui/widgets/button"
	"runtime"
)

var (
	nsAlert = objc.Get("NSAlert")

	nsAlertController      objc.Object
	nsAlertControllerClass objc.Class
)

func init() {
	nsAlertControllerClass = objc.NewClass("DefaultNSAlertDelegate", "NSAlert")
	nsAlertControllerClass.AddMethod("beginSheetModalForWindow:", func(alert, handler objc.Object) {
		println(alert)
	})
	objc.RegisterClass(nsAlertControllerClass)
	nsAlertController = objc.Get("DefaultNSAlertDelegate").Alloc().Init()
}

type NSAlert struct {
	objc.Object
}

func NewNSAlert() NSAlert {
	alert := NSAlert{nsAlert.Alloc().Init()}
	alert.SetDelegate(nsAlertController)
	return alert
}

func (i NSAlert) SetDelegate(delegate objc.Object) {
	i.Send("setDelegate:", delegate)
}

func (i NSAlert) Delegate() objc.Object {
	return i.Send("delegate")
}

func (i NSAlert) MessageText() string {
	return i.Get("messageText").String()
}

func (i NSAlert) SetMessageText(s string) {
	i.Set("messageText:", core.String(s))
}

func (i NSAlert) InformativeText() string {
	return i.Get("informativeText").String()
}

func (i NSAlert) SetInformativeText(s string) {
	i.Set("informativeText:", core.String(s))
}

func (i NSAlert) AddButtonWithTitle(s string) {
	i.Send("addButtonWithTitle:", core.String(s))
}

func (i NSAlert) BeginSheetModalForWindow(win cocoa.NSWindow) objc.Object {
	return i.Send("beginSheetModalForWindow:completionHandler:", &win, nil)
}

func main() {
	runtime.LockOSThread()
	app := cocoa.NSApp_WithDidLaunch(wenAppLaunch)
	app.Retain()

	itemQuit := cocoa.NSMenuItem_Init("Quit", objc.Sel("terminate:"), "q")
	menu := cocoa.NSMenu_Init("MenuInit")
	menu.AddItem(itemQuit)
	app.SetMainMenu(menu)

	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}

func wenAppLaunch(notification objc.Object) {
	window := cocoa.NSWindow_Init(
		core.Rect(0, 0, 600, 665),
		cocoa.NSClosableWindowMask|
			cocoa.NSResizableWindowMask|
			cocoa.NSMiniaturizableWindowMask|
			cocoa.NSFullSizeContentViewWindowMask|
			cocoa.NSTitledWindowMask,
		cocoa.NSBackingStoreBuffered,
		false,
	)
	window.SetHasShadow(true)

	rect := core.Rect(0, 0, 600, 665)
	rootView := cocoa.NSView_Init(rect)
	// rootView.Set("setTranslatesAutoresizingMaskIntoConstraints:", false)

	subView := cocoa.NSView_Init(rect)

	rootView.Send("addSubview:", &subView)
	topConstraint := widgets.NewNSLayoutConstraint()
	// topConstraint.SetConstraintWithItem(subView, widgets.NSLayoutConstraintAttributeTop, widgets.NSLayoutConstraintRelationEqual, rootView, widgets.NSLayoutConstraintAttributeBottom, 1.0, 40)
	rootView.Send("addConstraints:", core.NSArray_WithObjects(topConstraint))
	// rootView.AddSubviewPositionedRelativeTo(subView, 3, rootView)

	nsButton := button.NSButton{NSView: cocoa.NSView_Init(core.Rect(0, 0, 0, 0))}
	nsButton.Set("title:", core.String("titled button"))
	subView.Send("addSubview:", &nsButton)

	btn1 := button.NewButtonWithFrame(core.Rect(100, 150, 200, 22))
	btn1.SetTitle("Change Title")
	btn1.SetBorderType(button.BezelStyleRounded)
	btn1.SetType()
	i := 0
	btn1.SetAction(func(b button.NSButton) {
		// rect := core.NSRect{
		//	Origin: core.NSPoint{100, 200},
		//	Size:   core.NSSize{400, 25},
		// }
		i++
		btn1.SetTitle(fmt.Sprintf("Changed Title With Action %d", i))
		// btn1.Set("frame:", rect)
	})
	subView.Send("addSubview:", &btn1)

	btn := button.NewButtonWithFrame(core.Rect(100, 100, 200, 22))
	btn.SetTitle("Show Alert with sheet modal")
	btn.SetBorderType(button.BezelStyleRounded)
	btn.SetType()
	btn.SetAction(func(b button.NSButton) {
		showAlertWithSheet(window)
	})
	subView.Send("addSubview:", &btn)

	disclosure := button.NewButtonWithFrame(core.Rect(100, 50, 200, 22))
	disclosure.SetTitle("")
	disclosure.SetBorderType(button.BezelStyleDisclosure)
	disclosure.SetType()
	disclosure.SetAction(func(b button.NSButton) {
		// state := objc.Get("NSControl.StateValue")
		// obj.Set("state:", state)
	})
	subView.Send("addSubview:", &disclosure)

	window.Send("setMinSize:", core.NSSize{
		Width:  300,
		Height: 300,
	})
	window.SetTitle("Test sheet modal alert")
	window.SetContentView(rootView)
	window.MakeKeyAndOrderFront(rootView)
	window.Center()
}

func showAlertWithSheet(window cocoa.NSWindow) {
	nsAlert := NewNSAlert()
	nsAlert.SetMessageText("Alert test sheet message")
	nsAlert.SetInformativeText("Detailed description of nsAlert message")
	nsAlert.AddButtonWithTitle("OK")
	nsAlert.AddButtonWithTitle("Second")
	nsAlert.BeginSheetModalForWindow(window)
}
