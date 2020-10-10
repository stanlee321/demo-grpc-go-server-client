package main

import (
	"github.com/jmoiron/sqlx"
)

// ProfilePoliticalParty is the struct resposible to represent the database entity
type ProfilePoliticalParty struct {
	ID                          int32  `bson:"id" db:"id"`
	ProfilePoliticalName        string `bson:"profile_political_name" db:"profile_political_name"`
	ProfilePoliticalSigla       string `bson:"profile_political_sigla" db:"profile_political_sigla"`
	ProfilePoliticalType        string `bson:"profile_political_type" db:"profile_political_type"`
	ProfilePoliticalTwitterID   int32  `bson:"profile_political_twitter_id" db:"profile_political_twitter_id"`
	ProfilePoliticalFacebookID  int32  `bson:"profile_political_facebook_id" db:"profile_political_facebook_id"`
	ProfilePoliticalInstagramID int32  `bson:"profile_political_instagram_id" db:"profile_political_instagram_id"`
	ProfilePoliticalYoutubeID   int32  `bson:"profile_political_youtube_id" db:"profile_political_youtube_id"`
	ProfilePoliticalTikTokID    int32  `bson:"profile_political_tiktok_id" db:"profile_political_tiktok_id"`
	CreatedAt                   string `bson:"created_at" db:"created_at"`
	UpdatedAt                   string `bson:"updated_at" db:"updated_at"`
}

func (u *ProfilePoliticalParty) getPoliticalPartyByID(db *sqlx.DB) error {
	return db.Get(u, `SELECT * FROM ProfilePoliticalParty WHERE Id=$1`, u.ID)
}

// Update the data in the db using the instance values
func (u *ProfilePoliticalParty) updatePolitcalParty(db *sqlx.DB) error {

	_, err := db.Exec(`UPDATE ProfilePoliticalParty SET 
										Profile_Political_Name				= $1,
										Profile_Political_Sigla				= $2,
										Profile_Political_Type 				= $3,
										Profile_Political_Twitter_Id		= $4, 
										Profile_Political_Facebook_Id		= $5,
										Profile_Political_Instagram_Id		= $6,
										Profile_Political_Youtube_Id		= $7,
										Profile_Political_TikTok_Id			= $8
			WHERE Id=$9`,
		u.ProfilePoliticalName,
		u.ProfilePoliticalSigla,
		u.ProfilePoliticalType,
		u.ProfilePoliticalTwitterID,
		u.ProfilePoliticalFacebookID,
		u.ProfilePoliticalInstagramID,
		u.ProfilePoliticalYoutubeID,
		u.ProfilePoliticalTikTokID,
		u.ID)
	return err
}

// Delete the date from the db using the instance values
func (u *ProfilePoliticalParty) deletePoliticalParty(db *sqlx.DB) error {

	// create query
	_, err := db.Exec(`DELETE FROM ProfilePoliticalParty WHERE Id=$1`, u.ID)

	return err
}

// Create a new user in the db using the instance values
func (u *ProfilePoliticalParty) createPoliticalParty(db *sqlx.DB) error {

	return db.QueryRow(
		`INSERT INTO ProfilePoliticalParty(
			Profile_Political_Name,
			Profile_Political_Sigla,
			Profile_Political_Type,
			Profile_Political_Twitter_Id,
			Profile_Political_Facebook_Id,
			Profile_Political_Instagram_Id,
			Profile_Political_Youtube_Id,
			Profile_Political_TikTok_Id
		) 
	
		VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING Id`,
		u.ProfilePoliticalName,
		u.ProfilePoliticalSigla,
		u.ProfilePoliticalType,
		u.ProfilePoliticalTwitterID,
		u.ProfilePoliticalFacebookID,
		u.ProfilePoliticalInstagramID,
		u.ProfilePoliticalYoutubeID,
		u.ProfilePoliticalTikTokID,
	).Scan(&u.ID)
}

// List return a list of users. Could be applied pagination
func listPoliticalParty(db *sqlx.DB, start int, count int) ([]ProfilePoliticalParty, error) {
	users := []ProfilePoliticalParty{}
	err := db.Select(&users, `SELECT * FROM ProfilePoliticalParty LIMIT $1 OFFSET $2`, count, start)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// List return a list of users. Could be applied pagination
func listtPoliticalPartyByType(db *sqlx.DB, profileType string, start int, count int) ([]ProfilePoliticalParty, error) {
	users := []ProfilePoliticalParty{}
	err := db.Select(&users, `SELECT * FROM ProfilePoliticalParty WHERE Profile_Political_Type=$1 LIMIT $2 OFFSET $3`, profileType, count, start)
	if err != nil {
		return nil, err
	}
	return users, nil
}
