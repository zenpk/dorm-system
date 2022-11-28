import {Layout} from "../../components"
import {useEffect, useState} from "react";
import {fetchWrapper} from "../../services/helpers/fetch_wrapper";
import {useRouter} from "next/router";

export default function Info() {

    const [alert, setAlert] = useState("");
    const [orders, setOrders] = useState([]);

    useEffect(() => {
        fetchWrapper.get("/order/")
            .then(resp => resp.json())
            .then(data => {
                    if (data.orders === undefined) {
                        setAlert("You haven't submitted any order");
                    } else {
                        setOrders(data.orders);
                    }
                }
            )
    }, []);

    const router = useRouter();

    function _delete() {
        fetchWrapper.delete("/order/delete");
        router.reload();
    }

    if (alert.length > 0) {
        return <Layout>
            <h2>{alert}</h2>
        </Layout>;
    } else {
        return <Layout>
            <h1>My order information</h1>
            {
                orders?.map((o, i) => {
                    return <table className={"table table-striped table-bordered"} key={i}>
                        <tbody>
                        <tr>
                            <th>Id</th>
                            <td>{o?.id}</td>
                        </tr>
                        <tr>
                            <th>Building number</th>
                            <td>{o?.buildingNum}</td>
                        </tr>
                        <tr>
                            <th>Dormitory number</th>
                            <td>{o?.dormNum}</td>
                        </tr>
                        <tr>
                            <th>Information</th>
                            <td>{o?.info}</td>
                        </tr>
                        <tr>
                            <th>Success</th>
                            <td>{o?.success}</td>
                        </tr>
                        </tbody>
                    </table>;
                })
            }
            <a href="#" onClick={_delete} className="my-button-big fw-bold">
                Delete Successful Order
            </a>
        </Layout>
    }
}
