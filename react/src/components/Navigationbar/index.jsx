// eslint-disable-next-line
import React, { useState } from "react";
import "./index.scss";
import { Navbar, Nav, Button } from "react-bootstrap";
import { NavLink, useHistory } from "react-router-dom";
import { useEffect } from "react";

import { TiThMenu } from "react-icons/ti";

export default function Navigationbar(props) {
  const history = useHistory();

  const LogoutHandler = () => {
    history.replace("/");
  };

  const [menu, setMenu] = useState(false);
  const toggleMenu = () => [setMenu(!menu)];
  const closeOpen = () => {
    setMenu(false);
  };
  return (
    <nav>

    </nav>
    
  );
}
