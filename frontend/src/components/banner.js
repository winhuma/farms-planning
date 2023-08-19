import "./banner.css";
import banner from "../assets/images/inspiration.png";
import { Carousel } from "antd";

const Banner = () => {
  return (
    <Carousel className="container-banner" autoplay>
      <div>
        <img src={banner} alt="banner" className="banner" />
      </div>
    </Carousel>
  );
};

export default Banner;
