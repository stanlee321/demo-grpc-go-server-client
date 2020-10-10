package main

import (
	profiles_pb "github.com/stanlee321/demo-grpc-go-server-client/proto"
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

func (handler *userDataHandler) CreateProfilePoliticalParty(ctx context.Context, req *profiles_pb.CreateProfilePoliticalPartyRequest) (*profiles_pb.CreateProfilePoliticalPartyResponse, error) {
	fmt.Println("Create User request")

	user := req.GetProfile()

	data := ProfilePoliticalParty{
		ProfilePoliticalName:        user.GetProfilePoliticalName(),
		ProfilePoliticalSigla:       user.GetProfilePoliticalSigla(),
		ProfilePoliticalType:        user.GetProfilePoliticalType(),
		ProfilePoliticalTwitterID:   user.GetProfilePoliticalTwitterId(),
		ProfilePoliticalFacebookID:  user.GetProfilePoliticalFacebookID(),
		ProfilePoliticalInstagramID: user.GetProfilePoliticalFacebookID(),
		ProfilePoliticalYoutubeID:   user.GetProfilePoliticalYoutubeID(),
		ProfilePoliticalTikTokID:    user.GetProfilePoliticalTikTokID(),
	}

	// Enqueue the request for create user

	err := data.createPoliticalParty(handler.app.DB)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err))
	}

	return &profiles_pb.CreateProfilePoliticalPartyResponse{
		Profile: &profiles_pb.ProfilePoliticalParty{
			Id:                          data.ID,
			ProfilePoliticalName:        user.GetProfilePoliticalName(),
			ProfilePoliticalSigla:       user.GetProfilePoliticalSigla(),
			ProfilePoliticalType:        user.GetProfilePoliticalType(),
			ProfilePoliticalTwitterId:   user.GetProfilePoliticalTwitterId(),
			ProfilePoliticalFacebookID:  user.GetProfilePoliticalFacebookID(),
			ProfilePoliticalInstagramID: user.GetProfilePoliticalFacebookID(),
			ProfilePoliticalYoutubeID:   user.GetProfilePoliticalYoutubeID(),
			ProfilePoliticalTikTokID:    user.GetProfilePoliticalTikTokID(),
		},
	}, nil

}

// Read User Domain
func (handler *userDataHandler) ReadProfilePoliticalPartyByID(ctx context.Context, req *profiles_pb.ReadProfilePoliticalPartyByIdRequest) (*profiles_pb.ReadProfilePoliticalPartyByIdResponse, error) {
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

	var userdata ProfilePoliticalParty
	var err error

	userID := req.GetId()

	fmt.Println("MODE: PASSING USER ID....", userID)



	// If value does not exist in Cache
	// Read data from DB
	if userdata, err = handler.app.getProfilePoliticalPartyFromDB(userID); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("INTERAL ERROR trying to get value from DB... %v", err),
		)
	}


	// Return Query in ReadUser response format
	return &profiles_pb.ReadProfilePoliticalPartyByIdResponse{
		Profile: dataToprofilespb(userdata),
	}, nil
}

// START Update User Domain
func (handler *userDataHandler) UpdateProfilePoliticalParty(ctx context.Context, req *profiles_pb.UpdateProfilePoliticalPartyRequest) (*profiles_pb.UpdateProfilePoliticalPartyResponse, error) {
	var newuserdata ProfilePoliticalParty

	userdata := req.GetProfile()

	user := ProfilePoliticalParty{
		ID:                          userdata.GetId(),
		ProfilePoliticalName:        userdata.GetProfilePoliticalName(),
		ProfilePoliticalSigla:       userdata.GetProfilePoliticalSigla(),
		ProfilePoliticalType:        userdata.GetProfilePoliticalType(),
		ProfilePoliticalTwitterID:   userdata.GetProfilePoliticalTwitterId(),
		ProfilePoliticalFacebookID:  userdata.GetProfilePoliticalFacebookID(),
		ProfilePoliticalInstagramID: userdata.GetProfilePoliticalFacebookID(),
		ProfilePoliticalYoutubeID:   userdata.GetProfilePoliticalYoutubeID(),
		ProfilePoliticalTikTokID:    userdata.GetProfilePoliticalTikTokID(),
	}

	err := user.updatePolitcalParty(handler.app.DB)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	if newuserdata, err = handler.app.getProfilePoliticalPartyFromDB(user.ID); err == nil {
		return &profiles_pb.UpdateProfilePoliticalPartyResponse{
			Profile: dataToprofilespb(newuserdata),
		}, nil
	}


	return nil, status.Errorf(
		codes.Internal,
		fmt.Sprintf("Internal error: %v", err),
	)
}

// END Update User Domain

// Start Delete User Domain
func (handler *userDataHandler) DeleteProfilePoliticalParty(ctx context.Context, req *profiles_pb.DeleteProfilePoliticalPartyRequest) (*profiles_pb.DeleteProfilePoliticalPartyResponse, error) {
	id := req.GetId()

	user := ProfilePoliticalParty{ID: id}

	if err := user.deletePoliticalParty(handler.app.DB); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &profiles_pb.DeleteProfilePoliticalPartyResponse{
		Id: id,
	}, nil
}

// END Delete User Domain

// Start List Users Domain
func (handler *userDataHandler) ListProfilePoliticalParty(req *profiles_pb.ListProfilePoliticalPartyRequest, stream profiles_pb.ProfileService_ListProfilePoliticalPartyServer) error {
	fmt.Println("ListUser request")

	count := req.GetCount()
	start := req.GetStart()

	if count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	users, err := listPoliticalParty(handler.app.DB, int(start), int(count))

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	for _, quser := range users {
		stream.Send(&profiles_pb.ListProfilePoliticalPartyResponse{
			Profile: dataToprofilespb(quser),
		})
	}

	return nil

}

// END List Users Domain

// Start List Users Domain
func (handler *userDataHandler) ListProfilePoliticalPartyByType(req *profiles_pb.ListProfilePoliticalPartyByTypeRequest, stream profiles_pb.ProfileService_ListProfilePoliticalPartyByTypeServer) error {
	fmt.Println("ListUser request")

	userType := req.GetProfileType()
	count := req.GetCount()
	start := req.GetStart()

	if count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	users, err := listtPoliticalPartyByType(handler.app.DB, userType, int(start), int(count))

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	for _, quser := range users {
		stream.Send(&profiles_pb.ListProfilePoliticalPartyByTypeResponse{
			Profile: dataToprofilespb(quser),
		})
	}

	return nil

}

// END List Users Domain

func (a *App) getProfilePoliticalPartyFromDB(id int32) (ProfilePoliticalParty, error) {
	user := ProfilePoliticalParty{
		ID: id,
	}

	if err := user.getPoliticalPartyByID(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			return user, err
		default:
			return user, err
		}
	}
	return user, nil
}

func dataToprofilespb(userdata ProfilePoliticalParty) *profiles_pb.ProfilePoliticalParty {
	return &profiles_pb.ProfilePoliticalParty{
		Id:                          userdata.ID,
		ProfilePoliticalName:        userdata.ProfilePoliticalName,
		ProfilePoliticalSigla:       userdata.ProfilePoliticalSigla,
		ProfilePoliticalType:        userdata.ProfilePoliticalType,
		ProfilePoliticalTwitterId:   userdata.ProfilePoliticalTwitterID,
		ProfilePoliticalFacebookID:  userdata.ProfilePoliticalFacebookID,
		ProfilePoliticalInstagramID: userdata.ProfilePoliticalFacebookID,
		ProfilePoliticalYoutubeID:   userdata.ProfilePoliticalYoutubeID,
		ProfilePoliticalTikTokID:    userdata.ProfilePoliticalTikTokID,
		CreatedAt:                   userdata.CreatedAt,
		UpdatedAt:                   userdata.UpdatedAt,
	}
}
