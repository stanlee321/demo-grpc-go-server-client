package main

import (
	"github.com/jmoiron/sqlx"
)

// ProfilePeople is the struct resposible to represent the database entity
type ProfilePeople struct {
	ID                 int32  `bson:"id" db:"id"`
	ProfileFirstName   string `bson:"profile_firstname" db:"profile_firstname"`
	ProfileLastName    string `bson:"profile_lastname" db:"profile_lastname"`
	ProfileOrientation string `bson:"profile_orientation" db:"profile_orientation"`
	ProfileTwitterID   int32  `bson:"profile_twitter_id" db:"profile_twitter_id"`
	ProfileFacebookID  int32  `bson:"profile_facebook_id" db:"profile_facebook_id"`
	ProfileInstagramID int32  `bson:"profile_instagram_id" db:"profile_instagram_id"`
	ProfileYoutubeID   int32  `bson:"profile_youtube_id" db:"profile_youtube_id"`
	ProfileTikTokID    int32  `bson:"profile_tiktok_id" db:"profile_tiktok_id"`
	ProfileKind        string `bson:"profile_kind" db:"profile_kind"`
	CreatedAt          string `bson:"created_at" db:"created_at"`
	UpdatedAt          string `bson:"updated_at" db:"updated_at"`
}

// getByUserName returns the quer SELECT by Email
func (u *ProfilePeople) getProfilePeopleByUserID(db *sqlx.DB) error {
	return db.Get(u, "SELECT * FROM ProfilePeople WHERE Id=$1", u.ID)
}

// Update the data in the db using the instance values
func (u *ProfilePeople) updateProfilePeople(db *sqlx.DB) error {

	_, err := db.Exec(`UPDATE ProfilePeople SET 
										Profile_FirstName		=$1,
										Profile_LastName		=$2,
										Profile_Orientation		=$3,
										Profile_Twitter_Id		=$4, 
										Profile_Facebook_Id		=$5,
										Profile_Instagram_Id	=$6,
										Profile_Youtube_Id		=$7,
										Profile_TikTok_Id		=$8,
										Profile_Kind			=$9
			WHERE Id=$10`,
		u.ProfileFirstName,
		u.ProfileLastName,
		u.ProfileOrientation,
		u.ProfileTwitterID,
		u.ProfileFacebookID,
		u.ProfileInstagramID,
		u.ProfileYoutubeID,
		u.ProfileTikTokID,
		u.ProfileKind,
		u.ID)

	return err
}

// Delete the date from the db using the instance values
func (u *ProfilePeople) deleteProfilePeople(db *sqlx.DB) error {

	// create query
	_, err := db.Exec(`DELETE FROM ProfilePeople WHERE Id=$1`, u.ID)

	return err
}

// Create a new user in the db using the instance values
func (u *ProfilePeople) createProfilePeople(db *sqlx.DB) error {

	return db.QueryRow(

		`INSERT INTO ProfilePeople(
			Profile_FirstName,
			Profile_LastName,
			Profile_Orientation,
			Profile_Twitter_Id, 
			Profile_Facebook_Id	,
			Profile_Instagram_Id,
			Profile_Youtube_Id,
			Profile_TikTok_Id,
			Profile_Kind
		) 
	
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING Id`,
		u.ProfileFirstName,
		u.ProfileLastName,
		u.ProfileOrientation,
		u.ProfileTwitterID,
		u.ProfileFacebookID,
		u.ProfileInstagramID,
		u.ProfileYoutubeID,
		u.ProfileTikTokID,
		u.ProfileKind,
	).Scan(&u.ID)
}

// List return a list of users. Could be applied pagination
func listProfilePeople(db *sqlx.DB, start, count int) ([]ProfilePeople, error) {
	users := []ProfilePeople{}
	err := db.Select(&users, `SELECT * FROM ProfilePeople LIMIT $1 OFFSET $2`, count, start)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// List return a list of users. Could be applied pagination
func listProfilePeopleByKind(db *sqlx.DB, profileKind string, start int, count int) ([]ProfilePeople, error) {
	users := []ProfilePeople{}
	err := db.Select(&users, `SELECT * FROM ProfilePeople WHERE Profile_Kind=$1 LIMIT $2 OFFSET $3`, profileKind, count, start)
	if err != nil {
		return nil, err
	}
	return users, nil
}
