import Link from "next/link";
import {useRouter} from "next/router";
import {useForm} from "react-hook-form"
import {Layout} from "../../components";
import {fetchWrapper} from "../../services/helpers/fetch_wrapper";

export default function Register() {
    const router = useRouter();
    const {register, handleSubmit} = useForm();

    function onSubmit(data) {
        fetchWrapper.post("/register", data)
            .then(data => console.log(data));
        router.push("/user/info");
    }

    return <Layout>
        <h1>Register</h1>
        <div className="row">
            <div className="col-4">
                <form onSubmit={handleSubmit(onSubmit)}>
                    <div className="mb-3">
                        <label className="form-label">Username</label>
                        <input type="text" className="form-control" {...register("username")}/>
                    </div>
                    <div className="mb-3">
                        <label className="form-label">Password</label>
                        <input type="password" className="form-control" {...register("password")}/>
                    </div>
                    <button type="submit" className="my-button">Register</button>
                    &nbsp;&nbsp;&nbsp;&nbsp;
                    <Link href="/user/register" className="a-dec">Login</Link>
                </form>
            </div>
        </div>
    </Layout>;
}
