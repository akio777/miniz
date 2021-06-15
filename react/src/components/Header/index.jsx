import './index.scss'

import {FiMenu} from 'react-icons/fi';
import Search from '../Search';

export default function Header(){
    return(
        <div className="header">
            <div className="menu"><FiMenu size={24}/></div>
            <div className="logo">LOGOAPP</div>
            <Search />
            <div className="account">{"Account"}</div>
        </div>
    )
}