import "./App.css";
import "antd/dist/reset.css";
import React, { useState } from "react";
import logo from "./assets/images/logo.png";
import { Layout, Space, Divider, Menu } from "antd";
import HomePage from "./views/homePage";
import FormWaterConsumption from "./views/formWaterConsumption";

const { Header, Content, Footer } = Layout;

const menuItems = [
  {
    key: "0",
    label: "หน้าแรก",
  },
  {
    key: "1",
    label: "แรงบันดาลใจ",
  },
  {
    key: "2",
    label: "ข้อจำกัด",
  },
  {
    key: "3",
    label: "ข้อดี",
  },
  {
    key: "4",
    label: "บริการของเรา",
    children: [
      {
        key: "41",
        label: "คำนวณความต้องการน้ำ",
      },
      // {
      //   key: '42',
      //   label: 'item2',
      // },
      // {
      //   key: '43',
      //   label: 'item3',
      // },
    ],
  },
  {
    key: "5",
    label: "คลังความรู้",
  },
  {
    key: "6",
    label: "ติดต่อ/ข้อเสนอแนะ",
  },
];

const App = () => {
  const [pageStage, setPageStage] = useState("0");

  const onClickMenu = (e) => {
    setPageStage(e.key);
  };

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
                onClick={onClickMenu}
              />
            </div>
          </Header>

          <Content className="app-content">
            {/* แก้เป็น child ภายหลัง */}
            {pageStage === "41" ? <FormWaterConsumption /> : <HomePage />}
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
