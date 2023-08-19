import "./homePage.css";
import Banner from "../components/banner";
import Topic from "../components/topic";
import CardContent from "../components/cardContent";
import { Row, Col } from "antd";
import GoodPoint from "../components/goodpoint";

const HomePage = () => {
  const topic1 = "ข้อจำกัด";
  // const topic2 = 'ข้อดี'
  // const topic3 = 'บริการของเรา'
  // const topic4 = 'คลังความรู้'
  // const topic5 = 'ติดต่อเรา'
  return (
    <div>
      <Banner />

      <div className="restriction" id="restriction">
        <Topic topicName={topic1} />
        <p className="topic-p">
          เว็บไซต์นี้อาจจะไม่ได้เหมาะกับเจ้าของธุรกิจการเกษตรทุกคน
        </p>

        <Row className="row-card" gutter={[16, 16]}>
          <Col span={8} className="card1">
            <CardContent
              number="01"
              text="เราไม่ได้ลงไปสำรวจแปลงด้วยตนเอง จึงอาจมีความคลาดคลื่อนในผลการออกแบบ"
            />
          </Col>

          <Col span={8}>
            <CardContent
              number="02"
              text="ข้อมูลสภาพอากาศยังไม่ครอบคลุมทุกพื้นที่"
            />
          </Col>

          <Col span={8}>
            <CardContent
              number="03"
              text="ฐานข้อมูลพืช ยังครอบคลุมพืชเศรษฐกิจที่ปลูกจริง"
            />
          </Col>

          <Col span={8}>
            <CardContent
              number="04"
              text="ถ้าซื้ออุปกรณ์น้ำตามร้านขายปลีก อาจจะยังหายากในหลายพื้นที่"
            />
          </Col>

          <Col span={8}>
            <CardContent
              number="05"
              text="เราออกแบบให้ฟรี แต่เจ้าของธุรกิจจะต้องสั่งซื้ออุปกรณ์ผ่านเว็บไซต์เราเท่านั้น"
            />
          </Col>

          <Col span={8}>
            <CardContent
              number="06"
              text="ขอสงวนสิทธิ์ในการรับผิดชอบต่อความเสียหายที่เกิดขึ้น ที่อาจเกิดจากการใช้ผลการออกแบบของเรา"
            />
          </Col>
        </Row>
      </div>

      <GoodPoint />
    </div>
  );
};

export default HomePage;
