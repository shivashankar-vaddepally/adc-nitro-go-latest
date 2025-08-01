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

package appflow

/**
* Binding class showing the appflowpolicy that can be bound to appflowglobal.
*/
type Appflowglobalappflowpolicybinding struct {
	/**
	* Name of the AppFlow policy.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Global bind point for which to show detailed information about the policies bound to the bind point.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Specifies the priority of the policy.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of policy label to invoke. Specify vserver for a policy label associated with a virtual server, or policylabel for a user-defined policy label.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the label to invoke if the current policy evaluates to TRUE.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* The number of policies bound to the bindpoint.
	*/
	Numpol int `json:"numpol,omitempty"`
	/**
	* Flow type of the bound AppFlow policy.
	*/
	Flowtype int `json:"flowtype,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`


}