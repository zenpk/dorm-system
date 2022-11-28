import Link from "next/link"

export function Sidebar() {
    return <>
        <div className="sidebar-all" id="sidebar-all">
            <div className="sidebar-title">Menu</div>
            <ul className="list-unstyled">
                <li>
                    <Link href="/" className="link-dark sidebar-item">
                        Homepage
                    </Link>
                </li>
                <li>
                    <Link href="/user/info" className="link-dark sidebar-item">
                        Info
                    </Link>
                </li>
                <li>
                    <Link href="/dorm/info" className="link-dark sidebar-item">
                        Dormitory
                    </Link>
                </li>
                <li>
                    <button
                        className="btn btn-toggle align-items-center collapsed sidebar-button"
                        data-bs-toggle="collapse"
                        data-bs-target="#menu-team-collapse"
                        aria-expanded="false"
                    >
                        Team
                    </button>
                </li>
                <div className="collapse" id="menu-team-collapse">
                    <ul className="list-unstyled">
                        <li>
                            <Link href="/team/info" className="sidebar-subitem">
                                - My Team
                            </Link>
                        </li>
                        <li>
                            <Link href="/team/join" className="sidebar-subitem">
                                - Join a Team
                            </Link>
                        </li>
                    </ul>
                </div>
                <li>
                    <button
                        className="btn btn-toggle align-items-center collapsed sidebar-button"
                        data-bs-toggle="collapse"
                        data-bs-target="#menu-order-collapse"
                        aria-expanded="false"
                    >
                        Order
                    </button>
                </li>
                <div className="collapse" id="menu-order-collapse">
                    <ul className="list-unstyled">
                        <li>
                            <Link href="/order/info" className="sidebar-subitem">
                                - My Order
                            </Link>
                        </li>
                        <li>
                            <Link href="/order/create" className="sidebar-subitem">
                                - Create an Order
                            </Link>
                        </li>
                    </ul>
                </div>
            </ul>
            <div className="sidebar-line"></div>
        </div>
    </>;
}
