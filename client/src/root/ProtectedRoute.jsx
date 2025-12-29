import { Navigate, Outlet } from "react-router-dom";

export const ProtectedRoute = () => {
  // 1. Cek Token di LocalStorage
  // jika sudah login akan ada
  const token = localStorage.getItem("token");

  // 2. Logika Redirect
  if (!token) {
    return <Navigate to="/login" replace />;
  }

  // 3. Kalau ada token
  return <Outlet />;
};
