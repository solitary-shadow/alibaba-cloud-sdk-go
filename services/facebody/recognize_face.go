package facebody

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

// RecognizeFace invokes the facebody.RecognizeFace API synchronously
// api document: https://help.aliyun.com/api/facebody/recognizeface.html
func (client *Client) RecognizeFace(request *RecognizeFaceRequest) (response *RecognizeFaceResponse, err error) {
	response = CreateRecognizeFaceResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeFaceWithChan invokes the facebody.RecognizeFace API asynchronously
// api document: https://help.aliyun.com/api/facebody/recognizeface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeFaceWithChan(request *RecognizeFaceRequest) (<-chan *RecognizeFaceResponse, <-chan error) {
	responseChan := make(chan *RecognizeFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeFace(request)
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

// RecognizeFaceWithCallback invokes the facebody.RecognizeFace API asynchronously
// api document: https://help.aliyun.com/api/facebody/recognizeface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeFaceWithCallback(request *RecognizeFaceRequest, callback func(response *RecognizeFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeFaceResponse
		var err error
		defer close(result)
		response, err = client.RecognizeFace(request)
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

// RecognizeFaceRequest is the request struct for api RecognizeFace
type RecognizeFaceRequest struct {
	*requests.RpcRequest
	ImageType requests.Integer `position:"Body" name:"ImageType"`
	ImageURL  string           `position:"Body" name:"ImageURL"`
}

// RecognizeFaceResponse is the response struct for api RecognizeFace
type RecognizeFaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeFaceRequest creates a request to invoke RecognizeFace API
func CreateRecognizeFaceRequest() (request *RecognizeFaceRequest) {
	request = &RecognizeFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "RecognizeFace", "facebody", "openAPI")
	return
}

// CreateRecognizeFaceResponse creates a response to parse from RecognizeFace response
func CreateRecognizeFaceResponse() (response *RecognizeFaceResponse) {
	response = &RecognizeFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}