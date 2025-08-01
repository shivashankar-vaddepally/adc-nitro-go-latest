/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package ns

/**
* Binding class showing the nsservicefunction that can be bound to nsservicepath.
*/
type Nsservicepathnsservicefunctionbinding struct {
	/**
	* List of service functions constituting the chain.
	*/
	Servicefunction string `json:"servicefunction,omitempty"`
	/**
	* The serviceindex of each servicefunction in path.
	*/
	Index int `json:"index,omitempty"`
	/**
	* Name for the Service path. Must begin with an ASCII alphanumeric or underscore (_) character, and must
		contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-)
		characters.
	*/
	Servicepathname string `json:"servicepathname,omitempty"`


}