package cdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeCdnDomainConfigs invokes the cdn.DescribeCdnDomainConfigs API synchronously
// api document: https://help.aliyun.com/api/cdn/describecdndomainconfigs.html
func (client *Client) DescribeCdnDomainConfigs(request *DescribeCdnDomainConfigsRequest) (response *DescribeCdnDomainConfigsResponse, err error) {
	response = CreateDescribeCdnDomainConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnDomainConfigsWithChan invokes the cdn.DescribeCdnDomainConfigs API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecdndomainconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnDomainConfigsWithChan(request *DescribeCdnDomainConfigsRequest) (<-chan *DescribeCdnDomainConfigsResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnDomainConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnDomainConfigs(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeCdnDomainConfigsWithCallback invokes the cdn.DescribeCdnDomainConfigs API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecdndomainconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnDomainConfigsWithCallback(request *DescribeCdnDomainConfigsRequest, callback func(response *DescribeCdnDomainConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnDomainConfigsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnDomainConfigs(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeCdnDomainConfigsRequest is the request struct for api DescribeCdnDomainConfigs
type DescribeCdnDomainConfigsRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	FunctionNames string           `position:"Query" name:"FunctionNames"`
}

// DescribeCdnDomainConfigsResponse is the response struct for api DescribeCdnDomainConfigs
type DescribeCdnDomainConfigsResponse struct {
	*responses.BaseResponse
	RequestId     string                                  `json:"RequestId" xml:"RequestId"`
	DomainConfigs DomainConfigsInDescribeCdnDomainConfigs `json:"DomainConfigs" xml:"DomainConfigs"`
}

// CreateDescribeCdnDomainConfigsRequest creates a request to invoke DescribeCdnDomainConfigs API
func CreateDescribeCdnDomainConfigsRequest() (request *DescribeCdnDomainConfigsRequest) {
	request = &DescribeCdnDomainConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeCdnDomainConfigs", "", "")
	return
}

// CreateDescribeCdnDomainConfigsResponse creates a response to parse from DescribeCdnDomainConfigs response
func CreateDescribeCdnDomainConfigsResponse() (response *DescribeCdnDomainConfigsResponse) {
	response = &DescribeCdnDomainConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
