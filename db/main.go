package db

import (
	"database/sql"
	"github.com/aprilcoskun/nanolinker/models"
	"github.com/aprilcoskun/nanolinker/utils/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/teris-io/shortid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var db *sqlx.DB

func init() {
	db = sqlx.MustOpen("sqlite3", "nanolinker.db")
	db.MustExec(schema)
	linkCache = &Cache{items: make(map[string]*models.CachedLink)}

	logger.Info("Database Initialized")
}

func IsConfigured() bool {
	var id int
	err := db.Get(&id, isConfigExistsQuery)
	if err != nil || id != 1 {
		return false
	}

	return true
}

func Configure(configData *models.ConfigureUserData) (err error) {
	if IsConfigured() {
		err = ErrAlreadyConfigured
		return
	}

	// Hash passwords with bcrypt(default cost = 10) algorithm before storing them in database.
	hashedP, err := bcrypt.GenerateFromPassword([]byte(configData.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	_, err = db.Exec(`INSERT INTO config (id, admin_username, admin_password) VALUES (1, $1, $2)`, configData.Username, hashedP)

	return
}

func SaveLink(link models.CachedLink) error {
	var result sql.Result
	var err error

	// Generate ID
	if link.ID == "" {
		link.ID, err = shortid.Generate()
		if err != nil {
			return err
		}
	}

	// Set with expire date if it's exists
	if link.ExpiredAt.Valid && time.Now().Before(link.ExpiredAt.Time) {
		result, err = db.Exec(insertLinkWithExpireQuery, link.ID, link.Url, link.ExpiredAt.Time)
	} else {
		result, err = db.Exec(insertLinkQuery, link.ID, link.Url)
	}

	if err != nil {
		return err
	}
	linkCache.Set(&link)
	return isSingleRowAffected(result)
}

func InsertClick(click *models.Click) (err error) {
	_, err = db.Exec(insertClick, click.LinkID, click.Ip, click.Referer, click.UserAgent)
	if err != nil {
		logger.Error(err)
	}

	return
}

func GetLinks(limit, offset int) (links []models.Link, count int, err error) {
	err = db.Get(&count, selectLinkCount)
	if err != nil {
		logger.Error(err)
		return
	}

	err = db.Select(&links, selectLinksQuery, limit, offset)
	if err == sql.ErrNoRows {
		err = ErrLinkNotFound
	}
	return
}

func GetLink(id string) (link models.CachedLink, err error) {
	cachedLink, found := linkCache.Get(id)
	if found {
		link = *cachedLink
		return
	}
	err = db.Get(&link, selectLinkQuery, id)
	if err == sql.ErrNoRows {
		err = ErrLinkNotFound
		return
	}

	linkCache.Set(&link)

	return
}

func DeleteLink(id string) error {
	result, err := db.Exec(deleteLinkQuery, id)
	if err != nil {
		return err
	}

	return isSingleRowAffected(result)
}

func AuthUser(configData *models.ConfigureUserData) error {
	var config models.Config
	err := db.Get(&config, selectConfig)
	if err != nil {
		logger.Info(err)
		return ErrNotConfigured
	}

	if config.Username != configData.Username {
		return ErrWrongUsername
	}

	if config.HashedPassword == "admin" && configData.Password == "admin" {
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(config.HashedPassword), []byte(configData.Password))
	if err != nil {
		return ErrWrongPassword
	}

	return nil
}

// Check if only one row is affected when dealing with user operations in database
func isSingleRowAffected(result sql.Result) error {
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error(err)
		return err
	}

	if rowsAffected == 0 {
		return ErrNoRowsAffected
	} else if rowsAffected != 1 {
		logger.Warn(rowsAffected, "Rows affected ??")
		return ErrMultipleRowAffected
	}
	return nil
}

// TODO: UPDATE CONFIG AND UPDATE Link
