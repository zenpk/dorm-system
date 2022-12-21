/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: false
}

module.exports = {
    ...nextConfig,
    publicRuntimeConfig: {
        // url: "http://127.0.0.1:8080"
        url: "http://172.17.0.1:8080" // Docker on Linux
    },
    env: {
        errOK: 2000,
        errUnknown: 5900,
        errNotLogin: 4001,
        errNoPermission: 4002,
        errInputHeader: 4003,
        errInputBody: 4004,
        errParseToken: 5201,
        errServiceConn: 5301,
        errLogic: 5501,
        errCacheConn: 5601,
        errNoCache: 5602,
        errDBConn: 5701,
        errNoRecord: 5702,
        errDuplicatedRecord: 5703,
        errTypeConv: 5901,
        errGenJWT: 5903,
        errParseJWT: 5903
    }
}
