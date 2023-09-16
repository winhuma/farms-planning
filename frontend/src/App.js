import "./App.css";
import "antd/dist/reset.css";
import React from "react";
import AppRoutes from "./routes/appRoutes";
import { Link } from "react-router-dom";
import logo from "./assets/images/logo.png";
import { Layout, Space, Divider, Menu } from "antd";

const { Header, Content, Footer } = Layout;

const menuItems = [
  {
    key: "0",
    label: <Link to="/"> หน้าแรก </Link>,
  },
  {
    key: "1",
    label: "บริการของเรา",
    children: [
      {
        key: "11",
        label: <Link to="/watercal"> คำนวณความต้องการน้ำ </Link>,
      },
    ],
  },
];

const App = () => {
  return (
    <div>
      <Space className="app-space" direction="vertical" size={[0, 48]}>
        <Layout className="app-layout">
          <Header className="app-header">
            <div className="navbar">
              <span className="logo">
                <img src={logo} alt="logo" /> FARM PLAN
                <Divider type="vertical" />
              </span>

              <Menu
                className="menu"
                mode="horizontal"
                defaultSelectedKeys={["0"]}
                items={menuItems}
              />
            </div>
          </Header>

          <Content className="app-content">
            <AppRoutes />
          </Content>

          <Footer className="app-footer">
            <span>Copyright ©2023 Farm plan. All rights reserved.</span>
          </Footer>
        </Layout>
      </Space>
    </div>
  );
};

export default App;
