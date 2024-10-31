package service

import (
	"english_app/internal/course_module/dto"
	"english_app/internal/course_module/entity"
	videorepository "english_app/internal/course_module/repository/video_repository"
	"english_app/pkg/errs"
	"english_app/pkg/gcloud"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type VideoPartService interface {
	Create(request dto.VideoPartRequest) (*dto.VideoPartResponse, errs.MessageErr)
	FindByID(id uuid.UUID) (*dto.VideoPartResponse, errs.MessageErr)
	Update(id uuid.UUID, request dto.VideoPartRequest) (*dto.VideoPartResponse, errs.MessageErr)
	Delete(id uuid.UUID) errs.MessageErr
}

type videoPartService struct {
	repo        videorepository.VideoPartRepository
	gcsUploader *gcloud.GCSUploader
}

func NewVideoPartService(repo videorepository.VideoPartRepository, gcsUploader *gcloud.GCSUploader) VideoPartService {
	return &videoPartService{repo: repo, gcsUploader: gcsUploader}
}

func (s *videoPartService) Create(request dto.VideoPartRequest) (*dto.VideoPartResponse, errs.MessageErr) {

	fileURL, err := s.gcsUploader.UploadFile(request.FileVideo, request.ContentType)
	if err != nil {
		return nil, err
	}

	videoPart := entity.VideoPart{
		ID:          uuid.New(),
		Title:       request.Title,
		Description: request.Description,
		URL:         fileURL,
		VideoExp:    request.VideoExp,
		VideoPoin:   request.VideoPoin,
	}
	if err := s.repo.CreateVideo(&videoPart); err != nil {
		return nil, err
	}
	return &dto.VideoPartResponse{
		ID:          videoPart.ID,
		Title:       videoPart.Title,
		Description: videoPart.Description,
		URL:         videoPart.URL,
		VideoExp:    videoPart.VideoExp,
		VideoPoin:   videoPart.VideoPoin,
		CreatedAt:   videoPart.CreatedAt,
		UpdatedAt:   videoPart.UpdatedAt,
	}, nil
}

func (s *videoPartService) FindByID(id uuid.UUID) (*dto.VideoPartResponse, errs.MessageErr) {
	videoPart, err := s.repo.FindVideoByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.VideoPartResponse{
		ID:          videoPart.ID,
		Title:       videoPart.Title,
		Description: videoPart.Description,
		URL:         videoPart.URL,
		VideoExp:    videoPart.VideoExp,
		VideoPoin:   videoPart.VideoPoin,
		CreatedAt:   videoPart.CreatedAt,
		UpdatedAt:   videoPart.UpdatedAt,
	}, nil
}

func (s *videoPartService) Update(id uuid.UUID, request dto.VideoPartRequest) (*dto.VideoPartResponse, errs.MessageErr) {
	oldvideoPart, err := s.repo.FindVideoByID(id)
	if err != nil {
		return nil, err
	}

	newVideoPart := entity.VideoPart{
		Title:       request.Title,
		Description: request.Description,
		VideoExp:    request.VideoExp,
		VideoPoin:   request.VideoPoin,
	}

	// videoPart.Title = request.Title
	// videoPart.Description = request.Description
	// videoPart.URL = request.URL
	// videoPart.VideoExp = request.VideoExp
	// videoPart.VideoPoin = request.VideoPoin
	// videoPart.VideoPoin = request.VideoPoin

	if request.ContentType != "" {
		objectName := oldvideoPart.URL[strings.LastIndex(oldvideoPart.URL, "/")+1:]
		err = s.gcsUploader.DeleteFile(objectName)
		if err != nil {
			return nil, err
		}

		fmt.Print(objectName)

		fileURL, err := s.gcsUploader.UploadFile(request.FileVideo, request.ContentType)
		if err != nil {
			return nil, err
		}

		newVideoPart.URL = fileURL

	}
	updatedVideo, err := s.repo.UpdateVideoBatch(oldvideoPart, &newVideoPart)

	if err != nil {
		return nil, err
	}
	return &dto.VideoPartResponse{
		ID:          updatedVideo.ID,
		Title:       updatedVideo.Title,
		Description: updatedVideo.Description,
		URL:         updatedVideo.URL,
		VideoExp:    updatedVideo.VideoExp,
		VideoPoin:   updatedVideo.VideoPoin,
		CreatedAt:   updatedVideo.CreatedAt,
		UpdatedAt:   updatedVideo.UpdatedAt,
	}, nil
}

func (s *videoPartService) Delete(id uuid.UUID) errs.MessageErr {

	videoPart, err := s.repo.FindVideoByID(id)
	if err != nil {
		return err
	}
	objectName := videoPart.URL[strings.LastIndex(videoPart.URL, "/")+1:]
	err = s.gcsUploader.DeleteFile(objectName)
	if err != nil {
		return err
	}

	return s.repo.DeleteVideoByID(id)
}
