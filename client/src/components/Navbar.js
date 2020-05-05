import React, { Component } from "react";
import './style/Navbar.css';
import { MDBNavbar,
        MDBNavbarBrand,
        MDBNavbarNav,
        MDBNavItem,
        MDBNavLink,
        MDBNavbarToggler,
        MDBCollapse,
        MDBFormInline,
        MDBInput
    } from "mdbreact";
import { BrowserRouter as Router } from 'react-router-dom';

class Navbar extends Component {

    constructor(props) {
        super(props);
        this.state = { isOpen: false, searchValue: "" };
        this.onSubmit = this.onSubmit.bind(this);
        this.onChange = this.onChange.bind(this);
        this.toggleCollapse = this.toggleCollapse.bind(this);
    }

    onSubmit(event) {
        this.props.searchByTopic(this.state.searchValue);
        this.setState({ searchValue: "" });
        event.preventDefault();
    }

    onChange(event) {
        this.setState({ searchValue: event.target.value });
    }

    toggleCollapse() {
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
                                <MDBFormInline waves onSubmit={this.onSubmit}>
                                    <MDBInput className="form-control mr-sm-2" type="text" hint="Search" value={this.state.searchValue} onChange={this.onChange}/>
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