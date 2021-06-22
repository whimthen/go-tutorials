package widgets

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

var (
	scrollView = objc.Get("NSScrollView")
)

type NSScrollView struct {
	cocoa.NSView
}

func NewNSScrollView(frame core.NSRect) NSScrollView {
	return NSScrollView{cocoa.NSView{scrollView.Alloc().Send("initWithFrame:", frame)}}
}

func (s NSScrollView) SetContentView(clip NSClipView) {
	s.Set("contentView:", &clip)
}

func (s NSScrollView) SetHorizontalScroller(scroller NSScroller) {
	s.Set("horizontalScroller:", &scroller)
}

// SetBorderType A value that specifies the appearance of the scroll view’s border.
// https://developer.apple.com/documentation/appkit/nsscrollview/1403528-bordertype
func (s NSScrollView) SetBorderType(border NSBorderType) {
	s.Set("borderType:", core.NSUInteger(border))
}

func (s NSScrollView) BorderType() NSBorderType {
	return NSBorderType(s.Get("borderType").Uint())
}
