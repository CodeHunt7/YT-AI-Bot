package app

import "strings"

type UserState string

const (
	StateIdle            UserState = "idle"
	StateAwaitingChannel UserState = "awaiting_channel_url"
	StateOnboardingDone  UserState = "completed"
)

type UserSession struct {
	State      UserState
	ChannelURL string
}

type OnboardingStore interface {
	SaveSession(userID int64, session UserSession)
	GetSession(userID int64) (UserSession, bool)
}

type InMemoryOnboardingStore struct {
	sessions map[int64]UserSession
}

func NewInMemoryOnboardingStore() *InMemoryOnboardingStore {
	return &InMemoryOnboardingStore{
		sessions: make(map[int64]UserSession),
	}
}

func (s *InMemoryOnboardingStore) SaveSession(userID int64, session UserSession) {
	s.sessions[userID] = session
}

func (s *InMemoryOnboardingStore) GetSession(userID int64) (UserSession, bool) {
	session, ok := s.sessions[userID]
	return session, ok
}

type OnboardingService struct {
	store OnboardingStore
}

func NewOnboardingService(store OnboardingStore) *OnboardingService {
	return &OnboardingService{
		store: store,
	}
}

func (s *OnboardingService) HandleStart(userID int64) string {
	s.store.SaveSession(userID, UserSession{State: StateAwaitingChannel})
	return "РџСЂРёРІРµС‚. РЇ РР-РїСЂРѕРґСЋСЃРµСЂ РґР»СЏ YouTube-Р°РІС‚РѕСЂРѕРІ. РћС‚РїСЂР°РІСЊ СЃСЃС‹Р»РєСѓ РЅР° СЃРІРѕР№ YouTube-РєР°РЅР°Р», С‡С‚РѕР±С‹ РЅР°С‡Р°С‚СЊ РѕРЅР±РѕСЂРґРёРЅРі."
}

func (s *OnboardingService) HandleText(userID int64, text string) string {
	session, _ := s.store.GetSession(userID)

	switch session.State {
	case StateAwaitingChannel:
		session.ChannelURL = strings.TrimSpace(text)
		session.State = StateOnboardingDone
		s.store.SaveSession(userID, session)
		return "РЎСЃС‹Р»РєР° РЅР° РєР°РЅР°Р» РїРѕР»СѓС‡РµРЅР°. РЎР»РµРґСѓСЋС‰РёР№ С€Р°Рі вЂ” Р°РЅР°Р»РёР· РєР°РЅР°Р»Р°."
	default:
		return "РќР°РїРёС€Рё /start, С‡С‚РѕР±С‹ РЅР°С‡Р°С‚СЊ РѕРЅР±РѕСЂРґРёРЅРі."
	}
}

func (s *OnboardingService) GetSession(userID int64) (UserSession, bool) {
	return s.store.GetSession(userID)
}
