// eslint-disable-next-line
import logo from "./logo.svg";
import "./App.scss";
import React, { useState, useEffect } from "react";
import {
  Switch,
  Route,
  Redirect,
  useLocation,
  useHistory,
  BrowserRouter as Router,
} from "react-router-dom";
import axios from "axios";
import config from "./config.json";
import "bootstrap/dist/css/bootstrap.min.css";
import Navigationbar from "./components/Navigationbar";

import {Home} from './routes';
import Content from "./components/Content";

export default function App() {
  const history = useHistory();
  const [IsLogin, setIsLogin] = useState(!false);
  const [role, setRole] = useState({
    role: 2,
    name: "Guest",
  });

  const checkLogin = async () => {
    let token = window.sessionStorage.getItem("token");
    if (token) {
      let res = await axios
        .get(config.API_URL + "/check", {
          headers: {
            Authorization: "Bearer " + token,
          },
        })
        .catch((err) => {
          console.log(err);
        });
      if (res) {
        console.log(res);
      }
    }
  };

  const Empty = (props) => {
    return (
      <Content ShowFooter={false}>{props.word}</Content>
    );
  };

  return (
    <React.StrictMode>
    <div>
      <Router>
        <Navigationbar />
        {/* <Navigationbar user={role} login={IsLogin}/> */}
        <Switch>
          <Route exact path="/page1" render={() => <Empty word={"page1"} />} />
          <Route exact path="/page2" render={() => <Empty word={"page2"} />} />
          <Route exact path="/page3" render={() => <Empty word={"page3"} />} />
          <Route exact path="/page4" render={() => <Empty word={"page4"} />} />
          <Route exact path="/" render={() => <Home />} />
        </Switch>
      </Router>
    </div>
    </React.StrictMode>
  );
}
