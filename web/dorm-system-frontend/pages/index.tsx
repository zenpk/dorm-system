import Link from "next/link"
import {Layout} from "../components"

export default function Home() {
    return <Layout>
            <div className="row g-3">
                <div className="col">
                    <div className={`font-art mm-1 fw-bold mfs-1 `}>Welcome to the dorm selecting system!</div>
                    <br/>
                    <Link href="/user/login" className="my-button-big mfs-5 mm-1 fw-bold">
                        Login
                    </Link>
                </div>
            </div>
        </Layout>;
}
