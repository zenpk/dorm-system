import {Header} from "./header";
import {Sidebar} from "./sidebar";
import React, {useEffect, useState} from "react";
import {alertWrapper} from "../services/alert_wrapper";

export function Layout({children}: { children: React.ReactNode }) {
    const [alertMsg, setAlertMsg] = useState("");
    useEffect(() => {
        alertWrapper.read().then(resp => setAlertMsg(resp.msg));
    }, []);

    return <>
        <div className="container-fluid p-0">
            <div className="row g-0">
                <div className="col-3">
                    <Sidebar/>
                </div>
                <div className="col-9">
                    <Header/>
                    <div className="page-body">
                        {alertMsg.length > 0 ?
                            <div className="alert alert-danger fs-6 py-1 px-3" role="alert">{alertMsg}</div> : null}
                        {children}
                    </div>
                </div>
            </div>
        </div>
    </>;
}
