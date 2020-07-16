// Copyright (c) 2020 Palantir Technologies. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httpclient

import (
	"github.com/palantir/witchcraft-go-server/witchcraft/refreshable"
)

type RefreshableClientConfig interface {
	refreshable.Refreshable

	CurrentClientConfig() ClientConfig
	MapClientConfig(func(ClientConfig) interface{}) refreshable.Refreshable
	SubscribeToClientConfig(consumer func(ClientConfig)) (unsubscribe func())
}

type RefreshingClientConfig struct {
	refreshable.Refreshable
}

func (r RefreshingClientConfig) CurrentClientConfig() ClientConfig {
	return r.Refreshable.Current().(ClientConfig)
}

func (r RefreshingClientConfig) MapClientConfig(f func(config ClientConfig) interface{}) refreshable.Refreshable {
	return r.Refreshable.Map(func(i interface{}) interface{} {
		return f(i.(ClientConfig))
	})
}

func (r RefreshingClientConfig) SubscribeToClientConfig(consumer func(config ClientConfig)) func() {
	return r.Subscribe(func(i interface{}) {
		consumer(i.(ClientConfig))
	})
}
