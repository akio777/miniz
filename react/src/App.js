// eslint-disable-next-line
import logo from './logo.svg';
import './App.scss';
import React,{useState, useEffect} from 'react';
import {
  Switch,
  Route,
  Redirect,
  useLocation,
  useHistory,
  BrowserRouter as Router
} from "react-router-dom";
import axios from 'axios';
import config from './config.json';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './components/Header';


export default function App() {

  const history = useHistory()
  const [IsLogin, setIsLogin] = useState(!false)
  const [role, setRole] = useState({
    role : 2,
    name : "Guest"
  })

  const checkLogin=async()=>{
    let token = window.sessionStorage.getItem('token')
    if(token){
      let res = await axios.get(config.API_URL+'/check',{
        headers:{
          Authorization : 'Bearer ' + token
        }
      })
      .catch(err=>{
        console.log(err)
      })
      if(res){
        console.log(res)
      }
    }
  }

  const Empty=(props)=>{
    return (
      <div style={{
        fontSize:'2rem',
        display:'flex',
        justifyContent:'center',
        marginTop:'2rem'
      }}>
        <h1>{props.word}</h1>
      </div>
      )
  }

  return (
    <div>
      <Header />
      {/* <Router>
        <Navigationbar user={role} login={IsLogin}/>
        <Switch>
          <Route exact path="/Overview" render={()=><Empty word={"Overview"}/>} />
          <Route exact path="/Customerinfo" render={()=><Empty word={"Customerinfo"}/>} />
          <Route exact path="/Download" render={()=><Empty word={"Download"}/>} />
          <Route exact path="/Accountmanage" render={()=><Empty word={"Accountmanage"}/>} />
          <Route exact path="/" render={()=><Empty word={""}/>} />
        </Switch>
      </Router> */}
  </div>
  );
}

