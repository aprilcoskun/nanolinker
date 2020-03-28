package db

import "errors"

const schema = `
CREATE TABLE IF NOT EXISTS users (
    username TEXT PRIMARY KEY,
	link_limit INT DEFAULT 10,
    password TEXT NOT NULL
);

CREATE TABLE  IF NOT EXISTS config (
    id INTEGER PRIMARY KEY CHECK (id = 1),
    admin_username TEXT DEFAULT 'admin',
    admin_password TEXT DEFAULT 'admin',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS links (
	id TEXT PRIMARY KEY,
	url TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	expires_at DATETIME DEFAULT NULL,
	deleted_at DATETIME DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS clicks (
    link_id TEXT NOT NULL,
    ip TEXT,
    referer TEXT,
    user_agent TEXT,
	clicked_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	
	CONSTRAINT fk_links
    FOREIGN KEY (link_id)
    REFERENCES links(id) ON DELETE CASCADE ON UPDATE CASCADE
);
`

const insertLinkQuery = `INSERT INTO links (id, url) VALUES ($1, $2)`
const updateLinkQuery = `UPDATE links SET id = $1, url = $2 WHERE id = $3`
const insertLinkWithExpireQuery = `INSERT INTO links (id, url, expires_at) VALUES ($1, $2, $3)`

const selectLinkQuery = `
SELECT id, url, expires_at FROM links 
WHERE id = $1 AND deleted_at is NULL AND (expires_at is NULL OR expires_at < CURRENT_TIMESTAMP)`

const selectLinksQuery = `
SELECT id, url, COALESCE(total_clicks, 0) AS total_clicks, COALESCE(unique_visitors, 0) as unique_visitors, created_at FROM links
LEFT JOIN (SELECT link_id, COUNT(*) AS total_clicks FROM clicks GROUP BY link_id)
C ON links.id = C.link_id
LEFT JOIN (SELECT link_id, COUNT(DISTINCT ip) AS unique_visitors FROM (SELECT DISTINCT link_id, ip FROM clicks GROUP BY link_id, ip) GROUP BY link_id)
U ON links.id = U.link_id
WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`

const insertClick = `INSERT INTO clicks (link_id, ip, referer, user_agent) VALUES ($1, $2, $3, $4)`
const selectConfig = `SELECT admin_username, admin_password, created_at, updated_at FROM config LIMIT 1`
const isConfigExistsQuery = `SELECT count(*) FROM config LIMIT 1`
const selectLinkCount = `SELECT count(*) FROM links`

const updateAdminNameQuery = `UPDATE config SET admin_username = $1`
const updateAdminPasswordQuery = `UPDATE config SET admin_password = $1`
const deleteLinkQuery = `UPDATE links SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1`

var ErrWrongPassword = errors.New("wrong password")
var ErrWrongUsername = errors.New("wrong username")
var ErrLinkNotFound = errors.New("url(s) not found")

var ErrNoRowsAffected = errors.New("no rows affected")
var ErrMultipleRowAffected = errors.New("multiple rows affected")
var ErrAlreadyConfigured = errors.New("already configured")
var ErrNotConfigured = errors.New("not configured")
