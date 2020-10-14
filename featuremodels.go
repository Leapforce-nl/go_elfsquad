package elfsquad

import (
	"fmt"

	types "github.com/Leapforce-nl/go_types"
)

type FeatureModelsResponse struct {
	Context  string         `json:"@odata.context"`
	Value    []FeatureModel `json:"value"`
	NextLink string         `json:"@odata.nextLink"`
}

type FeatureModel struct {
	RootFeatureID    types.GUID `json:"rootFeatureId"`
	Order            int32      `json:"order"`
	DisplayPrices    bool       `json:"displayPrices"`
	HideInShowroom   bool       `json:"hideInShowroom"`
	HideInOrderEntry bool       `json:"hideInOrderEntry"`
	AutodeskUrn      string     `json:"autodeskUrn"`
	ID               types.GUID `json:"id"`
	CreatedDate      string     `json:"createdDate"`
	UpdatedDate      string     `json:"updatedDate"`
	OrganizationID   types.GUID `json:"organizationId"`
	CreatorID        types.GUID `json:"creatorId"`
}

func (es *Elfsquad) GetFeatureModels() (*[]FeatureModel, error) {
	top := 100
	skip := 0

	featureModels := []FeatureModel{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		url := fmt.Sprintf("%s/featuremodels?$top=%v&$skip=%v", apiURLData, top, skip)

		featureModelsReponse := FeatureModelsResponse{}

		_, err := es.oAuth2.Get(url, &featureModelsReponse)
		if err != nil {
			return nil, err
		}

		rowCount = len(featureModelsReponse.Value)

		if rowCount > 0 {
			featureModels = append(featureModels, featureModelsReponse.Value...)
		}

		skip += top
	}

	return &featureModels, nil
}
