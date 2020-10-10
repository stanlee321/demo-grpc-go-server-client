package main

import (
	profiles_pb "../../protofiles"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/getsentry/sentry-go"
)

func (handler *userDataHandler) CreateProfileStateOrganization(ctx context.Context, req *profiles_pb.CreateProfileStateOrganizationRequest) (*profiles_pb.CreateProfileStateOrganizationResponse, error) {
	fmt.Println("Create User request")

	user := req.GetProfile()

	data := ProfileStateOrganization{
		ID:                   user.GetId(),
		ProfileName:          user.GetProfileName(),
		ProfileSigla:         user.GetProfileSigla(),
		ProfilePoliticalType: user.GetProfilePoliticalType(),
		ProfileFase:          user.GetProfileFase(),
		ProfileTwitterID:     user.GetProfileTwitterId(),
		ProfileFacebookID:    user.GetProfileFacebookId(),
		ProfileInstagramID:   user.GetProfileFacebookId(),
		ProfileYoutubeID:     user.GetProfileYoutubeId(),
		ProfileTikTokID:      user.GetProfileTikTokId(),
	}

	handler.app.DB.Get(&data.ID, "SELECT nextval('ProfileStateOrganization_Profile_State_Organization_Id_seq')")

	// Enqueue the request for create user

	err := data.createProfileStateOrganization(handler.app.DB)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err))
	}

	return &profiles_pb.CreateProfileStateOrganizationResponse{
		Profile: dataToProfileStateOrganizationpb(data),
	}, nil

}

// Read User Domain
func (handler *userDataHandler) ReadProfileStateOrganizationByID(ctx context.Context, req *profiles_pb.ReadProfileStateOrganizationByIdRequest) (*profiles_pb.ReadProfileStateOrganizationByIdResponse, error) {
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

	var userdata ProfileStateOrganization
	var err error

	userID := req.GetId()

	fmt.Println("MODE: PASSING USER ID....", userID)

	// If value does not exist in Cache
	// Read data from DB
	if userdata, err = handler.app.getProfileStateOrganizationFromDB(userID); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("INTERAL ERROR trying to get value from DB... %v", err),
		)
	}

	// Set cache with Key/Values
	response, _ := json.Marshal(userdata)

	// Set JSON values to Cache
	if err := handler.app.Cache.setValue(userID, response); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("INTERAL ERROR trying to set Value in Cache:... %v", err),
		)

	}

	// Return Query in ReadUser response format
	return &profiles_pb.ReadProfileStateOrganizationByIdResponse{
		Profile: dataToProfileStateOrganizationpb(userdata),
	}, nil
}

// START Update User Domain
func (handler *userDataHandler) UpdateProfileStateOrganization(ctx context.Context, req *profiles_pb.UpdateProfileStateOrganizationRequest) (*profiles_pb.UpdateProfileStateOrganizationResponse, error) {
	var newuserdata ProfileStateOrganization

	userdata := req.GetProfile()

	user := ProfileStateOrganization{
		ID:                   userdata.GetId(),
		ProfileName:          userdata.GetProfileName(),
		ProfileSigla:         userdata.GetProfileSigla(),
		ProfilePoliticalType: userdata.GetProfilePoliticalType(),
		ProfileFase:          userdata.GetProfileFase(),
		ProfileTwitterID:     userdata.GetProfileTwitterId(),
		ProfileFacebookID:    userdata.GetProfileFacebookId(),
		ProfileInstagramID:   userdata.GetProfileFacebookId(),
		ProfileYoutubeID:     userdata.GetProfileYoutubeId(),
		ProfileTikTokID:      userdata.GetProfileTikTokId(),
	}

	err := user.updateProfileStateOrganization(handler.app.DB)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	if newuserdata, err = handler.app.getProfileStateOrganizationFromDB(user.ID); err == nil {
		return &profiles_pb.UpdateProfileStateOrganizationResponse{
			Profile: dataToProfileStateOrganizationpb(newuserdata),
		}, nil
	}

	// Set cache with Key/Values
	response, _ := json.Marshal(newuserdata)

	// Set JSON values to Cache
	if err := handler.app.Cache.setValue(user.ID, response); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("INTERAL ERROR trying to set Value in Cache:... %v", err),
		)

	}
	return nil, status.Errorf(
		codes.Internal,
		fmt.Sprintf("Internal error: %v", err),
	)
}

// END Update User Domain

// Start Delete User Domain
func (handler *userDataHandler) DeleteProfileStateOrganization(ctx context.Context, req *profiles_pb.DeleteProfileStateOrganizationRequest) (*profiles_pb.DeleteProfileStateOrganizationResponse, error) {
	id := req.GetId()

	user := ProfileStateOrganization{ID: id}

	if err := user.deleteProfileStateOrganization(handler.app.DB); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	// Delete  JSON values from Cache
	if err := handler.app.Cache.deleteKey(user.ID); err != nil {
		sentry.CaptureException(err)

		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("INTERAL ERROR trying to DEKETE Value in Cache:... %v", err),
		)

	}

	return &profiles_pb.DeleteProfileStateOrganizationResponse{
		Id: id,
	}, nil
}

// END Delete User Domain

// Start List Users Domain
func (handler *userDataHandler) ListProfileStateOrganization(req *profiles_pb.ListProfileStateOrganizationRequest, stream profiles_pb.ProfileService_ListProfileStateOrganizationServer) error {
	fmt.Println("ListUser request")

	count := req.GetCount()
	start := req.GetStart()

	if count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	users, err := listProfileStateOrganization(handler.app.DB, int(start), int(count))

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	for _, quser := range users {
		stream.Send(&profiles_pb.ListProfileStateOrganizationResponse{
			Profile: dataToProfileStateOrganizationpb(quser),
		})
	}

	return nil

}

// END List Users Domain

// Start List Users Domain
func (handler *userDataHandler) ListProfileStateOrganizationByType(req *profiles_pb.ListProfileStateOrganizationByTypeRequest, stream profiles_pb.ProfileService_ListProfileStateOrganizationByTypeServer) error {
	fmt.Println("ListUser request")

	usertype := req.GetProfileType()
	count := req.GetCount()
	start := req.GetStart()

	if count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	users, err := listProfileStateOrganizationByType(handler.app.DB, usertype, int(start), int(count))

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	for _, quser := range users {
		stream.Send(&profiles_pb.ListProfileStateOrganizationByTypeResponse{
			Profile: dataToProfileStateOrganizationpb(quser),
		})
	}

	return nil

}

// END List Users Domain

func (a *App) getProfileStateOrganizationFromDB(id int32) (ProfileStateOrganization, error) {
	user := ProfileStateOrganization{ID: id}

	if err := user.getProfileStateOrganizationByID(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			return user, err
		default:
			return user, err
		}
	}
	return user, nil
}

func dataToProfileStateOrganizationpb(userdata ProfileStateOrganization) *profiles_pb.ProfileStateOrganization {
	return &profiles_pb.ProfileStateOrganization{
		Id:                   userdata.ID,
		ProfileName:          userdata.ProfileName,
		ProfileSigla:         userdata.ProfileSigla,
		ProfilePoliticalType: userdata.ProfilePoliticalType,
		ProfileFase:          userdata.ProfileFase,
		ProfileTwitterId:     userdata.ProfileTwitterID,
		ProfileFacebookId:    userdata.ProfileFacebookID,
		ProfileInstagramId:   userdata.ProfileFacebookID,
		ProfileYoutubeId:     userdata.ProfileYoutubeID,
		ProfileTikTokId:      userdata.ProfileTikTokID,
		CreatedAt:            userdata.CreatedAt,
		UpdatedAt:            userdata.UpdatedAt,
	}
}
