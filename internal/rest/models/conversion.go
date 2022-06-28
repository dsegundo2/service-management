package models

import (
	"time"

	"github.com/dsegundo2/service-management/internal/servicedb"
)

/*	File for converting database structs to api structs	*/

func ConvertListServicesResponse(s []*servicedb.Service, totalCount int) *ListServicesResponse {
	returnService := &ListServicesResponse{
		TotalCount: int32(totalCount),
		Services: func() []Service {
			sList := []Service{}
			for _, ser := range s {
				sList = append(sList, ConvertToApiService(ser))
			}
			return sList
		}(),
	}

	return returnService
}

func ConvertToApiService(s *servicedb.Service) Service {
	apiService := Service{
		Id:          s.ID,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
		Title:       s.Title,
		Description: s.Description,
		Versions: func() []ServiceVersion {
			versions := []ServiceVersion{}
			for _, sv := range s.Versions {
				versions = append(versions, ConvertToApiVersion(sv))
			}
			return versions
		}(),
		VersionCount: float64(s.VersionCount),
	}
	if s.DeletedAt != nil {
		apiService.DeletedAt = s.DeletedAt.Time
	}
	return apiService
}

func ConvertToApiVersion(v *servicedb.ServiceVersion) ServiceVersion {
	return ServiceVersion{
		Version:   v.Version,
		ServiceId: v.ServiceID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		DeletedAt: func() time.Time {
			if v.DeletedAt != nil {
				return v.DeletedAt.Time
			}
			return time.Time{}
		}(),
		ServiceExampleField: v.ServiceExampleField,
	}
}
