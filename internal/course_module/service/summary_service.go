package service

import (
	"english_app/internal/course_module/dto"
	"english_app/internal/course_module/entity"
	summaryrepository "english_app/internal/course_module/repository/summary_repository"
	"english_app/pkg/errs"
	"english_app/pkg/gcloud"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type SummaryPartService interface {
	Create(request dto.SummaryPartRequest) (*dto.SummaryPartResponse, errs.MessageErr)
	FindByID(id uuid.UUID) (*dto.SummaryPartResponse, errs.MessageErr)
	Update(id uuid.UUID, request dto.SummaryPartRequest) (*dto.SummaryPartResponse, errs.MessageErr)
	Delete(id uuid.UUID) errs.MessageErr
}

type summaryPartService struct {
	repo        summaryrepository.SummaryPartRepository
	gcsUploader *gcloud.GCSUploader
	// storage *storage.Client
	// bucket  string
}

// Delete implements SummaryPartService.
func (s *summaryPartService) Delete(id uuid.UUID) errs.MessageErr {
	summary, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	objectName := summary.URL[strings.LastIndex(summary.URL, "/")+1:]
	err = s.gcsUploader.DeleteFile(objectName)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// FindByID implements SummaryPartService.
func (s *summaryPartService) FindByID(id uuid.UUID) (*dto.SummaryPartResponse, errs.MessageErr) {
	videoPart, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.SummaryPartResponse{
		ID:          videoPart.ID,
		Description: videoPart.Description,
		URL:         videoPart.URL,
		CreatedAt:   videoPart.CreatedAt,
		UpdatedAt:   videoPart.UpdatedAt,
	}, nil
}

// Update implements SummaryPartService.
func (s *summaryPartService) Update(id uuid.UUID, request dto.SummaryPartRequest) (*dto.SummaryPartResponse, errs.MessageErr) {
	oldSummary, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	newSummary := entity.SummaryPart{
		Description: request.Description,
	}

	// videoPart.Title = request.Title
	// videoPart.Description = request.Description
	// videoPart.URL = request.URL
	// videoPart.VideoExp = request.VideoExp
	// videoPart.VideoPoin = request.VideoPoin
	// videoPart.VideoPoin = request.VideoPoin

	if request.ContentType != "" {
		objectName := oldSummary.URL[strings.LastIndex(oldSummary.URL, "/")+1:]
		err = s.gcsUploader.DeleteFile(objectName)
		if err != nil {
			return nil, err
		}

		fmt.Print(objectName)

		fileURL, err := s.gcsUploader.UploadFile(request.FileSumary, request.ContentType)
		if err != nil {
			return nil, err
		}

		newSummary.URL = fileURL

	}
	updatedVideo, err := s.repo.Update(oldSummary, &newSummary)

	if err != nil {
		return nil, err
	}
	return &dto.SummaryPartResponse{
		ID:          updatedVideo.ID,
		Description: updatedVideo.Description,
		URL:         updatedVideo.URL,
		CreatedAt:   updatedVideo.CreatedAt,
		UpdatedAt:   updatedVideo.UpdatedAt,
	}, nil
}

func NewSummaryPartService(repo summaryrepository.SummaryPartRepository, gcsUploader *gcloud.GCSUploader) SummaryPartService {
	return &summaryPartService{repo: repo, gcsUploader: gcsUploader}
}

func (s *summaryPartService) Create(request dto.SummaryPartRequest) (*dto.SummaryPartResponse, errs.MessageErr) {

	fileURL, err := s.gcsUploader.UploadFile(request.FileSumary, request.ContentType)
	if err != nil {
		return nil, err
	}

	summaryPart := entity.SummaryPart{
		ID:          uuid.New(),
		Description: request.Description,
		URL:         fileURL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(&summaryPart); err != nil {
		return nil, err
	}

	return &dto.SummaryPartResponse{
		ID:          summaryPart.ID,
		Description: summaryPart.Description,
		URL:         summaryPart.URL,
		CreatedAt:   summaryPart.CreatedAt,
		UpdatedAt:   summaryPart.UpdatedAt,
	}, nil
}
