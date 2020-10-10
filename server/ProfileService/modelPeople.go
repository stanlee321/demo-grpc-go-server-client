package main

import (
	"github.com/jmoiron/sqlx"
)

// ProfilePeople is the struct resposible to represent the database entity
type ProfilePeople struct {
	ID                 int32  `bson:"Profile_People_Id" db:"Profile_People_Id"`
	ProfileFirstName   string `bson:"Profile_FirstName" db:"Profile_FirstName"`
	ProfileLastName    string `bson:"Profile_LastName" db:"Profile_LastName"`
	ProfileOrientation string `bson:"Profile_Orientation" db:"Profile_Orientation"`
	ProfileTwitterID   int32  `bson:"Profile_Twitter_Id" db:"Profile_Twitter_Id"`
	ProfileFacebookID  int32  `bson:"Profile_Facebook_Id" db:"Profile_Facebook_Id"`
	ProfileInstagramID int32  `bson:"Profile_Instagram_Id" db:"Profile_Instagram_Id"`
	ProfileYoutubeID   int32  `bson:"Profile_Youtube_Id" db:"Profile_Youtube_Id"`
	ProfileTikTokID    int32  `bson:"Profile_TikTok_Id" db:"Profile_TikTok_Id"`
	ProfileKind        string `bson:"Profile_Kind" db:"Profile_Kind"`
	CreatedAt          string `bson:"created_at" db:"created_at"`
	UpdatedAt          string `bson:"updated_at" db:"updated_at"`
}

// getByUserName returns the quer SELECT by Email
func (u *ProfilePeople) getProfilePeopleByUserID(db *sqlx.DB) error {
	return db.Get(u, "SELECT * FROM ProfilePeople WHERE Profile_People_Id=$1", u.ID)
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
										Profile_Kind			=$9,
			WHERE Profile_People_Id=$10`,
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
	_, err := db.Exec(`DELETE FROM ProfilePeople WHERE Profile_People_Id=$1`, u.ID)

	return err
}

// Create a new user in the db using the instance values
func (u *ProfilePeople) createProfilePeople(db *sqlx.DB) error {

	return db.QueryRow(
		`INSERT INTO ProfilePeople(
			Profile_People_Id, 
			Profile_FirstName,
			Profile_LastName,
			Profile_Orientation	,
			Profile_Twitter_Id, 
			Profile_Facebook_Id	,
			Profile_Instagram_Id,
			Profile_Youtube_Id,
			Profile_TikTok_Id,
			Profile_Kind,
		) 
	
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING Profile_People_Id`,
		u.ID,
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
