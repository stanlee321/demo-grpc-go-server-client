package main

import (
	"github.com/jmoiron/sqlx"
)

// ProfileStateOrganization is the struct resposible to represent the database entity
type ProfileStateOrganization struct {
	ID                   int32  `bson:"id" db:"id"`
	ProfileName          string `bson:"profile_name" db:"profile_name"`
	ProfileSigla         string `bson:"profile_sigla" db:"profile_sigla"`
	ProfilePoliticalType string `bson:"profile_political_type" db:"profile_political_type"`
	ProfileFase          string `bson:"profile_fase" db:"profile_fase"`
	ProfileTwitterID     int32  `bson:"profile_twitter_id" db:"profile_twitter_id"`
	ProfileFacebookID    int32  `bson:"profile_facebook_id" db:"profile_facebook_id"`
	ProfileInstagramID   int32  `bson:"profile_instagram_id" db:"profile_instagram_id"`
	ProfileYoutubeID     int32  `bson:"profile_youtube_id" db:"profile_youtube_id"`
	ProfileTikTokID      int32  `bson:"profile_tiktok_id" db:"profile_tiktok_id"`
	CreatedAt            string `bson:"created_at" db:"created_at"`
	UpdatedAt            string `bson:"updated_at" db:"updated_at"`
}

// getByUserName returns the quer SELECT by Email
func (u *ProfileStateOrganization) getProfileStateOrganizationByID(db *sqlx.DB) error {
	return db.Get(u, "SELECT * FROM ProfileStateOrganization WHERE Id=$1", u.ID)
}

// Update the data in the db using the instance values
func (u *ProfileStateOrganization) updateProfileStateOrganization(db *sqlx.DB) error {

	_, err := db.Exec(`UPDATE ProfileStateOrganization SET 
										Profile_Name				=$1,
										Profile_Sigla				=$2,
										Profile_Political_Type		=$3,
										Profile_Fase				=$4,
										Profile_Twitter_Id			=$5, 
										Profile_Facebook_Id			=$6,
										Profile_Instagram_Id		=$7,
										Profile_Youtube_Id			=$8,
										Profile_TikTok_Id			=$9
			WHERE Id=$10`,
		u.ProfileName,
		u.ProfileSigla,
		u.ProfilePoliticalType,
		u.ProfileFase,
		u.ProfileTwitterID,
		u.ProfileFacebookID,
		u.ProfileInstagramID,
		u.ProfileYoutubeID,
		u.ProfileTikTokID,
		u.ID)
	return err
}

// Delete the date from the db using the instance values
func (u *ProfileStateOrganization) deleteProfileStateOrganization(db *sqlx.DB) error {

	// create query
	_, err := db.Exec(`DELETE FROM ProfileStateOrganization WHERE Id=$1`, u.ID)

	return err
}

// Create a new user in the db using the instance values
func (u *ProfileStateOrganization) createProfileStateOrganization(db *sqlx.DB) error {

	return db.QueryRow(
		`INSERT INTO ProfileStateOrganization(

			Profile_Name			,
			Profile_Sigla			,
			Profile_Political_Type	,
			Profile_Fase			,
			Profile_Twitter_Id		, 
			Profile_Facebook_Id		,
			Profile_Instagram_Id	,
			Profile_Youtube_Id		,
			Profile_TikTok_Id		
		) 
	
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING Id`,

		u.ProfileName,
		u.ProfileSigla,
		u.ProfilePoliticalType,
		u.ProfileFase,
		u.ProfileTwitterID,
		u.ProfileFacebookID,
		u.ProfileInstagramID,
		u.ProfileYoutubeID,
		u.ProfileTikTokID,
	).Scan(&u.ID)
}

// List return a list of users. Could be applied pagination
func listProfileStateOrganization(db *sqlx.DB, start, count int) ([]ProfileStateOrganization, error) {
	users := []ProfileStateOrganization{}
	err := db.Select(&users, `SELECT * FROM ProfileStateOrganization LIMIT $1 OFFSET $2`, count, start)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// List return a list of users. Could be applied pagination
func listProfileStateOrganizationByType(db *sqlx.DB, politicalType string, start int, count int) ([]ProfileStateOrganization, error) {
	users := []ProfileStateOrganization{}
	err := db.Select(&users, `SELECT * FROM ProfileStateOrganization WHERE Profile_Political_Type=$1 LIMIT $2 OFFSET $3`, politicalType, count, start)
	if err != nil {
		return nil, err
	}
	return users, nil
}
