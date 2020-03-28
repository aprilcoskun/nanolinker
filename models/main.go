package models

import (
	"database/sql"
	"time"
)

type Config struct {
	Username       string `db:"admin_username" json:"username"`
	HashedPassword string `db:"admin_password" json:"-"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Link struct {
	ID          string       `db:"id" json:"id"`
	Url         string       `db:"url" json:"url"`
	TotalClicks int          `db:"total_clicks" json:"total_clicks"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	ExpiredAt   sql.NullTime `db:"expires_at" json:"expires_at"`
	UpdatedAt   time.Time    `db:"updated_at" json:"updated_at"`
}

type LinkWithStats struct {
	ID             string    `db:"id" json:"id"`
	Url            string    `db:"url" json:"url"`
	TotalClicks    int       `db:"total_clicks" json:"total_clicks"`
	UniqueVisitors int       `db:"unique_visitors" json:"unique_visitors"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}

type Click struct {
	LinkID    string    `db:"link_id" json:"link_id"`
	Ip        string    `db:"ip" json:"ip"`
	Referer   string    `db:"referer" json:"referer"`
	UserAgent string    `db:"user_agent" json:"user_agent"`
	ClickedAt time.Time `db:"clicked_at" json:"clicked_at"`
}

type ConfigureUserData struct {
	Username   string `json:"username" binding:"required,gte=3,alphanumunicode"`
	Password   string `json:"password" binding:"required,gte=6"`
	RememberMe bool   `json:"remember_me"`
}

type UpdateConfigData struct {
	Username string `db:"admin_username" json:"username"`
	Password string `db:"admin_password" json:"password"`
}

type Stats struct {
	TotalClicks    int    `db:"total_clicks" json:"total_clicks"`
	UniqueVisitors string `db:"unique_visitors" json:"unique_visitors"`
}

type CachedLink struct {
	ID        string       `db:"id" json:"id"`
	Url       string       `db:"url" json:"url"`
	ExpiredAt sql.NullTime `db:"expires_at" json:"expires_at"`
}
