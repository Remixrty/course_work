import React from 'react'
import { NavLink } from 'react-router-dom'

export default function Navbar() {
    return (
            <nav className="navbar navbar-expand-lg navbar-light bg-light">
                <div className="container-fluid">
                   <NavLink
                   className="nav-link"
                   to="/">
                       ONIT
                   </NavLink>
                   
                    <div className="collapse navbar-collapse" id="navbarSupportedContent">
                        <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                            <li className="nav-item">
                                <NavLink
                                    className="nav-link"
                                    to="/">
                                    Home
                                </NavLink>
                            </li>
                            <li className="nav-item">
                                <NavLink
                                className="nav-link"
                                to='/first'>
                                    FirstQ
                                </NavLink>
                            </li>
                            <li className="nav-item">
                                <NavLink
                                className="nav-link"
                                to='/second'>
                                    SecondQ
                                </NavLink>
                            </li>
                            <li className="nav-item">
                                <NavLink
                                className="nav-link"
                                to='/third'>
                                    ThirdQ
                                </NavLink>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
    )
}
