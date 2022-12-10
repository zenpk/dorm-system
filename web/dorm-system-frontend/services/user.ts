import {fetchWrapper} from "./fetch_wrapper";

export const user = {
    login: login,
    register: register
}

type Credential = {
    username: string,
    password: string
}

function genCredential(username: string, password: string) {
    const cred: Credential = {
        username: username,
        password: password
    }
    return cred;
}

function login(username: string, password: string) {
    const cred = genCredential(username, password);
    fetchWrapper.post("/login", cred)
        .then(data => console.log(data))
        .catch(err => console.log(err));
}

function register(username: string, password: string) {
    const cred = genCredential(username, password);


}