import "./index.scss";

import { FiMenu } from "react-icons/fi";
import Search from "../Search";
import { Navbar, Container, Nav } from "react-bootstrap";
import { NavLink } from "react-router-dom";

export default function Navigationbar() {
  return (
    <Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
      <Container className="testX">
        <NavLink exact to="/" className="navbar-brand">
          LOGO APP
        </NavLink>
        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
        <Navbar.Collapse id="responsive-navbar-nav">
          <Nav className="me-auto">
            <NavLink exact to="/page1" className="navlink">
              page1
            </NavLink>
            <NavLink exact to="/page2" className="navlink">
              page2
            </NavLink>
            <NavLink exact to="/page3" className="navlink">
              page3
            </NavLink>
            <NavLink exact to="/page4" className="navlink">
              page4
            </NavLink>
            <NavLink exact to="/" className="navlink in-logout">
              logout
            </NavLink>
          </Nav>
        </Navbar.Collapse>
        <NavLink exact to="/" className="navlink  right-logout">
          logout
        </NavLink>
      </Container>
    </Navbar>
  );
}
