import {Header} from "./header";
import {Sidebar} from "./sidebar";

export function Layout({children}) {
    return <>
        <div className="container-fluid p-0">
            <div className="row g-0">
                <div className="col-3">
                    <Sidebar/>
                </div>
                <div className="col-9">
                    <Header/>
                    <div className="page-body">
                        {children}
                    </div>
                </div>
            </div>
        </div>
    </>;
}
