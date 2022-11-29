import Link from "next/link";
import {useRouter} from "next/router";
import {useForm} from "react-hook-form"
import {Layout} from "../../components";
import {fetchWrapper} from "../../services/helpers/fetch_wrapper";

export default function Create() {
    const router = useRouter();
    const {register, handleSubmit} = useForm();

    function onSubmit(data) {
        fetchWrapper.post("/order/create", data)
            .then(resp => resp.json())
            .then(data => console.log(data));
        router.push("/order/info")
    }

    return <Layout>
        <h1>Create an Order</h1>
        <h2>Please enter a building number</h2>
        <div className="row">
            <div className="col-4">
                <form onSubmit={handleSubmit(onSubmit)}>
                    <div className="mb-3">
                        <label className="form-label">Building number</label>
                        <input type="text" className="form-control" {...register("buildingNum")}/>
                    </div>
                    <button type="submit" className="my-button">Submit</button>
                </form>
            </div>
        </div>
    </Layout>;
}
