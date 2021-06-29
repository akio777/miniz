import './index.scss';

export default function Multiplecard(props){
    return(
        <div className="multiplecard">
            {props.children}
        </div>
    )
}