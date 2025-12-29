import "./App.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { ProtectedRoute } from "./root/ProtectedRoute";
import { LoginPage } from "./pages/authPages/LoginPage";
import { SignupPage } from "./pages/authPages/SignupPage";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          // Route Public login dan register
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<SignupPage />} />
          // route Protected ( harus login )
          <Route element={<ProtectedRoute />}></Route>
          <Route path="*" />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
