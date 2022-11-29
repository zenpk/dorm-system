import {Layout} from "../../components"
import {useEffect, useState} from "react";
import {fetchWrapper} from "../../services/helpers/fetch_wrapper";
import {useRouter} from "next/router";

export default function Info() {

    const [alert, setAlert] = useState("");
    const [team, setTeam] = useState({});

    useEffect(() => {
        fetchWrapper.get("/team/")
            .then(resp => resp.json())
            .then(data => {
                    console.log(data);
                    if (data.err.code !== process.env.errOK) {
                        setAlert("You haven't joined any team");
                    } else {
                        setTeam(data.team);
                    }
                }
            )
    }, []);

    const router = useRouter();

    function leave() {
        fetchWrapper.delete("/team/leave")
            .then(resp => resp.json())
            .then(data => console.log(data));
        router.reload();
    }

    if (alert.length > 0) {
        return <Layout>
            <h2>{alert}</h2>
        </Layout>;
    } else {
        return <Layout>
            <h1>My team information</h1>
            <table className={"table table-striped table-bordered"}>
                <tbody>
                <tr>
                    <th>Id</th>
                    <td>{team?.id}</td>
                </tr>
                <tr>
                    <th>Code</th>
                    <td>{team?.code}</td>
                </tr>
                <tr>
                    <th>Gender</th>
                    <td>{team?.gender}</td>
                </tr>
                <tr>
                    <th>Owner</th>
                    <td className={"p-0"}>
                        <table className={"table table-striped table-bordered m-0"}>
                            <tbody>
                            <tr>
                                <th className={"w-50"}>Id</th>
                                <td className={"w-50"}>{team?.owner?.id}</td>
                            </tr>
                            <tr>
                                <th className={"w-50"}>Name</th>
                                <td className={"w-50"}>{team?.owner?.name}</td>
                            </tr>
                            <tr>
                                <th className={"w-50"}>Student number</th>
                                <td className={"w-50"}>{team?.owner?.studentNum}</td>
                            </tr>
                            </tbody>
                        </table>
                    </td>
                </tr>
                <tr>
                    <th>Members</th>
                    <td className={"p-0"}>
                        {
                            team?.members?.map((m, i) => {
                                return <table className={"table table-striped table-bordered my-3"} key={i}>
                                    <tbody>
                                    <tr>
                                        <th className={"w-50"}>Id</th>
                                        <td className={"w-50"}>{m?.id}</td>
                                    </tr>
                                    <tr>
                                        <th className={"w-50"}>Name</th>
                                        <td className={"w-50"}>{m?.name}</td>
                                    </tr>
                                    <tr>
                                        <th className={"w-50"}>Student number</th>
                                        <td className={"w-50"} >{m?.studentNum}</td>
                                    </tr>
                                    </tbody>
                                </table>;
                            })
                        }
                    </td>
                </tr>
                </tbody>
            </table>
            <a href="#" onClick={leave} className="my-button-big fw-bold">
                Leave
            </a>
        </Layout>
    }
}
