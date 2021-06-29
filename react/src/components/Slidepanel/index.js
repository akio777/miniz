import "./index.scss";
import { Button } from "react-bootstrap";
import Carditem from "../Carditem";
import { useEffect, useState } from "react";
import Multiplecard from "../Multiplecard";
export default function Slidepanel(props) {
  const [SHOW_PIC, setSHOW_PIC] = useState(4);
  
  return (
    <div className="slidepanel">
      <div className="name-button">
        <p>{props.name}</p>
        <Button variant="primary">see more</Button>{" "}
      </div>
      <div className="panel">{}</div>
    </div>
  );
}
