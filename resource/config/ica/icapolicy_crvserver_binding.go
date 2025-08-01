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

package ica

/**
* Binding class showing the crvserver that can be bound to icapolicy.
*/
type Icapolicycrvserverbinding struct {
	/**
	* Location where policy is bound.
	*/
	Boundto string `json:"boundto,omitempty"`
	/**
	* Specifies the priority of the policy.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Indicates whether policy is bound or not.
	*/
	Activepolicy int `json:"activepolicy,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Name of the policy about which to display detailed information.
	*/
	Name string `json:"name,omitempty"`


}