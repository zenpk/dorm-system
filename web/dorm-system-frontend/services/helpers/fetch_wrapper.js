import getConfig from "next/config";

const {publicRuntimeConfig} = getConfig();

export const fetchWrapper = {
    get: get,
    post: post,
    put: put,
    delete: _delete
};

function get(url) {
    const options = {
        method: "GET",
        credentials: "include"
    };
    return fetch(publicRuntimeConfig.url + url, options).then(handleResponse);
}

function post(url, body) {
    const options = {
        method: "POST",
        headers: {"Content-Type": "application/json;charset=UTF-8"},
        credentials: "include",
        body: JSON.stringify(body)
    };
    return fetch(publicRuntimeConfig.url + url, options).then(handleResponse);
}

function put(url, body) {
    const options = {
        method: "PUT",
        headers: {"Content-Type": "application/json;charset=UTF-8"},
        credentials: "include",
        body: JSON.stringify(body)
    };
    return fetch(publicRuntimeConfig.url + url, options).then(handleResponse);
}

// prefixed with underscored because delete is a reserved word in javascript
function _delete(url) {
    const options = {
        method: "DELETE",
        credentials: "include"
    };
    return fetch(publicRuntimeConfig.url + url, options).then(handleResponse);
}

function handleResponse(resp) {
    return resp.json().catch(err => console.log(err));
}
