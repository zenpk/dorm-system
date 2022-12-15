import {useEffect, useState} from "react";
import Link from "next/link";
import {fetchWrapper} from "../services/fetch_wrapper";
import {alertWrapper} from "../services/alert_wrapper";
import {user} from "../services/user";

export function Header() {
    const [name, setName] = useState("");
    useEffect(() => {
        fetchWrapper.get("/user/info")
            .then(data => {
                    if (data.err.code === process.env.errOK) {
                        setName(data.user.name);
                        user.setInfo(data.user);
                    }
                }
            )
            .catch(err => console.log(err));
    }, []);

    return <>
        <div className="page-header">
            <div className="row justify-content-between">
                <div className="col-auto">
                    <span className="font-code mfs-6">Source Code on ={">"}&nbsp;
                        <a className="a-dec" href="https://github.com/zenpk/dorm-system" tabIndex={0}>
                            GitHub
                        </a>
                    </span>
                </div>
                <div className="col-auto">
                    {name?.length > 0 ?
                        <span className="mfs-6">Welcome: <Link href="/user/info" className="a-dec">{name}</Link></span>
                        : <span className="mfs-6"> <Link href="/user/login" className="a-dec">Login</Link></span>
                    }
                </div>
            </div>
        </div>
    </>;
}
