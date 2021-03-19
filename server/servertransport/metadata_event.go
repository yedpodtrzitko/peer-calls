package servertransport

import (
	"fmt"

	"github.com/peer-calls/peer-calls/server/transport"
)

type metadataEvent struct {
	// Type must always be set
	Type metadataEventType `json:"type"`

	// Track will be set only when Type is metadataEventTypeTrackEvent.
	Track *trackEventJSON `json:"trackEvent"`
}

type metadataEventType int

const (
	// Track event contains the information about tracks.
	metadataEventTypeTrack metadataEventType = iota + 1
)

func (m metadataEventType) String() string {
	switch m {
	case metadataEventTypeTrack:
		return "TrackEvent"
	default:
		return fmt.Sprintf("Unknown(%d)", m)
	}
}

type initEventJSON struct {
	ClientID string
}

// trackEventJSON is used instead of TrackEvent because JSON cannot deserialize
// to Track interface, so a SimpleTrack is used.
type trackEventJSON struct {
	ClientID  string
	TrackInfo trackInfoJSON
	Type      transport.TrackEventType
}

func newTrackEventJSON(trackEvent transport.TrackEvent) trackEventJSON {
	// TODO watch out for possible panics.
	track := trackEvent.TrackWithMID.Track.(transport.SimpleTrack)

	return trackEventJSON{
		ClientID: trackEvent.ClientID,
		TrackInfo: trackInfoJSON{
			Track: track,
			Mid:   trackEvent.TrackWithMID.Mid,
		},
		Type: trackEvent.Type,
	}
}

// trackEvent converts the trackEventJSON to TrackEvent.
func (t trackEventJSON) trackEvent(clientID string) transport.TrackEvent {
	return transport.TrackEvent{
		ClientID: clientID,
		TrackWithMID: transport.TrackInfo{
			Track: t.TrackInfo.Track,
			Mid:   t.TrackInfo.Mid,
		},
		Type: t.Type,
	}
}

type trackInfoJSON struct {
	Track transport.SimpleTrack
	Mid   string
}

type byeEventJSON struct{}
