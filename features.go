package elfsquad

import (
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
	types "github.com/leapforce-libraries/go_types"
)

type FeaturesResponse struct {
	Context  string    `json:"@odata.context"`
	Value    []Feature `json:"value"`
	NextLink string    `json:"@odata.nextLink"`
}

type Feature struct {
	Name             string     `json:"name"`
	ArticleCode      string     `json:"articleCode"`
	Type             string     `json:"type"`
	SalesPrice       float64    `json:"salesPrice"`
	DisallowDiscount bool       `json:"disallowDiscount"`
	MinValue         float64    `json:"minValue"`
	MaxValue         float64    `json:"maxValue"`
	StepValue        float64    `json:"stepValue"`
	PackingUnit      float64    `json:"packingUnit"`
	CategoryID       types.GUID `json:"categoryId"`
	MarginPct        float64    `json:"marginPct"`
	CardImageURL     string     `json:"cardImageUrl"`
	ID               types.GUID `json:"id"`
	CreatorID        types.GUID `json:"creatorId"`
	Reference        types.GUID `json:"reference"`
	Synced           bool       `json:"synced"`
	Inactive         bool       `json:"inactive"`
	CreatedDate      string     `json:"createdDate"`
	UpdatedDate      string     `json:"updatedDate"`
}

func (service *Service) GetFeatures() (*[]Feature, *errortools.Error) {
	top := 100
	skip := 0

	features := []Feature{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("features?$top=%v&$skip=%v", top, skip)

		featuresResponse := FeaturesResponse{}
		requestConfig := oauth2.RequestConfig{
			URL:           service.url(urlPath),
			ResponseModel: &featuresResponse,
		}
		_, _, e := service.get(&requestConfig)
		if e != nil {
			return nil, e
		}

		rowCount = len(featuresResponse.Value)

		if rowCount > 0 {
			features = append(features, featuresResponse.Value...)
		}

		skip += top
	}

	return &features, nil
}
