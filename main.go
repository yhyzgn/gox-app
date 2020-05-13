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
// time   : 2019-11-25 14:47
// version: 1.0.0
// desc   :

package main

import (
	"fmt"
	"github.com/yhyzgn/gog"
	"github.com/yhyzgn/gox"
	"gox-app/api/controller"
	"gox-app/config"
	"net/http"
	"os"
)

const (
	port = 8080
)

func init() {
	env := os.Getenv("ENV")
	if env == "prod" {
		// 生产环境
		gog.SetFormatter(gog.NewJSONFormatter())
		gog.SetLevel(gog.INFO)
	}
	gog.Async(true)
}

func main() {
	x := gox.NewGoX()

	web := &config.WebConfig{}
	err := x.Load("config/config.yml", web)
	fmt.Println(web, err)

	initWeb(x)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", web.Server.Port),
		Handler: x,
	}

	x.Run(server)
}

func initWeb(r *gox.GoX) {
	r.
		Configure(NewConfig()).
		Mapping("/api/normal", new(controller.NormalController)).
		Mapping("/api/param", new(controller.ParamController)).
		Mapping("/api/log", new(controller.LogController))
}
