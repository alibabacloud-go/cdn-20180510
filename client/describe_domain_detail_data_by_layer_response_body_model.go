// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDetailDataByLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDomainDetailDataByLayerResponseBodyData) *DescribeDomainDetailDataByLayerResponseBody
	GetData() *DescribeDomainDetailDataByLayerResponseBodyData
	SetRequestId(v string) *DescribeDomainDetailDataByLayerResponseBody
	GetRequestId() *string
}

type DescribeDomainDetailDataByLayerResponseBody struct {
	// The number of IPv6 requests per second.
	Data *DescribeDomainDetailDataByLayerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The number of queries per second.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainDetailDataByLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailDataByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailDataByLayerResponseBody) GetData() *DescribeDomainDetailDataByLayerResponseBodyData {
	return s.Data
}

func (s *DescribeDomainDetailDataByLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainDetailDataByLayerResponseBody) SetData(v *DescribeDomainDetailDataByLayerResponseBodyData) *DescribeDomainDetailDataByLayerResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBody) SetRequestId(v string) *DescribeDomainDetailDataByLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainDetailDataByLayerResponseBodyData struct {
	DataModule []*DescribeDomainDetailDataByLayerResponseBodyDataDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainDetailDataByLayerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailDataByLayerResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailDataByLayerResponseBodyData) GetDataModule() []*DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	return s.DataModule
}

func (s *DescribeDomainDetailDataByLayerResponseBodyData) SetDataModule(v []*DescribeDomainDetailDataByLayerResponseBodyDataDataModule) *DescribeDomainDetailDataByLayerResponseBodyData {
	s.DataModule = v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainDetailDataByLayerResponseBodyDataDataModule struct {
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2
	Acc *int64 `json:"Acc,omitempty" xml:"Acc,omitempty"`
	// The bandwidth of IPv6 requests. Unit: bit/s.
	//
	// example:
	//
	// 21060178715.146667
	Bps *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The number of requests.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// - You can call this operation up to 20 times per second per account.
	//
	// - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
	//
	// **Time granularity**
	//
	// The following table describes the time granularity, the time period within which historical data is available, and the data delay, which vary with the maximum time range per query.
	//
	// | Time granularity | Maximum time range per query | Historical data available | Data delay |
	//
	// | ---------------- | ---------------------------- | ------------------------- | ---------- |
	//
	// | 5 minutes | 3 days | 93 days | 15 minutes |
	//
	// | 1 hour | 31 days | 186 days | 4 hours |
	//
	// | 1 days | 366 days | 366 days | 04:00 on the next day |
	//
	// example:
	//
	// 200:44349|206:1753719|304:45|403:1095|416:1|499:332
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 1
	Ipv6Acc *int64 `json:"Ipv6Acc,omitempty" xml:"Ipv6Acc,omitempty"`
	// The number of IPv6 requests.
	//
	// example:
	//
	// 0
	Ipv6Bps *float32 `json:"Ipv6Bps,omitempty" xml:"Ipv6Bps,omitempty"`
	// The amount of network traffic generated by IPv6 requests. Unit: bytes.
	//
	// example:
	//
	// 0
	Ipv6Qps *float32 `json:"Ipv6Qps,omitempty" xml:"Ipv6Qps,omitempty"`
	// The proportions of HTTP status codes.
	//
	// example:
	//
	// 141718944482
	Ipv6Traf *int64 `json:"Ipv6Traf,omitempty" xml:"Ipv6Traf,omitempty"`
	// The number of requests.
	//
	// example:
	//
	// 5998.47
	Qps *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The domain name.
	//
	// example:
	//
	// 2020-07-05T22:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The bandwidth of IPv6 requests. Unit: bit/s.
	//
	// example:
	//
	// 789756701818
	Traf *int64 `json:"Traf,omitempty" xml:"Traf,omitempty"`
}

func (s DescribeDomainDetailDataByLayerResponseBodyDataDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GetAcc() *int64 {
	return s.Acc
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GetBps() *float32 {
	return s.Bps
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GetHttpCode() *string {
	return s.HttpCode
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GetIpv6Acc() *int64 {
	return s.Ipv6Acc
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GetIpv6Bps() *float32 {
	return s.Ipv6Bps
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GetIpv6Qps() *float32 {
	return s.Ipv6Qps
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GetIpv6Traf() *int64 {
	return s.Ipv6Traf
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GetQps() *float32 {
	return s.Qps
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) GetTraf() *int64 {
	return s.Traf
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetAcc(v int64) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Acc = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetBps(v float32) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Bps = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetDomainName(v string) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetHttpCode(v string) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.HttpCode = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetIpv6Acc(v int64) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Ipv6Acc = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetIpv6Bps(v float32) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Ipv6Bps = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetIpv6Qps(v float32) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Ipv6Qps = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetIpv6Traf(v int64) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Ipv6Traf = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetQps(v float32) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Qps = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetTimeStamp(v string) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) SetTraf(v int64) *DescribeDomainDetailDataByLayerResponseBodyDataDataModule {
	s.Traf = &v
	return s
}

func (s *DescribeDomainDetailDataByLayerResponseBodyDataDataModule) Validate() error {
	return dara.Validate(s)
}
