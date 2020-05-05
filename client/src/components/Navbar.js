import React, { Component } from "react";
import { MDBNavbar,
        MDBNavbarBrand,
        MDBNavbarNav,
        MDBNavItem,
        MDBNavLink,
        MDBNavbarToggler,
        MDBCollapse,
        MDBFormInline
    } from "mdbreact";
import { BrowserRouter as Router } from 'react-router-dom';

class Navbar extends Component {
    state = { isOpen: false };

    toggleCollapse = () => {
        this.setState({ isOpen: !this.state.isOpen });
    }

    render() {
        return (
            <Router>
                <MDBNavbar color="black" dark expand="md">
                    <MDBNavbarBrand>
                    <strong className="white-text">theAnalyst</strong>
                    </MDBNavbarBrand>
                    <MDBNavbarToggler onClick={this.toggleCollapse} />
                    <MDBCollapse id="navbarCollapse3" isOpen={this.state.isOpen} navbar>
                        <MDBNavbarNav left>
                            <MDBNavItem>
                                <MDBNavLink to="/">Home</MDBNavLink>
                            </MDBNavItem>
                            <MDBNavItem>
                                <MDBNavLink to="/">Graphs</MDBNavLink>
                            </MDBNavItem>
                        </MDBNavbarNav>
                        <MDBNavbarNav right>
                            <MDBNavItem>
                                <MDBFormInline waves>
                                    <div className="md-form my-0">
                                    <input className="form-control mr-sm-2" type="text" placeholder="Search" aria-label="Search" />
                                    </div>
                                </MDBFormInline>
                            </MDBNavItem>
                        </MDBNavbarNav>
                    </MDBCollapse>
                </MDBNavbar>
            </Router>
        );
    }
}

export default Navbar;