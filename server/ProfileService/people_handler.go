package main

import (
	profiles_pb "github.com/stanlee321/demo-grpc-go-server-client/proto"
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/getsentry/sentry-go"
)

func (handler *userDataHandler) CreateProfilePeople(ctx context.Context, req *profiles_pb.CreateProfilePeopleRequest) (*profiles_pb.CreateProfilePeopleResponse, error) {
	fmt.Println("Create User request")

	user := req.GetProfile()

	data := ProfilePeople{
		ID:                 user.GetId(),
		ProfileFirstName:   user.GetProfileFirstName(),
		ProfileLastName:    user.GetProfileLastName(),
		ProfileOrientation: user.GetProfileOrientation(),
		ProfileKind:        user.GetProfileKind(),
		ProfileTwitterID:   user.GetProfileTwitterId(),
		ProfileFacebookID:  user.GetProfileFacebookId(),
		ProfileInstagramID: user.GetProfileFacebookId(),
		ProfileYoutubeID:   user.GetProfileYoutubeId(),
		ProfileTikTokID:    user.GetProfileTikTokId(),
	}

	handler.app.DB.Get(&data.ID, "SELECT nextval('ProfilePeople_Profile_People_Id_seq')")

	err := data.createProfilePeople(handler.app.DB)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err))
	}

	return &profiles_pb.CreateProfilePeopleResponse{
		Profile: dataToProfilePeoplepb(data),
	}, nil

}

// Read User Domain
func (handler *userDataHandler) ReadProfilePeopleByID(ctx context.Context, req *profiles_pb.ReadProfilePeopleByIdRequest) (*profiles_pb.ReadProfilePeopleByIdResponse, error) {
	/*
		This method holds the Read User logic.

		PARAMS:
		------

		req : {
			.GetEmail(): string with the email value ,
			.GetUserId(): string with the userId value
		}
	*/
	fmt.Println("Reading ReadUser request")

	var userdata ProfilePeople
	var err error

	userID := req.GetId()

	fmt.Println("MODE: PASSING USER ID....", userID)

	// If value does not exist in Cache
	// Read data from DB
	if userdata, err = handler.app.getProfilePeopleFromDB(userID); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("INTERAL ERROR trying to get value from DB... %v", err),
		)
	}

	// Return Query in ReadUser response format
	return &profiles_pb.ReadProfilePeopleByIdResponse{
		Profile: dataToProfilePeoplepb(userdata),
	}, nil
}

// START Update User Domain
func (handler *userDataHandler) UpdateProfilePeople(ctx context.Context, req *profiles_pb.UpdateProfilePeopleRequest) (*profiles_pb.UpdateProfilePeopleResponse, error) {
	var newuserdata ProfilePeople

	userdata := req.GetProfile()

	user := ProfilePeople{
		ID:                 userdata.GetId(),
		ProfileFirstName:   userdata.GetProfileFirstName(),
		ProfileLastName:    userdata.GetProfileLastName(),
		ProfileOrientation: userdata.GetProfileOrientation(),
		ProfileKind:        userdata.GetProfileKind(),
		ProfileTwitterID:   userdata.GetProfileTwitterId(),
		ProfileFacebookID:  userdata.GetProfileFacebookId(),
		ProfileInstagramID: userdata.GetProfileFacebookId(),
		ProfileYoutubeID:   userdata.GetProfileYoutubeId(),
		ProfileTikTokID:    userdata.GetProfileTikTokId(),
	}

	err := user.updateProfilePeople(handler.app.DB)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	if newuserdata, err = handler.app.getProfilePeopleFromDB(user.ID); err == nil {
		return &profiles_pb.UpdateProfilePeopleResponse{
			Profile: dataToProfilePeoplepb(newuserdata),
		}, nil
	}

	return nil, status.Errorf(
		codes.Internal,
		fmt.Sprintf("Internal error: %v", err),
	)
}

// END Update User Domain

// Start Delete User Domain
func (handler *userDataHandler) DeleteProfilePeople(ctx context.Context, req *profiles_pb.DeleteProfilePeopleRequest) (*profiles_pb.DeleteProfilePeopleResponse, error) {
	id := req.GetId()

	user := ProfilePeople{ID: id}

	if err := user.deleteProfilePeople(handler.app.DB); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &profiles_pb.DeleteProfilePeopleResponse{
		Id: id,
	}, nil
}

// END Delete User Domain

// Start List Users Domain
func (handler *userDataHandler) ListProfilePeople(req *profiles_pb.ListProfilePeopleRequest, stream profiles_pb.ProfileService_ListProfilePeopleServer) error {
	fmt.Println("ListUser request")

	count := req.GetCount()
	start := req.GetStart()

	if count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	users, err := listProfilePeople(handler.app.DB, int(start), int(count))

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	for _, quser := range users {
		stream.Send(&profiles_pb.ListProfilePeopleResponse{
			Profile: dataToProfilePeoplepb(quser),
		})
	}

	return nil

}

// END List Users Domain

// Start List Users Domain
func (handler *userDataHandler) ListProfilePeopleByKind(req *profiles_pb.ListProfilePeopleByKindRequest, stream profiles_pb.ProfileService_ListProfilePeopleByKindServer) error {
	fmt.Println("ListUser request")

	userkind := req.GetProfileKind()
	count := req.GetCount()
	start := req.GetStart()

	if count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	users, err := listProfilePeopleByKind(handler.app.DB, userkind, int(start), int(count))

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	for _, quser := range users {
		stream.Send(&profiles_pb.ListProfilePeopleByKindResponse{
			Profile: dataToProfilePeoplepb(quser),
		})
	}

	return nil

}

// END List Users Domain

func (a *App) getProfilePeopleFromDB(id int32) (ProfilePeople, error) {
	user := ProfilePeople{ID: id}

	if err := user.getProfilePeopleByUserID(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			return user, err
		default:
			return user, err
		}
	}
	return user, nil
}

func dataToProfilePeoplepb(userdata ProfilePeople) *profiles_pb.ProfilePeople {
	return &profiles_pb.ProfilePeople{
		Id:                 userdata.ID,
		ProfileFirstName:   userdata.ProfileFirstName,
		ProfileLastName:    userdata.ProfileLastName,
		ProfileOrientation: userdata.ProfileOrientation,
		ProfileKind:        userdata.ProfileKind,
		ProfileTwitterId:   userdata.ProfileTwitterID,
		ProfileFacebookId:  userdata.ProfileFacebookID,
		ProfileInstagramId: userdata.ProfileFacebookID,
		ProfileYoutubeId:   userdata.ProfileYoutubeID,
		ProfileTikTokId:    userdata.ProfileTikTokID,
		CreatedAt:          userdata.CreatedAt,
		UpdatedAt:          userdata.UpdatedAt,
	}
}
