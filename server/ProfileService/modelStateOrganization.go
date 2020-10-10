package main

import (
	"github.com/jmoiron/sqlx"
)

// ProfileStateOrganization is the struct resposible to represent the database entity
type ProfileStateOrganization struct {
	ID                   int32  `bson:"Profile_State_Organization_Id" db:"Profile_State_Organization_Id"`
	ProfileName          string `bson:"Profile_Name" db:"Profile_Name"`
	ProfileSigla         string `bson:"Profile_Sigla" db:"Profile_Sigla"`
	ProfilePoliticalType string `bson:"Profile_Political_Type" db:"Profile_Political_Type"`
	ProfileFase          string `bson:"Profile_Fase" db:"Profile_Fase"`
	ProfileTwitterID     int32  `bson:"Profile_Twitter_Id" db:"Profile_Twitter_Id"`
	ProfileFacebookID    int32  `bson:"Profile_Facebook_Id" db:"Profile_Facebook_Id"`
	ProfileInstagramID   int32  `bson:"Profile_Instagram_Id" db:"Profile_Instagram_Id"`
	ProfileYoutubeID     int32  `bson:"Profile_Youtube_Id" db:"Profile_Youtube_Id"`
	ProfileTikTokID      int32  `bson:"Profile_TikTok_Id" db:"Profile_TikTok_Id"`
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
										Profile_Twitter_Id		=$4, 
										Profile_Facebook_Id		=$5,
										Profile_Instagram_Id	=$6,
										Profile_Youtube_Id		=$7,
										Profile_TikTok_Id		=$8,
										password				=$9,
			WHERE Profile_State_Organization_Id=$10`,
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
			Profile_State_Organization_Id, 
			Profile_Name			,
			Profile_Sigla			,
			Profile_Political_Type	,
			Profile_Twitter_Id		, 
			Profile_Facebook_Id		,
			Profile_Instagram_Id	,
			Profile_Youtube_Id		,
			Profile_TikTok_Id		,
			password				,
		) 
	
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING Profile_State_Organization_Id`,
		u.ID,
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
