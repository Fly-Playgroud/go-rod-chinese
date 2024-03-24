package rod

import "github.com/go-rod/rod/lib/proto"

// FrameTree Keeps track of the page frame tree and it's is managed by
// FrameManager. FrameTree uses frame IDs to reference frame and it
// means that referenced frames might not be in the tree anymore. Thus, the tree
// structure is eventually consistent.
type FrameTree struct {
	frames    map[proto.PageFrameID]*Frame
	parentIds map[proto.PageFrameID]proto.PageFrameID
	childIds  map[proto.PageFrameID]map[proto.PageFrameID]struct{}
	mainFrame *Frame
}

// getMainFrame returns the frame's main frame
func (ft *FrameTree) getMainFrame() *Frame {
	return ft.mainFrame
}

// getById returns the frame by its id
func (ft *FrameTree) getById(frameId proto.PageFrameID) *Frame {
	return ft.frames[frameId]
}

// allFrames returns all frames
func (ft *FrameTree) allFrames() []*Frame {
	var frames []*Frame
	for _, frame := range ft.frames {
		frames = append(frames, frame)
	}
	return frames
}

// addFrame add a frame into frame tree
// TODO: how to make sure that the frame is added into the tree.
func (ft *FrameTree) addFrame(frame *Frame) {
	ft.frames[frame.id] = frame

	if frame.parentFrameId != "" {
		ft.parentIds[frame.id] = frame.parentFrameId

		if _, ok := ft.childIds[frame.parentFrameId]; !ok {
			ft.childIds[frame.parentFrameId] = make(map[proto.PageFrameID]struct{})
		}

		ft.childIds[frame.parentFrameId][frame.id] = struct{}{}
	}
}

// removeFrame a frame from frame tree
func (ft *FrameTree) removeFrame(frame *Frame) {
	delete(ft.frames, frame.id)
	delete(ft.parentIds, frame.id)

	if frame.parentFrameId != "" {
		if childIds, ok := ft.childIds[frame.parentFrameId]; ok {
			delete(childIds, frame.id)
		}
	} else {
		ft.mainFrame = nil
	}
}

// childFrames get frame's all children
func (ft *FrameTree) childFrames(frameId proto.PageFrameID) []*Frame {
	childIds, ok := ft.childIds[frameId]
	if !ok {
		return nil
	}

	var children []*Frame
	for childId := range childIds {
		if frame := ft.getById(childId); frame != nil {
			children = append(children, frame)
		}
	}

	return children
}

// parentFrame return frame's parent
func (ft *FrameTree) parentFrame(frameId proto.PageFrameID) *Frame {

	parentId, ok := ft.parentIds[frameId]
	if !ok {
		return nil
	}

	return ft.getById(parentId)
}
