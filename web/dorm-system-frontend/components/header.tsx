import {useEffect, useState} from "react";
import Link from "next/link";

export function Header() {
    const [name, setName] = useState("");
    useEffect(() => {
        let allCookie = document.cookie;
        console.log(allCookie);
    }, [])

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
                    {name.length > 0 ?
                        <span className="mfs-6">Welcome: <Link href="/user/info" className="a-dec">{name}</Link></span>
                        : <span className="mfs-6"> <Link href="/user/login" className="a-dec">Login</Link></span>
                    }
                </div>
            </div>
        </div>
    </>;
}
