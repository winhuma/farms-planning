import { Route, Routes, Link } from "react-router-dom";
import HomePage from "../views/homePage";
import FormWaterConsumption from "../views/formWaterConsumption";

const NoMatch = () => {
  return (
    <div style={{ textAlign: "center" }}>
      <h1 style={{ marginTop: "25%", marginBottom: "25%" }}>
        <Link to="/"> 404 - NOT FOUND! </Link>
      </h1>
    </div>
  );
};

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/watercal" element={<FormWaterConsumption />} />
      <Route path="*" element={<NoMatch />} />
    </Routes>
  );
};

export default AppRoutes;
