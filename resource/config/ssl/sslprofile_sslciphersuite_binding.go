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

package ssl

/**
* Binding class showing the sslciphersuite that can be bound to sslprofile.
*/
type Sslprofilesslciphersuitebinding struct {
	/**
	* The cipher group/alias/individual cipher configuration
	*/
	Ciphername string `json:"ciphername,omitempty"`
	/**
	* cipher priority
	*/
	Cipherpriority int `json:"cipherpriority,omitempty"`
	/**
	* The cipher suite description.
	*/
	Description string `json:"description,omitempty"`
	/**
	* Name of the SSL profile.
	*/
	Name string `json:"name,omitempty"`


}