import { Route, Routes } from "react-router-dom";
import HomePage from "../views/homePage";
import FormWaterConsumption from "../views/formWaterConsumption";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/watercal" element={<FormWaterConsumption />} />
    </Routes>
  );
};

export default AppRoutes;
