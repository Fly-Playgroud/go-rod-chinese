package rod

import (
	"context"
	"github.com/go-rod/rod/lib/proto"
)

// Frame implements these interfaces.
var (
	_ proto.Client      = &Frame{}
	_ proto.Contextable = &Frame{}
	_ proto.Sessionable = &Frame{}
)

type Frame struct {

	// frame or iframe's url
	url string

	// check if frame is already detached from main frame.
	detached bool

	// 和执行上下文有关，在 Frame 发生更新时会自动变更上下文中的 bindings 类似于 expose
	//worlds map[string]struct{}

	//frameManager *FrameManager

	// id of the frame's parent frame
	parentFrameId proto.PageFrameID

	//SessionID of the frame's sessionID
	sessionID proto.TargetSessionID

	// frame's id
	id proto.PageFrameID
	// frame's loaderId
	loaderId proto.NetworkLoaderID
	// record the lifecycleEvents of the frame,has issued
	lifecycleEvents map[string]struct{}
}

func (f *Frame) GetSessionID() proto.TargetSessionID {
	//TODO implement me
	panic("implement me")
}

func (f *Frame) GetContext() context.Context {
	//TODO implement me
	panic("implement me")
}

func (f *Frame) Call(ctx context.Context, sessionID, methodName string, params interface{}) (res []byte, err error) {
	//TODO implement me
	panic("implement me")
}

// UpdateId Updates the frame ID with the new ID. This happens when the main frame is
// replaced by a different frame.
func (f *Frame) updateId(frameId proto.PageFrameID) {
	f.id = frameId
}

func (f *Frame) updateSessionID(sessionID proto.TargetSessionID, keepWords bool) {

}

func (f *Frame) page() *Page {
	return nil
}
