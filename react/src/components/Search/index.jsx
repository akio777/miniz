import React from "react";
import "./index.scss";
import { FiSearch } from "react-icons/fi";

import useSearch from "../../hooks/useSearch";

export default function Search(props) {

  const [show, setShow] = useSearch()

  return (
    <>
    
    <div className="SeachBox">
      <input placeholder="ค้นหา" />
      <div className="box">
        <FiSearch className="button" />
      </div>
    </div>

    {/* {show&&<div className="SeachBoxMobile">
      <input placeholder="ค้นหา" onBlur={()=>setShow(false)} autoFocus/>
      <div className="box">
        <FiSearch className="button" />
      </div>
    </div>}
    {!show&&<FiSearch className="single" onClick={()=>setShow(true)}/>} */}


    </>
    
  );
}
