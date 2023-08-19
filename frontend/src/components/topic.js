import "./topic.css";

const Topic = (props) => {
  return (
    <>
      <div className="block-size"></div>
      <h3 className="topic">{props.topicName}</h3>
    </>
  );
};

export default Topic;
