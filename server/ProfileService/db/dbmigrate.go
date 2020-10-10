package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const tableCreationQuery = `
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TABLE IF NOT EXISTS ProfilePoliticalParty
(
	Id 		SERIAL PRIMARY KEY,
	Profile_Political_Name			TEXT NOT NULL,
	Profile_Political_Sigla 		TEXT,
	Profile_Political_Type			TEXT,
	Profile_Political_Twitter_Id	BIGINT,
	Profile_Political_Facebook_Id	BIGINT,
	Profile_Political_Instagram_Id	BIGINT,
	Profile_Political_Youtube_Id	BIGINT,
	Profile_Political_TikTok_Id		BIGINT,

	Created_at 				TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	Updated_at 				TIMESTAMPTZ NOT NULL DEFAULT NOW()

);

CREATE TRIGGER set_timestamp
	BEFORE UPDATE ON ProfilePoliticalParty
	FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


CREATE TABLE IF NOT EXISTS ProfileStateOrganization
(
	Id								SERIAL PRIMARY KEY,
	Profile_Name					TEXT NOT NULL,
	Profile_Sigla					TEXT,
	Profile_Fase					TEXT,
	Profile_Political_Type			TEXT,
	Profile_Twitter_Id				BIGINT,
	Profile_Facebook_Id				BIGINT,
	Profile_Instagram_Id			BIGINT,
	Profile_Youtube_Id				BIGINT,
	Profile_TikTok_Id				BIGINT,
	Created_at 				TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	Updated_at 				TIMESTAMPTZ NOT NULL DEFAULT NOW()

);
CREATE TRIGGER set_timestamp
	BEFORE UPDATE ON ProfileStateOrganization
	FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


CREATE TABLE IF NOT EXISTS ProfilePeople
(
	Id 								SERIAL PRIMARY KEY,
	Profile_FirstName				TEXT NOT NULL,
	Profile_LastName				TEXT,
	Profile_Orientation				TEXT,
	Profile_Twitter_Id				BIGINT,
	Profile_Facebook_Id				BIGINT,
	Profile_Instagram_Id			BIGINT,
	Profile_Youtube_Id				BIGINT,
	Profile_TikTok_Id				BIGINT,
	Profile_Kind 					TEXT NOT NULL,
	Created_at 				TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	Updated_at 				TIMESTAMPTZ NOT NULL DEFAULT NOW()

);
CREATE TRIGGER set_timestamp
	BEFORE UPDATE ON ProfilePeople
	FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


`

func main() {
	fmt.Println("Starting Twitter migration")
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_DEV_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished Twitter migration")
}
