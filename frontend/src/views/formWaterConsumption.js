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
const getPlant = apiEndpoint + "/waters?type=plant";
const postCalWaterDay = apiEndpoint + "/waters/calculate/day";
const postCalWaterAgri = apiEndpoint + "/waters/calculate/plant";

const FormWaterConsumption = () => {
  const [area, setArea] = useState(null);
  const [numPeople, setNumPeople] = useState(null);
  const [waterDemand, setWaterDemand] = useState(null);
  const [resultDay, setResultDay] = useState(null);
  const [resultAgri, setResultAgri] = useState(null);
  const [activeTabKey, setActiveTabKey] = useState("day");
  const [selectedProvince, setSelectedProvince] = useState(null);
  const [searchProvince, setSearchProvince] = useState(null);
  const [province, setProvince] = useState(null);
  const [selectedPlant, setSelectedPlant] = useState(null);
  const [searchPlant, setSearchPlant] = useState(null);
  const [plant, setPlant] = useState(null);
  const [resPlant, setResPlant] = useState(null);
  const [areaAgri, setAreaAgri] = useState(null);
  const [result, setResult] = useState(null);

  const onChangeProvince = (provId) => {
    setSelectedProvince(provId);
    for (let prov of resPlant) {
      if (prov["province_id"] === provId) {
        let plantList = [];
        for (var [key, value] of Object.entries(prov["data"])) {
          plantList = [
            ...plantList,
            {
              value: value,
              label: key,
            },
          ];
        }
        setPlant(plantList);
        break;
      }
    }
  };

  const onSearchProvince = (value) => {
    setSearchProvince(value);
  };

  const filterOptionProvince = (input, option) =>
    (option?.label ?? "").toLowerCase().includes(input.toLowerCase());

  const onChangePlant = (planVal) => {
    setSelectedPlant(planVal);
  };

  const onSearchPlant = (value) => {
    setSearchPlant(value);
  };

  const filterOptionPlant = (input, option) =>
    (option?.label ?? "").toLowerCase().includes(input.toLowerCase());

  const onTabChange = async (key) => {
    setActiveTabKey(key);
    if (key === "agriculture") {
      await axios.get(getPlant).then((response) => {
        let resp = [];
        for (let prov of response.data.result) {
          resp = [
            ...resp,
            {
              value: prov["province_id"],
              label: prov["province_name"],
            },
          ];
        }
        setProvince(resp);
        setResPlant(response.data.result);
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

  const onChangeAreaAgri = (e) => {
    setAreaAgri(e.target.value);
  };

  const onClickCalculate = async () => {
    let sum = 0

    await axios
      .post(postCalWaterDay, {
        area_id: area,
        number_person: Number(numPeople),
        number_day: Number(waterDemand),
      })
      .then((response) => {
        sum += response.data.result
        setResultDay(response.data.result);
      });

    await axios
      .post(postCalWaterAgri, {
        plant_value: Number(selectedPlant),
        farm_area: Number(areaAgri),
      })
      .then((response) => {
        sum += response.data.result
        setResultAgri(response.data.result);
      });

      setResult(sum)
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
          <Radio value={5}> กรุงเทพฯ </Radio>
          <Radio value={1}> องการบริหารส่วนตำบล </Radio>
          <Radio value={2}> เทศบาลตำบล </Radio>
          <Radio value={3}> เทศบาลนคร </Radio>
          <Radio value={4}> เมืองพัทยา </Radio>
          <Radio value={6}> เทศบาลเมือง </Radio>
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

  const formAgr = (
    <Form
      className="water-consumption-form"
      labelCol={{ span: 3 }}
      wrapperCol={{ span: 24 }}
      layout="horizontal"
      labelAlign={"left"}
    >
      <Form.Item label="เลือกจังหวัด">
        <Select
          className="water-consumption-dropdown"
          showSearch
          placeholder="จังหวัด"
          optionFilterProp="children"
          value={selectedProvince}
          onChange={onChangeProvince}
          onSearch={onSearchProvince}
          filterOption={filterOptionProvince}
          options={province}
        />
      </Form.Item>

      <Form.Item label="เลือกพันธุ์พืช">
        <Select
          className="water-agriculture-dropdown"
          showSearch
          placeholder="พันธุ์พืช"
          optionFilterProp="children"
          value={selectedPlant}
          onChange={onChangePlant}
          onSearch={onSearchPlant}
          filterOption={filterOptionPlant}
          options={plant}
        />
      </Form.Item>

      <Form.Item label="พื้นที่" onChange={onChangeAreaAgri}>
        <InputNumber
          className="water-agriculture-area-number"
          value={areaAgri}
        />{" "}
        ไร่
      </Form.Item>

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
