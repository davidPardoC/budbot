import Cookies from "js-cookie";
import { Redirect, Route, RouteProps } from "wouter";

export const ProtectedRoute = (props:RouteProps) => {
  const isAuth = Cookies.get("token");

  if (!isAuth) {
    return <Redirect to="/login" />;
  }

  return <Route {...props} />;
};
