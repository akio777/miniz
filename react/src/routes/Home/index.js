import Content from '../../components/Content';
import './index.scss';
import Slidepanel from '../../components/Slidepanel';

export function Home(props){
    return(
        <Content ShowFooter={!true}>
            {/* <Slidepanel name="New"/> */}
            {/* <Slidepanel name="Popular"/> */}
        </Content>
    )
}