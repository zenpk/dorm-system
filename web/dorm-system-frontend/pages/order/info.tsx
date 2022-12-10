import {Layout} from "../../components"
import {useEffect, useState} from "react";
import {fetchWrapper} from "../../services/fetch_wrapper";
import {useRouter} from "next/router";
import {alertWrapper} from "../../services/alert_wrapper";

type OrderInfo = {
    id: number,
    buildingNum: string,
    dormNum: string,
    info: string,
    success: boolean,
    deleted: boolean,
}

export default function Info() {
    const [orders, setOrders] = useState<[OrderInfo]>();

    useEffect(() => {
        fetchWrapper.get("/order")
            .then(data => {
                    if (data.err.code !== process.env.errOK) {
                        alertWrapper.write(data.err.msg);
                    } else {
                        setOrders(data.orders);
                    }
                }
            )
    }, []);

    const router = useRouter();

    function _delete(id: number) {
        fetchWrapper.delete(`/order/delete?orderId=${id}`)
            .then(data => console.log(data));
        router.reload();
    }

    return <Layout>
        <h1>My order information</h1>
        {
            orders?.map((o, i) => {
                return <div className="mb-3" key={i}>
                    <table className="table table-striped table-bordered">
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
