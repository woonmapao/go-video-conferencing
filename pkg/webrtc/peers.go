package webrtc

import (
	"go-video-conferencing/pkg/chat"
	"sync"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

type Peers struct {
	Listlock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

func (p *Peers) DispatchKeyFrame() {

}
