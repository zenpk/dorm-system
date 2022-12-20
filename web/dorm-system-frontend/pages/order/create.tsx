import {useRouter} from "next/router";
import {Layout} from "../../components";
import {fetchWrapper} from "../../services/fetch_wrapper";
import React from "react";

export default function Create() {
    const router = useRouter();

    async function submit(data: {
        buildingNum: string
    }) {
        let resp = await fetchWrapper.post("/order/submit", data);
        console.log(resp);
        await router.push("/order/info");
    }

    return <Layout>
        <h1>Create an Order</h1>
        <h2>Please enter a building number</h2>
        <div className="row">
            <div className="col-4">
                <form
                    onSubmit={async function (e: React.SyntheticEvent) {
                        e.preventDefault();
                        const target = e.target as typeof e.target & {
                            buildingNum: { value: string };
                        };
                        const buildingNum = target.buildingNum.value;
                        await submit({buildingNum});
                    }}
                >
                    <div className="mb-3">
                        <label className="form-label">Building number</label>
                        <input type="text" className="form-control" name="buildingNum"/>
                    </div>
                    <button type="submit" className="my-button">Submit</button>
                </form>
            </div>
        </div>
    </Layout>;
}
