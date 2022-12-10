export const alertWrapper = {
    write: write,
    read: read
};

function write(msg: string) {
    const options: RequestInit = {
        method: "POST",
        body: msg
    }
    return fetch("/api/alert", options).then(resp => resp.json()).catch(err => console.log(err));
}

function read() {
    return fetch("/api/alert").then(resp => resp.json()).catch(err => console.log(err));
}
