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

package gslb

/**
* Statistics for GSLB domain resource.
*/

type Gslbdomainstats struct {
	/**
	* Name of the GSLB domain for which to display statistics. If you do not specify a name, statistics are shown for all configured GSLB
		domains.
	*/
	Name string `json:"name,omitempty"`
	Dnsrecordtype string `json:"dnsrecordtype,omitempty"`
	/**
	* Clear the statsistics / counters
	*/
	Clearstats string `json:"clearstats,omitempty"`
	Dnstotalqueries int `json:"dnstotalqueries,omitempty"`
	/**
	* Total number of DNS queries received.
	*/
	Dnsqueriesrate float64 `json:"dnsqueriesrate,omitempty"`
	Gslbdnsrectype string `json:"gslbdnsrectype,omitempty"`

}
