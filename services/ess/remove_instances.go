package ess

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

func (client *Client) RemoveInstances(request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
	response = CreateRemoveInstancesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RemoveInstancesWithChan(request *RemoveInstancesRequest) (<-chan *RemoveInstancesResponse, <-chan error) {
	responseChan := make(chan *RemoveInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveInstances(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) RemoveInstancesWithCallback(request *RemoveInstancesRequest, callback func(response *RemoveInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveInstancesResponse
		var err error
		defer close(result)
		response, err = client.RemoveInstances(request)
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

type RemoveInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceId10         string `position:"Query" name:"InstanceId.10"`
	InstanceId12         string `position:"Query" name:"InstanceId.12"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	InstanceId11         string `position:"Query" name:"InstanceId.11"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceId19         string `position:"Query" name:"InstanceId.19"`
	InstanceId17         string `position:"Query" name:"InstanceId.17"`
	InstanceId18         string `position:"Query" name:"InstanceId.18"`
	InstanceId6          string `position:"Query" name:"InstanceId.6"`
	InstanceId15         string `position:"Query" name:"InstanceId.15"`
	InstanceId7          string `position:"Query" name:"InstanceId.7"`
	InstanceId16         string `position:"Query" name:"InstanceId.16"`
	InstanceId8          string `position:"Query" name:"InstanceId.8"`
	InstanceId13         string `position:"Query" name:"InstanceId.13"`
	InstanceId9          string `position:"Query" name:"InstanceId.9"`
	InstanceId14         string `position:"Query" name:"InstanceId.14"`
	InstanceId2          string `position:"Query" name:"InstanceId.2"`
	InstanceId3          string `position:"Query" name:"InstanceId.3"`
	InstanceId4          string `position:"Query" name:"InstanceId.4"`
	InstanceId5          string `position:"Query" name:"InstanceId.5"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	InstanceId1          string `position:"Query" name:"InstanceId.1"`
	InstanceId20         string `position:"Query" name:"InstanceId.20"`
}

type RemoveInstancesResponse struct {
	*responses.BaseResponse
	ScalingActivityId string `json:"ScalingActivityId" xml:"ScalingActivityId"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
}

func CreateRemoveInstancesRequest() (request *RemoveInstancesRequest) {
	request = &RemoveInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "RemoveInstances", "", "")
	return
}

func CreateRemoveInstancesResponse() (response *RemoveInstancesResponse) {
	response = &RemoveInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}