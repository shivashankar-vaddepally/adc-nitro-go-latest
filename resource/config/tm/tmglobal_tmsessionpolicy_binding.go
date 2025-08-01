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

package tm

/**
* Binding class showing the tmsessionpolicy that can be bound to tmglobal.
*/
type Tmglobaltmsessionpolicybinding struct {
	/**
	* The name of the policy.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* The priority of the policy.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Bound policy type
	*/
	Bindpolicytype int `json:"bindpolicytype,omitempty"`
	/**
	* Indicates that a variable is a built-in (SYSTEM INTERNAL) type.
	*/
	Builtin []string `json:"builtin,omitempty"`
	/**
	* The feature to be checked while applying this config
	*/
	Feature string `json:"feature,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`


}