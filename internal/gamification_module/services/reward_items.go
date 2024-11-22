package services

// package services

// import (
// 	"english_app/internal/gamification_module/dto"
// 	reward_items_repo "english_app/internal/gamification_module/repository/reward_items"
// 	"english_app/pkg/errs"

// 	"github.com/google/uuid"
// )

// type RewardService interface {
// 	GetAllRewards() ([]*dto.RewardItemsResponse, errs.MessageErr)
// 	GetRewardDetail(id uuid.UUID) (*dto.RewardItemsResponse, errs.MessageErr)
// 	RedeemReward(id uuid.UUID) (*dto.ReedemItemResponse, errs.MessageErr)
// }

// type rewardService struct {
// 	repo        reward_items_repo.RewardRepository
// 	userService userRewardService
// }

// func NewRewardService(repo reward_items_repo.RewardRepository, userService userRewardService) RewardService {
// 	return &rewardService{repo, userService}
// }

// func (s *rewardService) GetAllRewards() ([]*dto.RewardItemsResponse, errs.MessageErr) {
// 	rewards, err := s.repo.FindAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Convert entity to DTO
// 	var rewardResponses []*dto.RewardItemsResponse
// 	for _, reward := range rewards {
// 		rewardResponses = append(rewardResponses, &dto.RewardItemsResponse{
// 			ID:          reward.ID,
// 			Name:        reward.Name,
// 			Points:      reward.Points,
// 			Description: reward.Description,
// 			Terms:       reward.Terms,
// 		})
// 	}

// 	return rewardResponses, nil
// }

// func (s *rewardService) GetRewardDetail(id uuid.UUID) (*dto.RewardItemsResponse, errs.MessageErr) {
// 	reward, err := s.repo.FindByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Convert entity to DTO
// 	rewardResponse := &dto.RewardItemsResponse{
// 		ID:          reward.ID,
// 		Name:        reward.Name,
// 		Points:      reward.Points,
// 		Description: reward.Description,
// 		Terms:       reward.Terms,
// 	}

// 	return rewardResponse, nil
// }

// // func (s *rewardService) RedeemReward(id uuid.UUID, userID uuid.UUID) (*dto.ReedemItemResponse, errs.MessageErr) {
// // 	reward, err := s.repo.FindByID(id)
// // 	if err != nil {
// // 		return nil,err
// // 	}

// // 	userReward,err := s.userService.GetByID(userID)

// // 	if err != nil {
// // 		return nil,err
// // 	}

// // 	if reward.Points <= 0 {
// // 		return "", errors.New("not enough points")
// // 	}

// // 	err = s.repo.RedeemReward(id)
// // 	if err != nil {
// // 		return "", err
// // 	}

// // 	return "Hadiah berhasil ditukar", nil
// // }
