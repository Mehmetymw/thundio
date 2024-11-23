package mqtt

import "log"

type PermissionChecker struct {
	allowedTopics map[string][]string // Kullanıcı bazında izinli topic listesi
}

func NewPermissionChecker() *PermissionChecker {
	return &PermissionChecker{
		allowedTopics: map[string][]string{
			"device-1": {"devices/device-1/data", "devices/device-1/commands"},
			"device-2": {"devices/device-2/data"},
		},
	}
}

func (p *PermissionChecker) CanPublish(clientID, topic string) bool {
	topics, ok := p.allowedTopics[clientID]
	if !ok {
		log.Printf("Client %s has no allowed topics", clientID)
		return false
	}
	for _, t := range topics {
		if t == topic {
			return true
		}
	}
	log.Printf("Client %s is not allowed to publish to topic %s", clientID, topic)
	return false
}

func (p *PermissionChecker) CanSubscribe(clientID, topic string) bool {
	topics, ok := p.allowedTopics[clientID]
	if !ok {
		log.Printf("Client %s has no allowed topics", clientID)
		return false
	}
	for _, t := range topics {
		if t == topic {
			return true
		}
	}
	log.Printf("Client %s is not allowed to subscribe to topic %s", clientID, topic)
	return false
}
