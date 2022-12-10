import getConfig from "next/config";

const {publicRuntimeConfig} = getConfig();

export const fetchWrapper = {
    get: get,
    post: post,
    put: put,
    delete: _delete
};

function get(url: string) {
    const options: RequestInit = {
        method: "GET",
        credentials: "include"
    };
    return fetch(publicRuntimeConfig.url + url, options).then(handleResponse);
}

function post(url: string, body: any) {
    const options: RequestInit = {
        method: "POST",
        headers: {"Content-Type": "application/json;charset=UTF-8"},
        credentials: "include",
        body: JSON.stringify(body)
    };
    return fetch(publicRuntimeConfig.url + url, options).then(handleResponse);
}

function put(url: string, body: any) {
    const options: RequestInit = {
        method: "PUT",
        headers: {"Content-Type": "application/json;charset=UTF-8"},
        credentials: "include",
        body: JSON.stringify(body)
    };
    return fetch(publicRuntimeConfig.url + url, options).then(handleResponse);
}

// prefixed with underscored because delete is a reserved word in javascript
function _delete(url: string) {
    const options: RequestInit = {
        method: "DELETE",
        credentials: "include"
    };
    return fetch(publicRuntimeConfig.url + url, options).then(handleResponse);
}

function handleResponse(resp: Response) {
    return resp.json().catch(err => console.log(err));
}
