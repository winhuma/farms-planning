import "./formWaterConsumption.css";
import React, { useState } from "react";
import axios from "axios";
import Topic from "../components/topic";
import { Form, Input, Radio, InputNumber, Button, Card, Divider } from "antd";

const apiEndpoint = process.env.REACT_APP_API_ENDPOINT;
const postUrlCalculate = apiEndpoint + "waters/calculate/day";
const tabList = [
  {
    key: "water-1",
    tab: "อุปโภค - บริโภค",
  },
  {
    key: "water-2",
    tab: "เพาะปลูก",
  },
  {
    key: "water-3",
    tab: "ปศุสัตว์",
  },
  {
    key: "water-4",
    tab: "อุตสาหกรรม",
  },
];

const FormWaterConsumption = () => {
  const [area, setArea] = useState(null);
  const [numPeople, setNumPeople] = useState(null);
  const [waterDemand, setWaterDemand] = useState(null);
  const [result, setResult] = useState(null);

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
      .post(postUrlCalculate, {
        area_name: area,
        number_person: Number(numPeople),
        number_day: Number(waterDemand),
      })
      .then((response) => {
        setResult(response.data.result);
      });
  };

  return (
    <div className="water-consumption-layout">
      <Topic
        className="water-consumption-topic"
        topicName={"คำนวณความต้องการน้ำ"}
      />

      <Card
        tabList={tabList}
        bordered={true}
        headStyle={{}}
        bodyStyle={{
          textAlign: "start",
          boxShadow: "2px 2px 2px 2px lightgray",
        }}
      >
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
            <InputNumber className="water-consumption-input-number" /> คน
          </Form.Item>

          <Form.Item label="ความต้องการน้ำ" onChange={onChangeWaterDemand}>
            <Input className="water-consumption-input-text" /> วัน
          </Form.Item>

          <Divider />

          <Form.Item label="ความต้องการน้ำทั้งหมด">
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
        </Form>
      </Card>
    </div>
  );
};

export default FormWaterConsumption;
