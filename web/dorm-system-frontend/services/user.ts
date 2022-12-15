import {fetchWrapper} from "./fetch_wrapper";


export const user = {
    login: login,
    register: register,
    getInfo: getInfo,
    setInfo: setInfo
};

type Credential = {
    username: string
    password: string
}

function genCredential(username: string, password: string) {
    const cred: Credential = {
        username: username,
        password: password
    };
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
    fetchWrapper.post("/register", cred)
        .then(data => console.log(data))
        .catch(err => console.log(err));
}

export type UserInfo = {
    id: number
    name: string
    gender: string
    studentNum: string
}

let info: UserInfo;

function setInfo(input: UserInfo) {
    info = input;
}

function getInfo() {
    return info;
}
