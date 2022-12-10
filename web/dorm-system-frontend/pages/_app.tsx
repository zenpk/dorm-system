// Copyright 2022 zenpk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import Head from "next/head"
import type { AppProps } from 'next/app'
import "bootstrap/dist/css/bootstrap.css"

import "../styles/sidebar.css"
import "../styles/mybasic.css"
import "../styles/my.css"
import "../styles/mysidebar.css"
import Script from "next/script";

export default function MyApp({ Component, pageProps }: AppProps) {
    return (
        <>
            <Head>
                <title>{"Dorm System"}</title>
                <meta charSet="utf-8"/>
                <meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=no"/>
            </Head>
            <Script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/js/bootstrap.bundle.min.js"
                    integrity="sha384-pprn3073KE6tl6bjs2QrFaJGz5/SUsLqktiwsUTF55Jfv3qYSDhgCecCxMW52nD2"
                    crossOrigin="anonymous"/>
            <Component {...pageProps} />
        </>
    )
}
