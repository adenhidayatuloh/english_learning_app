// package services
package services

// import (
// 	"english_app/internal/gamification_module/dto"
// 	"english_app/internal/gamification_module/entity"
// 	userreward "english_app/internal/gamification_module/repository/user_reward"
// 	"english_app/pkg/common"
// 	"english_app/pkg/errs"

// 	"github.com/google/uuid"
// )

// type UserRewardService interface {
// 	Create(input *dto.CreateUserRewardRequest) (*dto.UserRewardResponse, errs.MessageErr)
// 	GetByID(userID uuid.UUID) (*dto.UserRewardResponse, errs.MessageErr)
// 	GetAll() ([]dto.UserRewardResponse, errs.MessageErr)
// 	Update(input *dto.CreateUserRewardRequest) (*dto.UserRewardResponse, errs.MessageErr)
// 	Delete(id uint) errs.MessageErr
// 	GetUserLevel(userID uuid.UUID) (*dto.UserRewardLevelResponse, errs.MessageErr)
// }

// type userRewardService struct {
// 	repo userreward.UserRewardRepository
// }

// func NewUserRewardService(repo userreward.UserRewardRepository) UserRewardService {
// 	return &userRewardService{repo}
// }

// func (s *userRewardService) Create(input *dto.CreateUserRewardRequest) (*dto.UserRewardResponse, errs.MessageErr) {
// 	userReward := entity.UserReward{
// 		UserID:      input.UserID,
// 		TotalPoints: input.TotalPoints,
// 		TotalExp:    input.TotalExp,
// 		HelpCount:   input.HelpCount,
// 		HealthCount: input.HealthCount,
// 	}
// 	err := s.repo.Create(&userReward)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &dto.UserRewardResponse{
// 		ID:          userReward.ID,
// 		UserID:      userReward.UserID,
// 		TotalPoints: userReward.TotalPoints,
// 		TotalExp:    userReward.TotalExp,
// 		HelpCount:   userReward.HelpCount,
// 		HealthCount: userReward.HealthCount,
// 	}, nil
// }

// func (s *userRewardService) GetByID(userID uuid.UUID) (*dto.UserRewardResponse, errs.MessageErr) {
// 	userReward, err := s.repo.GetByUserID(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &dto.UserRewardResponse{
// 		ID:          userReward.ID,
// 		UserID:      userReward.UserID,
// 		TotalPoints: userReward.TotalPoints,
// 		TotalExp:    userReward.TotalExp,
// 		HelpCount:   userReward.HelpCount,
// 		HealthCount: userReward.HealthCount,
// 	}, nil
// }

// func (s *userRewardService) GetAll() ([]dto.UserRewardResponse, errs.MessageErr) {
// 	userRewards, err := s.repo.GetAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var responses []dto.UserRewardResponse
// 	for _, userReward := range userRewards {
// 		responses = append(responses, dto.UserRewardResponse{
// 			ID:          userReward.ID,
// 			UserID:      userReward.UserID,
// 			TotalPoints: userReward.TotalPoints,
// 			TotalExp:    userReward.TotalExp,
// 			HelpCount:   userReward.HelpCount,
// 			HealthCount: userReward.HealthCount,
// 		})
// 	}

// 	return responses, nil
// }

// func (s *userRewardService) Update(input *dto.CreateUserRewardRequest) (*dto.UserRewardResponse, errs.MessageErr) {
// 	userReward, err := s.repo.GetByUserID(input.UserID)
// 	if err != nil {
// 		createReward, er := s.Create(input)

// 		if er != nil {
// 			return nil, er
// 		}

// 		return &dto.UserRewardResponse{
// 			ID:          createReward.ID,
// 			UserID:      createReward.UserID,
// 			TotalPoints: createReward.TotalPoints,
// 			TotalExp:    createReward.TotalExp,
// 			HelpCount:   createReward.HelpCount,
// 			HealthCount: createReward.HealthCount,
// 		}, nil
// 	}

// 	userReward.TotalPoints = userReward.TotalPoints + input.TotalPoints
// 	userReward.TotalExp = userReward.TotalExp + input.TotalExp
// 	// userReward.HelpCount = input.HelpCount
// 	// userReward.HealthCount = input.HealthCount

// 	err = s.repo.Update(userReward)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &dto.UserRewardResponse{
// 		ID:          userReward.ID,
// 		UserID:      userReward.UserID,
// 		TotalPoints: userReward.TotalPoints,
// 		TotalExp:    userReward.TotalExp,
// 		HelpCount:   userReward.HelpCount,
// 		HealthCount: userReward.HealthCount,
// 	}, nil
// }

// func (s *userRewardService) Delete(id uint) errs.MessageErr {
// 	return s.repo.Delete(id)
// }

// func (s *userRewardService) GetUserLevel(userID uuid.UUID) (*dto.UserRewardLevelResponse, errs.MessageErr) {
// 	// Ambil data user reward berdasarkan user ID
// 	userReward, err := s.repo.GetByUserID(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Hitung level
// 	level, nextLevelExp := common.CalculateLevel(userReward.TotalExp)

// 	// Buat response
// 	response := &dto.UserRewardLevelResponse{
// 		Level:        level,
// 		CurrentExp:   userReward.TotalExp,
// 		NextLevelExp: nextLevelExp,
// 		TotalPoints:  userReward.TotalPoints,
// 	}
// 	return response, nil
// }
