import {useRouter} from "next/router";
import {useForm} from "react-hook-form"
import {Layout} from "../../components";
import {fetchWrapper} from "../../services/helpers/fetch_wrapper";
import {useState} from "react";


export default function Edit() {
    const router = useRouter();
    const {register, handleSubmit} = useForm();

    async function onSubmit(data) {
        let postBody = {};
        postBody.user = data;
        postBody.user.id = parseInt(data.id);
        let resp = await fetchWrapper.put("/user/info", postBody);
        let body = resp.json();
        if (resp.status !== 200) {
            alert(body);
        } else {
            router.push("/user/info")
        }
    }

    const info = router.query;
    const [name, setName] = useState(info.name);
    const [studentNum, setStudentNum] = useState(info.studentNum);
    const [gender, setGender] = useState(info.gender);

    return <Layout>
        <h1>Edit</h1>
        <div className="row">
            <div className="col-4">
                <form onSubmit={handleSubmit(onSubmit)}>
                    <div className="mb-3">
                        <label className="form-label">Name</label>
                        <input type="text" className="form-control" defaultValue={name}
                               onChange={evt => setName(evt.target.value)} {...register("name")}/>
                    </div>
                    <div className="mb-3">
                        <label className="form-label">Student number</label>
                        <input type="text" className="form-control" defaultValue={studentNum}
                               onChange={evt => setStudentNum(evt.target.value)} {...register("studentNum")}/>
                    </div>
                    <div className="mb-3">
                        <label className="form-label">Gender</label>
                        <input type="text" className="form-control" defaultValue={gender}
                               onChange={evt => setGender(evt.target.value)} {...register("gender")}/>
                    </div>
                    <button type="submit" className="my-button">Confirm</button>
                </form>
            </div>
        </div>
    </Layout>;
}
