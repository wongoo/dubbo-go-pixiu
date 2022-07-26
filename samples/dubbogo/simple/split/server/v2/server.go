/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	routers := []string{"/user", "/user/pixiu", "/prefix", "/health"}

	for _, router := range routers {
		msg := router[strings.LastIndex(router, "/")+1:]
		http.HandleFunc(router, func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(fmt.Sprintf(`{"server": "v2","message":"%s","status":200}`, msg)))
		})
	}

	remote()

	log.Println("Starting sample server ...")
	log.Fatal(http.ListenAndServe(":1316", nil))
}

func remote() {
	http.HandleFunc("/remote", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"keys":[{"kty":"RSA","e":"AQAB","kid":"ee8d626d","n":"gRda5b0pkgTytDuLrRnNSYhvfMIyM0ASq2ZggY4dVe12JV8N7lyXilyqLKleD-2lziivvzE8O8CdIC2vUf0tBD7VuMyldnZruSEZWCuKJPdgKgy9yPpShmD2NyhbwQIAbievGMJIp_JMwz8MkdY5pzhPECGNgCEtUAmsrrctP5V8HuxaxGt9bb-DdPXkYWXW3MPMSlVpGZ5GiIeTABxqYNG2MSoYeQ9x8O3y488jbassTqxExI_4w9MBQBJR9HIXjWrrrenCcDlMY71rzkbdj3mmcn9xMq2vB5OhfHyHTihbUPLSm83aFWSuW9lE7ogMc93XnrB8evIAk6VfsYlS9Q"},{"kty":"EC","crv":"P-256","kid":"711d48d1","x":"tfXCoBU-wXemeQCkME1gMZWK0-UECCHIkedASZR0t-Q","y":"9xzYtnKQdiQJHCtGwpZWF21eP1fy5x4wC822rCilmBw"},{"kty":"EC","crv":"P-384","kid":"d52c9829","x":"tFx6ev6eLs9sNfdyndn4OgbhV6gPFVn7Ul0VD5vwuplJLbIYeFLI6T42tTaE5_Q4","y":"A0gzB8TqxPX7xMzyHH_FXkYG2iROANH_kQxBovSeus6l_QSyqYlipWpBy9BhY9dz"},{"kty":"RSA","e":"AQAB","kid":"ecac72e5","n":"nLbnTvZAUxdmuAbDDUNAfha6mw0fri3UpV2w1PxilflBuSnXJhzo532-YQITogoanMjy_sQ8kHUhZYHVRR6vLZRBBbl-hP8XWiCe4wwioy7Ey3TiIUYfW-SD6I42XbLt5o-47IR0j5YDXxnX2UU7-UgR_kITBeLDfk0rSp4B0GUhPbP5IDItS0MHHDDS3lhvJomxgEfoNrp0K0Fz_s0K33hfOqc2hD1tSkX-3oDTQVRMF4Nxax3NNw8-ahw6HNMlXlwWfXodgRMvj9pcz8xUYa3C5IlPlZkMumeNCFx1qds6K_eYcU0ss91DdbhhE8amRX1FsnBJNMRUkA5i45xkOIx15rQN230zzh0p71jvtx7wYRr5pdMlwxV0T9Ck5PCmx-GzFazA2X6DJ0Xnn1-cXkRoZHFj_8Mba1dUrNz-NWEk83uW5KT-ZEbX7nzGXtayKWmGb873a8aYPqIsp6bQ_-eRBd8TDT2g9HuPyPr5VKa1p33xKaohz4DGy3t1Qpy3UWnbPXUlh5dLWPKz-TcS9FP5gFhWVo-ZhU03Pn6P34OxHmXGWyQao18dQGqzgD4e9vY3rLhfcjVZJYNlWY2InsNwbYS-DnienPf1ws-miLeXxNKG3tFydoQzHwyOxG6Wc-HBfzL_hOvxINKQamvPasaYWl1LWznMps6elKCgKDc"},{"kty":"EC","crv":"P-521","kid":"c570888f","x":"AHNpXq0J7rikNRlwhaMYDD8LGVAVJzNJ-jEPksUIn2LB2LCdNRzfAhgbxdQcWT9ktlc9M1EhmTLccEqfnWdGL9G1","y":"AfHPUW3GYzzqbTczcYR0nYMVMFVrYsUxv4uiuSNV_XRN3Jf8zeYbbOLJv4S3bUytO7qHY8bfZxPxR9nn3BBTf5ol"}]}`))
	})
}
