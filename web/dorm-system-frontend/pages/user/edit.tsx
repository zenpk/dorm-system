import {useRouter} from "next/router";
import {Layout} from "../../components";
import {fetchWrapper} from "../../services/fetch_wrapper";
import React, {useState} from "react";
import {alertWrapper} from "../../services/alert_wrapper";
import {UserInfo} from "./info";


export default function Edit() {
    const router = useRouter();

    function submit(info: UserInfo) {
        type PostData = {
            user: UserInfo
        }
        const data: PostData = {
            user: info
        }
        fetchWrapper.put("/user/info", data)
            .then(data => {
                if (data.err.code !== process.env.errOK) {
                    alertWrapper.write(data.err.msg);
                }
            })
        router.push("/user/info")
    }

    const info = router.query;
    const [name, setName] = useState(info.name);
    const [studentNum, setStudentNum] = useState(info.studentNum);
    const [gender, setGender] = useState(info.gender);

    return <Layout>
        <h1>Edit</h1>
        <div className="row">
            <div className="col-4">
                <form
                    onSubmit={async function (e: React.SyntheticEvent) {
                        e.preventDefault();
                        const target = e.target as typeof e.target & {
                            name: { value: string };
                            studentNum: { value: string };
                            gender: { value: string };
                        };
                        const data: UserInfo = {
                            id: parseInt(info.id as string, 10),
                            name: target.name.value,
                            studentNum: target.studentNum.value,
                            gender: target.gender.value
                        }
                        await submit(data);
                    }}
                >
                    <div className="mb-3">
                        <label className="form-label">Name</label>
                        <input type="text" className="form-control" defaultValue={name}
                               onChange={evt => setName(evt.target.value)} name="name"/>
                    </div>
                    <div className="mb-3">
                        <label className="form-label">Student number</label>
                        <input type="text" className="form-control" defaultValue={studentNum}
                               onChange={evt => setStudentNum(evt.target.value)} name="studentNum"/>
                    </div>
                    <div className="mb-3">
                        <label className="form-label">Gender</label>
                        <input type="text" className="form-control" defaultValue={gender}
                               onChange={evt => setGender(evt.target.value)} name="gender"/>
                    </div>
                    <button type="submit" className="my-button">Confirm</button>
                </form>
            </div>
        </div>
    </Layout>;
}
