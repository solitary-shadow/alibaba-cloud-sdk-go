package qualitycheck

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

// ResultInfo is a nested struct in qualitycheck response
type ResultInfo struct {
	Comments        string                `json:"Comments" xml:"Comments"`
	Status          int                   `json:"Status" xml:"Status"`
	RuleName        string                `json:"RuleName" xml:"RuleName"`
	HitId           string                `json:"HitId" xml:"HitId"`
	ErrorMessage    string                `json:"ErrorMessage" xml:"ErrorMessage"`
	ReviewStatus    int                   `json:"ReviewStatus" xml:"ReviewStatus"`
	CreateTime      string                `json:"CreateTime" xml:"CreateTime"`
	AsrMsg          string                `json:"AsrMsg" xml:"AsrMsg"`
	ReviewResult    int                   `json:"ReviewResult" xml:"ReviewResult"`
	Tid             string                `json:"Tid" xml:"Tid"`
	Resolver        string                `json:"Resolver" xml:"Resolver"`
	TaskId          string                `json:"TaskId" xml:"TaskId"`
	Rid             int64                 `json:"Rid" xml:"Rid"`
	TaskName        string                `json:"TaskName" xml:"TaskName"`
	Reviewer        string                `json:"Reviewer" xml:"Reviewer"`
	Score           int                   `json:"Score" xml:"Score"`
	HandScoreIdList HandScoreIdList       `json:"HandScoreIdList" xml:"HandScoreIdList"`
	Agent           Agent                 `json:"Agent" xml:"Agent"`
	Recording       Recording             `json:"Recording" xml:"Recording"`
	Rules           RulesInUploadDataSync `json:"Rules" xml:"Rules"`
	AsrResult       AsrResultInGetResult  `json:"AsrResult" xml:"AsrResult"`
	HitResult       HitResult             `json:"HitResult" xml:"HitResult"`
}
