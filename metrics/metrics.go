// Copyright (C) 2015-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License,
// Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.

package metrics

type Metric struct {
	Key           string  `json:"name"`
	Value         float64 `json:"value"`
	Unit          string  `json:"unit"`
	AppGUID       string  `json:"app_guid"`
	InstanceIndex uint32  `json:"instance_index"`
}

type Metrics []Metric

func (m Metrics) Find(key string) (Metric, bool) {
	for _, metric := range m {
		if metric.Key == key {
			return metric, true
		}
	}

	return Metric{}, false
}
