package app

import "testing"

type fakeOnboardingStore struct {
	sessions map[int64]UserSession
}

func newFakeOnboardingStore() *fakeOnboardingStore {
	return &fakeOnboardingStore{
		sessions: make(map[int64]UserSession),
	}
}

func (s *fakeOnboardingStore) SaveSession(userID int64, session UserSession) {
	s.sessions[userID] = session
}

func (s *fakeOnboardingStore) GetSession(userID int64) (UserSession, bool) {
	session, ok := s.sessions[userID]
	return session, ok
}

func TestHandleStartSetsAwaitingChannelState(t *testing.T) {
	svc := NewOnboardingService(NewInMemoryOnboardingStore())
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
	svc := NewOnboardingService(NewInMemoryOnboardingStore())
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
	svc := NewOnboardingService(NewInMemoryOnboardingStore())
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

func TestServiceWorksViaStoreInterface(t *testing.T) {
	var store OnboardingStore = newFakeOnboardingStore()
	svc := NewOnboardingService(store)
	userID := int64(404)
	channelURL := "https://youtube.com/@interface-check"

	_ = svc.HandleStart(userID)
	_ = svc.HandleText(userID, channelURL)

	session, ok := svc.GetSession(userID)
	if !ok {
		t.Fatalf("expected session to exist")
	}
	if session.State != StateOnboardingDone {
		t.Fatalf("expected state %q, got %q", StateOnboardingDone, session.State)
	}
	if session.ChannelURL != channelURL {
		t.Fatalf("expected channel url %q, got %q", channelURL, session.ChannelURL)
	}
}
