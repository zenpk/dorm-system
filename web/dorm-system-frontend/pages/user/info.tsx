import {Layout} from "../../components"
import {useEffect, useState} from "react";
import {fetchWrapper} from "../../services/fetch_wrapper";
import Link from "next/link";
import {useRouter} from "next/router";
import {alertWrapper} from "../../services/alert_wrapper";


export type UserInfo = {
    id: number,
    name: string,
    studentNum: string,
    gender: string
};

export default function Info() {

    const [info, setInfo] = useState<UserInfo>();

    useEffect(() => {
        fetchWrapper.get("/user/info")
            .then(data => {
                    if (data.err.code === process.env.errNotLogin) {
                        alertWrapper.write("You need to login first");
                    } else if (data.err.code !== process.env.errOK) {
                        alertWrapper.write(data.err.msg);
                    } else {
                        setInfo(data.user);
                    }
                }
            )
            .catch(err => console.log(err));
    }, []);

    const router = useRouter();

    function logout() {
        fetchWrapper.get("/user/logout");
        router.push("/user/login");
    }

    return <Layout>
        <h1>My information</h1>
        <table className={"table table-striped table-bordered"}>
            <tbody>
            <tr>
                <th>Id</th>
                <td>{info?.id}</td>
            </tr>
            <tr>
                <th>Name</th>
                <td>{info?.name}</td>
            </tr>
            <tr>
                <th>Student number</th>
                <td>{info?.studentNum}</td>
            </tr>
            <tr>
                <th>Gender</th>
                <td>{info?.gender}</td>
            </tr>
            </tbody>
        </table>
        <Link href={{
            pathname: "/user/edit",
            query: info,
        }} className="my-button-big fw-bold">
            Edit
        </Link>
        &nbsp;&nbsp;
        <a href="#" onClick={logout} className="my-button-big fw-bold">
            Logout
        </a>
    </Layout>
}
