import Link from "next/link";
import {useRouter} from "next/router";
import {useForm} from "react-hook-form"
import {Layout} from "../../components";
import {fetchWrapper} from "../../services/helpers/fetch_wrapper";

export default function Join() {
    const router = useRouter();
    const {register, handleSubmit} = useForm();

    function onSubmit(data) {
        fetchWrapper.post("/join", data)
        router.push("/team/")
    }

    function create() {
        fetchWrapper.post("/create")
        router.push("/team/")
    }

    return <Layout>
        <h1>Please enter a team code</h1>
        <div className="row">
            <div className="col-4">
                <form onSubmit={handleSubmit(onSubmit)}>
                    <div className="mb-3">
                        <label className="form-label">Code</label>
                        <input type="text" className="form-control" {...register("code")}/>
                    </div>
                    <button type="submit" className="my-button">Join</button>
                    &nbsp;&nbsp;&nbsp;&nbsp;
                    <a href="#" onClick={create} className="a-dec">Create New</a>
                </form>
            </div>
        </div>
    </Layout>;
}
