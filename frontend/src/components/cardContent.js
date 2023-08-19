import "./cardContent.css";
import { Card } from "antd";

const CardContent = (props) => {
  return (
    <Card className="card-content">
      <h1 className="card-number">{props.number}</h1>
      <p>{props.text}</p>
    </Card>
  );
};

export default CardContent;
