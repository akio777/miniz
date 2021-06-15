
import { useState } from "react";

export default function useSearch(defaulValue=false){
    
    const [show, setShow] = useState(defaulValue)
    return [show, setShow]
}