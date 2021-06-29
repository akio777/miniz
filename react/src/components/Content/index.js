import Footer from "../Footer";
import "./index.scss";

export default function Content({children, ShowFooter}) {
  return (
    <>
      <div className="content">
        <div className="inside">
            {children}
        </div>
      </div>
      {ShowFooter&&<Footer />}
    </>
  );
}
