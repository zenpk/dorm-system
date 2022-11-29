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
                        console.log(data);
                        setOrders(data.orders);
                    }
                }
            )
    }, []);

    const router = useRouter();

    function _delete(id) {
        let data = {};
        data.orderId = parseInt(id);
        fetchWrapper.delete("/order/delete", data)
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
            <h1>My order information</h1>
            {
                orders?.map((o, i) => {
                    return <div className={"mb-3"} key={i}>
                        <table className={"table table-striped table-bordered"}>
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
                                <td>{o?.success.toString()}</td>
                            </tr>
                            </tbody>
                        </table>
                        <a href="#" onClick={() => _delete(o?.id)} className="my-button-big fw-bold">
                            Delete
                        </a>
                    </div>;
                })
            }
        </Layout>
    }
}
