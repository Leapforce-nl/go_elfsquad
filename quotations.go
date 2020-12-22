package elfsquad

import (
	"fmt"
	url "net/url"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	types "github.com/leapforce-libraries/go_types"
)

type QuotationsResponse struct {
	Context  string      `json:"@odata.context"`
	Value    []Quotation `json:"value"`
	NextLink string      `json:"@odata.nextLink"`
}

type Quotation struct {
	SellerID            *types.GUID `json:"sellerId,omitempty"`
	SellerContactID     *types.GUID `json:"sellerContactId,omitempty"`
	DebtorID            *types.GUID `json:"debtorId,omitempty"`
	DebtorContactID     *types.GUID `json:"debtorContactId,omitempty"`
	ShipToID            *types.GUID `json:"shipToId,omitempty"`
	ShipToContactID     *types.GUID `json:"shipToContactId,omitempty"`
	Synced              bool        `json:"synced,omitempty"`
	QuotationNumber     int64       `json:"quotationNumber,omitempty"`
	VersionNumber       int32       `json:"versionNumber,omitempty"`
	Status              string      `json:"status,omitempty"`
	Subject             string      `json:"subject,omitempty"`
	TotalPrice          float64     `json:"totalPrice,omitempty"`
	IsVerified          bool        `json:"isVerified,omitempty"`
	CustomerReference   string      `json:"customerReference,omitempty"`
	QuotationReference  string      `json:"quotationReference,omitempty"`
	Deliverydate        string      `json:"deliverydate,omitempty"`
	Remarks             string      `json:"remarks,omitempty"`
	ExpiresDate         string      `json:"expiresDate,omitempty"`
	QuotationTemplateID *types.GUID `json:"quotationTemplateId,omitempty"`
	ID                  *types.GUID `json:"id,omitempty"`
	CreatedDate         string      `json:"createdDate,omitempty"`
	UpdatedDate         string      `json:"updatedDate,omitempty"`
	OrganizationID      *types.GUID `json:"organizationId,omitempty"`
	CreatorID           *types.GUID `json:"creatorId,omitempty"`
}

type GetQuotationsParams struct {
	QuotationNumber *int64
	Select          *[]string
}

func (es *Elfsquad) GetQuotations(params *GetQuotationsParams) (*[]Quotation, *errortools.Error) {
	top := 100
	skip := 0

	filter := []string{}

	if params != nil {
		if params.QuotationNumber != nil {
			filter = append(filter, fmt.Sprintf("QuotationNumber eq %v", *params.QuotationNumber))
		}
	}

	quotations := []Quotation{}

	rowCount := 0

	for skip == 0 || rowCount > 0 {
		urlPath := fmt.Sprintf("quotations?$top=%v&$skip=%v", top, skip)

		if len(filter) > 0 {
			urlPath = fmt.Sprintf("%s&$filter=%s", urlPath, url.QueryEscape(strings.Join(filter, " AND ")))
		}
		if params.Select != nil {
			urlPath = fmt.Sprintf("%s&$select=%s", urlPath, strings.Join(*params.Select, ","))
		}

		quotationsReponse := QuotationsResponse{}
		_, _, e := es.get(urlPath, &quotationsReponse)
		if e != nil {
			return nil, e
		}

		rowCount = len(quotationsReponse.Value)

		if rowCount > 0 {
			quotations = append(quotations, quotationsReponse.Value...)
		}

		skip += top
	}

	return &quotations, nil
}

func (es *Elfsquad) UpdateQuotation(quotationID types.GUID, quotationUpdate *Quotation) *errortools.Error {
	urlPath := fmt.Sprintf("quotations(%s)", quotationID.String())

	_, _, e := es.patch(urlPath, quotationUpdate, nil)

	return e
}
