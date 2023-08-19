import "./navbar.css";
import logo from "../assets/images/logo.png";
import { Divider, Menu } from "antd";

const Navbar = (props) => {
  return (
    <div className="navbar">
      <span className="logo">
        <img src={logo} alt="logo" /> FARM PLAN
        <Divider type="vertical" />
      </span>
      <Menu
        className="menu"
        mode="horizontal"
        defaultSelectedKeys={["0"]}
        items={props.menuItems}
      />
    </div>
  );
};

export default Navbar;

// ย้ายไป app.js แล้ว ลบได้
