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

package cluster

/**
* Statistics for cluster instance resource.
*/

type Clusterinstancestats struct {
	/**
	* ID of the cluster instance for which to display statistics.
	*/
	Clid int `json:"clid,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Clbkplanetx int `json:"clbkplanetx,omitempty"`
	/**
	* Traffic transmitted from backplane (in mbits)
	*/
	Clbkplanetxrate float64 `json:"clbkplanetxrate,omitempty"`
	Clbkplanerx int `json:"clbkplanerx,omitempty"`
	/**
	* Traffic received on backplane (in mbits)
	*/
	Clbkplanerxrate float64 `json:"clbkplanerxrate,omitempty"`
	Clnumnodes int `json:"clnumnodes,omitempty"`
	Clcurstatus string `json:"clcurstatus,omitempty"`
	Clviewleader string `json:"clviewleader,omitempty"`
	Totsteeredpkts int `json:"totsteeredpkts,omitempty"`
	/**
	* Total number of packets steered on the cluster backplane.
	*/
	Steeredpktsrate float64 `json:"steeredpktsrate,omitempty"`
	Numdfddroppkts int `json:"numdfddroppkts,omitempty"`
	Totpropagationtimeout int `json:"totpropagationtimeout,omitempty"`

}
