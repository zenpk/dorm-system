import {Layout} from "../../components"
import {useEffect, useState} from "react";
import {fetchWrapper} from "../../services/helpers/fetch_wrapper";

export default function Info() {
    const [info, setInfo] = useState([]);
    const [rem, setRem] = useState({});

    useEffect(() => {
        fetchWrapper.get("/buildings")
            .then(data => {
                setInfo(data.infos);
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
        <h2>Total remaining beds count: {rem.all}</h2>
        {
            info.map((bd, i) => {
                return <div className={"row mt-5 p-0"} key={i}>
                    <div className={"col"}>
                        <img className="image-70 center"
                             src={bd.imgUrl}
                        />
                    </div>
                    <div className={"col"}>
                        <ul>
                            <li>Building number: {bd.num}</li>
                            <li>Information: {bd.info}</li>
                            <li>Remaining beds count: {rem[`${bd.num}`]}</li>
                        </ul>
                    </div>
                </div>
            })
        }
    </Layout>
}
