package apimaster

type Handler struct {
	URL           string
	AuthHeaderKey string
	APIKeys       []APIKey
}

func NewHandler(url string, authHeaderKey string, apiKeys []APIKey) Handler {
	return Handler{
		URL:           url,
		AuthHeaderKey: authHeaderKey,
		APIKeys:       apiKeys,
	}
}

type APIKey struct {
	Value          string
	RPSLimit       int
	DailyLimit     int
	DailyLimitUsed int
}

func NewAPIKey(value string, rpsLimit int, dailyLimit int, dailyLimitUsed int) APIKey {
	return APIKey{
		Value:          value,
		RPSLimit:       rpsLimit,
		DailyLimit:     dailyLimit,
		DailyLimitUsed: dailyLimitUsed,
	}
}

func (h *Handler) GetAvailableKeys() APIKey {

}

func (k *APIKey) IncrementDailyLimitUsed(usage int) {
	k.DailyLimitUsed += usage
}

func (k *APIKey) CanUseKey() bool {
	return k.DailyLimitUsed >= k.DailyLimit
}
