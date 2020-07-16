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
	"context"
	"net/http"
	"sync/atomic"
)

var _ Client = (*refreshableClientImpl)(nil)

type refreshableClientImpl struct {
	val    *atomic.Value
	params []ClientParam
}

func (r refreshableClientImpl) Do(ctx context.Context, params ...RequestParam) (*http.Response, error) {
	return r.val.Load().(Client).Do(ctx, params...)
}

func (r refreshableClientImpl) Get(ctx context.Context, params ...RequestParam) (*http.Response, error) {
	return r.val.Load().(Client).Get(ctx, params...)
}

func (r refreshableClientImpl) Head(ctx context.Context, params ...RequestParam) (*http.Response, error) {
	return r.val.Load().(Client).Head(ctx, params...)
}

func (r refreshableClientImpl) Post(ctx context.Context, params ...RequestParam) (*http.Response, error) {
	return r.val.Load().(Client).Post(ctx, params...)
}

func (r refreshableClientImpl) Put(ctx context.Context, params ...RequestParam) (*http.Response, error) {
	return r.val.Load().(Client).Put(ctx, params...)
}

func (r refreshableClientImpl) Delete(ctx context.Context, params ...RequestParam) (*http.Response, error) {
	return r.val.Load().(Client).Delete(ctx, params...)
}

func (r refreshableClientImpl) Update(config ClientConfig) {
	params := append(r.params, WithConfig(config))
	client, err := NewClient(params...)
	if err != nil {
		// log or something
		return
	}
	r.val.Store(client)
}
