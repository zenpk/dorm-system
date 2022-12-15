import {Layout} from "../../components"
import {useEffect, useState} from "react";
import {fetchWrapper} from "../../services/fetch_wrapper";
import {alertWrapper} from "../../services/alert_wrapper";

type BuildingInfo = {
    num: string
    info: string
    imgUrl: string
}

type RemainCnt = {
    all: string
    [key: string]: string
}

export default function Info() {
    const [info, setInfo] = useState<[BuildingInfo]>();
    const [rem, setRem] = useState<RemainCnt>();

    useEffect(() => {
        fetchWrapper.get("/buildings")
            .then(data=>{
                if (data.err.code !== process.env.errOK) {
                    alertWrapper.write(data.err.msg);
                }else{
                    setInfo(data.infos);
                }
            })
    }, []);

    useEffect(() => {
        fetchWrapper.get("/remain-cnt")
            .then(data => {
                setRem(data.remainCnt);
            })
    }, []);

    return <Layout>
        <h1>Building information</h1>
        <h2>Total remaining beds count: {rem?.all}</h2>
        {
            info?.map((bd, i) => {
                return <div className="row mt-5 p-0" key={i}>
                    <div className="col">
                        <img className="image-70 center"
                             src={bd.imgUrl}
                             alt="building image"/>
                    </div>
                    <div className="col">
                        <ul>
                            <li>Building number: {bd.num}</li>
                            <li>Information: {bd.info}</li>
                            <li>Remaining beds count: {rem ? rem[`${bd.num}`] : "NaN"}</li>
                        </ul>
                    </div>
                </div>
            })
        }
    </Layout>
}
