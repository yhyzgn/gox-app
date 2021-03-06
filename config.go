// Copyright 2019 yhyzgn gox
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

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2019-11-26 15:10
// version: 1.0.0
// desc   :

package main

import (
	"github.com/yhyzgn/gox/ctx"
	"github.com/yhyzgn/gox/of/filter/cors"
	"net/http"

	"github.com/yhyzgn/gox/component/filter"
	"github.com/yhyzgn/gox/component/interceptor"
	testFilter "gox-app/filter"
	testInterceptor "gox-app/interceptor"
)

type Config struct {
}

func NewConfig() *Config {
	return new(Config)
}

func (c *Config) Context(ctx *ctx.GoXContext) {
	ctx.
		SetStaticDir("static").
		SetNotFoundHandler(func(writer http.ResponseWriter, request *http.Request) {
			http.Error(writer, "你的请求被外星人绑架啦~", http.StatusNotFound)
		})
}

func (c *Config) ConfigFilter(chain *filter.Chain) {
	chain.
		AddFilters("/", cors.NewXCorsFilter().AllowedHeaders("Token"), testFilter.NewTestFilter()).
		AddFilters("/api/*", testFilter.NewLogFilter()).
		Exclude("/api/log/*")
}

func (c *Config) ConfigInterceptor(register *interceptor.Register) {
	register.AddInterceptors("/", testInterceptor.NewTestInterceptor())
}
