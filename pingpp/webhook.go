package pingpp

func ParseWebhooks(webhooks []byte) (*Event, error) {
	var event Event
	if webhooks != nil && len(webhooks) > 0 {
		err := JsonDecode(webhooks, &event)
		if err != nil {
			return nil, err
		}
	}
	return &event, nil
}
