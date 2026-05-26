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

type OnboardingService struct {
	sessions map[int64]UserSession
}

func NewOnboardingService() *OnboardingService {
	return &OnboardingService{
		sessions: make(map[int64]UserSession),
	}
}

func (s *OnboardingService) HandleStart(userID int64) string {
	s.sessions[userID] = UserSession{State: StateAwaitingChannel}
	return "Привет. Я ИИ-продюсер для YouTube-авторов. Отправь ссылку на свой YouTube-канал, чтобы начать онбординг."
}

func (s *OnboardingService) HandleText(userID int64, text string) string {
	session := s.sessions[userID]

	switch session.State {
	case StateAwaitingChannel:
		session.ChannelURL = strings.TrimSpace(text)
		session.State = StateOnboardingDone
		s.sessions[userID] = session
		return "Ссылка на канал получена. Следующий шаг — анализ канала."
	default:
		return "Напиши /start, чтобы начать онбординг."
	}
}

func (s *OnboardingService) GetSession(userID int64) (UserSession, bool) {
	session, ok := s.sessions[userID]
	return session, ok
}
