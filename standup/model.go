package standup

import "time"

// StandupUpdate represents a team member's daily standup update
type StandupUpdate struct {
	ID        int       `json:"id"`
	UserID    string    `json:"userId"`
	Update    string    `json:"update"`
	Blockers  string    `json:"blockers"`
	CreatedAt time.Time `json:"createdAt"`
}
