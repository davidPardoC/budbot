import axios from "axios";
import Cookies from "js-cookie";
import { Redirect, Route, RouteProps } from "wouter";

export const ProtectedRoute = (props:RouteProps) => {
  const token = Cookies.get("token");

  if (!token) {
    return <Redirect to="/login" />;
  }

  axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;

  return <Route {...props} />;
};
