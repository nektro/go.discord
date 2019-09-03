package discord

type GuildRole struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Color       int    `json:"color"`
	Hoist       bool   `json:"hoist"`
	Position    int    `json:"position"`
	Permissions int    `json:"permissions"`
	Managed     bool   `json:"managed"`
	Mentionable bool   `json:"mentionable"`
}

type GuildMember struct {
	User         User     `json:"user"`
	Nickname     string   `json:"nick"`
	Roles        []string `json:"roles"`
	JoinedAt     string   `json:"joined_at"`
	PremiumSince string   `json:"premium_since"`
	Deaf         bool     `json:"deaf"`
	Mute         bool     `json:"mute"`
}

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
	MFAEnabled    bool   `json:"mfa_enabled"`
	Locale        string `json:"locale"`
	Verified      bool   `json:"verified"`
	Email         bool   `json:"email"`
	Flags         int    `json:"flags"`
	PremiumType   int    `json:"premium_type"`
}
