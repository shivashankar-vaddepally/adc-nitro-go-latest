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

package system

/**
* Binding class showing the systemuser that can be bound to systemgroup.
*/
type Systemgroupsystemuserbinding struct {
	/**
	* The system user.
	*/
	Username string `json:"username,omitempty"`
	/**
	* Name of the system group.
	*/
	Groupname string `json:"groupname,omitempty"`


}