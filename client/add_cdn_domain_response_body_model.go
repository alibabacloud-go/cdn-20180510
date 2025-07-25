// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCdnDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddCdnDomainResponseBody
	GetRequestId() *string
}

type AddCdnDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCdnDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCdnDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddCdnDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCdnDomainResponseBody) SetRequestId(v string) *AddCdnDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCdnDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
