import Link from "next/link";
import {useRouter} from "next/router";
import {Layout} from "../../components";
import React from "react";
import {user} from "../../services/user";

export default function Register() {
    const router = useRouter();

    function submit(username: string, password: string) {
        user.register(username, password);
        router.push("/user/info");
    }

    return <Layout>
        <h1>Register</h1>
        <div className="row">
            <div className="col-4">
                <form
                    onSubmit={async function (e: React.SyntheticEvent) {
                        e.preventDefault();
                        const target = e.target as typeof e.target & {
                            username: { value: string };
                            password: { value: string };
                        };
                        await submit( target.username.value,target.password.value);
                    }}
                >
                    <div className="mb-3">
                        <label className="form-label">Username</label>
                        <input type="text" className="form-control" name="username"/>
                    </div>
                    <div className="mb-3">
                        <label className="form-label">Password</label>
                        <input type="password" className="form-control" name="password"/>
                    </div>
                    <button type="submit" className="my-button">Register</button>
                    &nbsp;&nbsp;&nbsp;&nbsp;
                    <Link href="/user/login" className="a-dec">Login</Link>
                </form>
            </div>
        </div>
    </Layout>;
}
