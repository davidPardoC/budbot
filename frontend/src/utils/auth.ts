import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

type TokenPayload = {
  chat_id: number;
  exp: number;
  iat: number;
  photo_url: string;
  user_id: number;
};

export const getUserFromToken = ():TokenPayload => {
  const token = Cookies.get("token");
  const decoded = jwtDecode<TokenPayload>(token as string);

  return decoded;
};
