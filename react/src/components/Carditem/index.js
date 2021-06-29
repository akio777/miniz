import { useEffect } from "react";
import "./index.scss";

export default function Carditem(props) {
  return (
    <div className="carditem" onClick={()=>{console.log('eeeeee')}}>
      <img src={props.src}/>
    </div>
  );
}
