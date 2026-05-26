package app

import "testing"

func TestHandleStartSetsAwaitingChannelState(t *testing.T) {
	svc := NewOnboardingService()
	userID := int64(101)

	_ = svc.HandleStart(userID)

	session, ok := svc.GetSession(userID)
	if !ok {
		t.Fatalf("expected session to be created")
	}
	if session.State != StateAwaitingChannel {
		t.Fatalf("expected state %q, got %q", StateAwaitingChannel, session.State)
	}
}

func TestTextAfterStartSavesChannelURL(t *testing.T) {
	svc := NewOnboardingService()
	userID := int64(202)
	channelURL := "https://youtube.com/@mychannel"

	_ = svc.HandleStart(userID)
	_ = svc.HandleText(userID, channelURL)

	session, ok := svc.GetSession(userID)
	if !ok {
		t.Fatalf("expected session to exist")
	}
	if session.ChannelURL != channelURL {
		t.Fatalf("expected channel url %q, got %q", channelURL, session.ChannelURL)
	}
}

func TestStateChangesToCompletedAfterChannelURL(t *testing.T) {
	svc := NewOnboardingService()
	userID := int64(303)

	_ = svc.HandleStart(userID)
	_ = svc.HandleText(userID, "https://youtube.com/@another")

	session, ok := svc.GetSession(userID)
	if !ok {
		t.Fatalf("expected session to exist")
	}
	if session.State != StateOnboardingDone {
		t.Fatalf("expected state %q, got %q", StateOnboardingDone, session.State)
	}
}
