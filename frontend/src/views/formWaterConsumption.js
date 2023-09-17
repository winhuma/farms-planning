import "./formWaterConsumption.css";
import React, { useState } from "react";
import axios from "axios";
import Topic from "../components/topic";
import {
  Form,
  Input,
  Radio,
  InputNumber,
  Button,
  Card,
  Divider,
  Select,
} from "antd";

const apiEndpoint = process.env.REACT_APP_API_ENDPOINT;
const postCalWaterDay = apiEndpoint + "waters/calculate/day";
const getProvince = apiEndpoint + "province";

const FormWaterConsumption = () => {
  const [area, setArea] = useState(null);
  const [numPeople, setNumPeople] = useState(null);
  const [waterDemand, setWaterDemand] = useState(null);
  const [result, setResult] = useState(null);
  const [activeTabKey, setActiveTabKey] = useState("day");
  const [selectedProvince, setSelectedProvince] = useState(null);
  const [searchProvince, setSearchProvince] = useState(null);
  const [province, setProvince] = useState(null);

  const onChangeProvince = (value) => {
    console.log("=selectedProvince=", selectedProvince);
    setSelectedProvince(value);
  };

  const onSearchProvince = (value) => {
    console.log("-searchProvince-", searchProvince);
    setSearchProvince(value);
  };

  const filterOptionProvince = (input, option) =>
    (option?.label ?? "").toLowerCase().includes(input.toLowerCase());

  const onTabChange = async (key) => {
    setActiveTabKey(key);
    if (key === "agriculture") {
      await axios.get(getProvince).then((response) => {
        let resp = [];
        for (const res of response.data.result) {
          resp = [
            ...resp,
            {
              value: res["id"],
              label: res["province_name"],
            },
          ];
        }
        setProvince(resp);
      });
    }
  };

  const onChangeRadioArea = (e) => {
    setArea(e.target.value);
  };

  const onChangeNumPeople = (e) => {
    setNumPeople(e.target.value);
  };

  const onChangeWaterDemand = (e) => {
    setWaterDemand(e.target.value);
  };

  const onClickCalculate = async () => {
    await axios
      .post(postCalWaterDay, {
        area_name: area,
        number_person: Number(numPeople),
        number_day: Number(waterDemand),
      })
      .then((response) => {
        setResult(response.data.result);
      });
  };

  //---------------------------------------------------------------------

  const cardTitle = (
    <Topic
      className="water-consumption-topic"
      topicName={"คำนวณความต้องการน้ำ"}
    />
  );

  const tabList = [
    {
      key: "day",
      tab: "อุปโภค - บริโภค",
    },
    {
      key: "agriculture",
      tab: "เพาะปลูก",
    },
    {
      key: "livestock",
      tab: "ปศุสัตว์",
    },
    {
      key: "industry",
      tab: "อุตสาหกรรม",
    },
  ];

  const formDay = (
    <Form
      className="water-consumption-form"
      labelCol={{ span: 3 }}
      wrapperCol={{ span: 24 }}
      layout="horizontal"
      labelAlign={"left"}
    >
      <Form.Item label="เขตพื้นที่">
        <Radio.Group onChange={onChangeRadioArea} value={area}>
          <Radio value="กรุงเทพฯ"> กรุงเทพฯ </Radio>
          <Radio value="องการบริหารส่วนตำบล"> องการบริหารส่วนตำบล </Radio>
          <Radio value="เทศบาลตำบล"> เทศบาลตำบล </Radio>
          <Radio value="เทศบาลนคร"> เทศบาลนคร </Radio>
          <Radio value="เมืองพัทยา"> เมืองพัทยา </Radio>
          <Radio value="เทศบาลเมือง"> เทศบาลเมือง </Radio>
        </Radio.Group>
      </Form.Item>

      <Form.Item label="จำนวนผู้ใช้น้ำ" onChange={onChangeNumPeople}>
        <InputNumber
          className="water-consumption-input-number"
          value={numPeople}
        />{" "}
        คน
      </Form.Item>

      <Form.Item label="ความต้องการน้ำ" onChange={onChangeWaterDemand}>
        <Input className="water-consumption-input-text" value={waterDemand} />{" "}
        วัน
      </Form.Item>
    </Form>
  );

  // const province = [
  //   {
  //     value: "1",
  //     label: "กรุงเทพฯ",
  //   },
  //   {
  //     value: "2",
  //     label: "สมุทรสาคร",
  //   },
  //   {
  //     value: "3",
  //     label: "นนทบุรี",
  //   },
  // ];

  const formAgr = (
    <Form
      className="water-consumption-form"
      labelCol={{ span: 3 }}
      wrapperCol={{ span: 24 }}
      layout="horizontal"
      labelAlign={"left"}
    >
      <Select
        className="water-consumption-dropdown"
        showSearch
        placeholder="จังหวัด"
        optionFilterProp="children"
        // value={}
        onChange={onChangeProvince}
        onSearch={onSearchProvince}
        filterOption={filterOptionProvince}
        options={province}
      />

      <br />
      <br />
    </Form>
  );

  const formLiv = <p> Livestock </p>;

  const formInd = <p> Industry </p>;

  const contentList = {
    day: formDay,
    agriculture: formAgr,
    livestock: formLiv,
    industry: formInd,
  };

  //======================================================================

  return (
    <div className="water-consumption-layout">
      <Card
        title={cardTitle}
        tabList={tabList}
        activeTabKey={activeTabKey}
        onTabChange={onTabChange}
        bordered={true}
        headStyle={{}}
        bodyStyle={{
          textAlign: "start",
          boxShadow: "2px 2px 2px 2px lightgray",
        }}
      >
        {contentList[activeTabKey]}

        <Divider />

        <Form.Item
          label="ความต้องการน้ำทั้งหมด"
          className="water-consumption-form-result"
        >
          <Input
            className="water-consumption-input-result"
            disabled
            value={result}
          />{" "}
          ลูกบาศก์เมตร
          <Button
            className="water-consumption-button"
            type="primary"
            onClick={onClickCalculate}
          >
            คำนวณ
          </Button>
        </Form.Item>
      </Card>
    </div>
  );
};

export default FormWaterConsumption;
