import {Layout} from "../../components";
import {useEffect, useState} from "react";
import {fetchWrapper} from "../../services/fetch_wrapper";
import {useRouter} from "next/router";
import {alertWrapper} from "../../services/alert_wrapper";
import {user} from "../../services/user";

type Member = {
    id: number
    name: string
    studentNum: string
}

type Team = {
    id: number
    code: string
    gender: string
    owner: Member
    members: Member[]
}

export default function Info() {

    const [team, setTeam] = useState<Team>();

    useEffect(() => {
        fetchWrapper.get("/team")
            .then(data => {
                    if (data.err.code !== process.env.errOK) {
                        alertWrapper.write(data.err.msg);
                    } else {
                        setTeam(data.team);
                    }
                }
            );
    }, []);

    const router = useRouter();

    function leave() {
        fetchWrapper.delete("/team/leave")
            .then(data => {
                if (data.err.code !== process.env.errOK) {
                    alertWrapper.write(data.err.msg);
                }
            });
        router.reload();
    }

    function transfer(newId: number) {
        const postData = {
            newOwnerId: newId
        };
        fetchWrapper.post("/team/transfer", postData)
            .then(data => {
                if (data.err.code !== process.env.errOK) {
                    alertWrapper.write(data.err.msg);
                }
            });
        router.reload();
    }

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
                            <th className="w-50">Id</th>
                            <td className="w-50">{team?.owner?.id}</td>
                        </tr>
                        <tr>
                            <th className="w-50">Name</th>
                            <td className="w-50">{team?.owner?.name}</td>
                        </tr>
                        <tr>
                            <th className="w-50">Student number</th>
                            <td className="w-50">{team?.owner?.studentNum}</td>
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
                                    <th className="w-50">Id</th>
                                    <td className="w-50">{m?.id}</td>
                                </tr>
                                <tr>
                                    <th className="w-50">Name</th>
                                    <td className="w-50">{m?.name}</td>
                                </tr>
                                <tr>
                                    <th className="w-50">Student number</th>
                                    <td className="w-50">{m?.studentNum}</td>
                                </tr>
                                {user.getInfo().id == team?.owner?.id ?
                                    <tr>
                                        <th className="w-50">Ownership</th>
                                        <td className="w-50">
                                            <button className="my-button" onClick={() => transfer(m?.id)}>transfer to
                                                                                                          this user
                                            </button>
                                        </td>
                                    </tr>
                                    : null}
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
    </Layout>;
}
