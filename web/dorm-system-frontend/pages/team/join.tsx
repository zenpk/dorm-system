import {useRouter} from "next/router";
import {Layout} from "../../components";
import {fetchWrapper} from "../../services/fetch_wrapper";
import React from "react";
import {alertWrapper} from "../../services/alert_wrapper";

export default function Join() {
    const router = useRouter();

    function submit(code: string) {
        type PostData = {
            code: string
        }
        const data: PostData = {
            code: code
        }
        fetchWrapper.post("/team/join", data)
            .then(data => {
                if (data.err.code !== process.env.errOK) {
                    alertWrapper.write(data.err.msg);
                }
            });
        router.push("/team/info")
    }

    function create() {
        fetchWrapper.post("/team/create", null)
            .then(data => {{
                if (data.err.code !== process.env.errOK) {
                    alertWrapper.write(data.err.msg);
                }
            }});
        router.push("/team/info")
    }

    return <Layout>
        <h1>Join a Team</h1>
        <h2>Please enter a team code</h2>
        <div className="row">
            <div className="col-4">
                <form
                    onSubmit={async function (e: React.SyntheticEvent) {
                        e.preventDefault();
                        const target = e.target as typeof e.target & {
                            code: { value: string };
                        };
                        await submit(target.code.value);
                    }}
                >
                    <div className="mb-3">
                        <label className="form-label">Code</label>
                        <input type="text" className="form-control" name="code"/>
                    </div>
                    <button type="submit" className="my-button">Join</button>
                    &nbsp;&nbsp;&nbsp;&nbsp;
                    <a href="#" onClick={create} className="a-dec">Create New</a>
                </form>
            </div>
        </div>
    </Layout>;
}
