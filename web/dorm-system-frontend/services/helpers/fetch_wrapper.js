import getConfig from "next/config";

const {publicRuntimeConfig} = getConfig();

export const fetchWrapper = {
    get,
    post,
    put,
    delete: _delete
};

function get(url) {
    const options = {
        method: "GET",
        credentials: "include"
    };
    return fetch(publicRuntimeConfig.url + url, options);
}

function post(url, body) {
    const options = {
        method: "POST",
        headers: {"Content-Type": "application/json;charset=UTF-8"},
        credentials: "include",
        body: JSON.stringify(body)
    };
    return fetch(publicRuntimeConfig.url + url, options);
}

function put(url, body) {
    const options = {
        method: "PUT",
        headers: {"Content-Type": "application/json;charset=UTF-8"},
        credentials: "include",
        body: JSON.stringify(body)
    };
    return fetch(publicRuntimeConfig.url + url, options);
}

// prefixed with underscored because delete is a reserved word in javascript
function _delete(url) {
    const options = {
        method: "DELETE",
        credentials: "include"
    };
    return fetch(publicRuntimeConfig.url + url, options);
}
