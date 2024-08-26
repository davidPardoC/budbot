import { useEffect, useRef } from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "../../ui/card";

const LoginForm = () => {

  const divref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const script = document.createElement("script");
    script.src = "https://telegram.org/js/telegram-widget.js?22";
    script.async = true;
    script.setAttribute("data-telegram-login", "PocketBudBot");
    script.setAttribute("data-size", "large");
    script.setAttribute(
      "data-auth-url",
      "http://127.0.0.1/api/v1/auth/telegram/callback"
    );
    script.setAttribute("data-request-access", "write");

    if (divref.current) {
      divref.current.appendChild(script);
    }
  })


  return (
    <Card className="max-w-fit">
      <CardHeader>
        <CardTitle className="text-center">
          <p className="text-4xl">🤖💸</p>
         <h1 className="mt-3">Wellcome to BudBot</h1> 
        </CardTitle>
        <CardDescription className="text-center">Please signin to see your dashboard </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="flex justify-center" ref={divref}></div>
      </CardContent>
      <CardFooter>
        <a href="https://t.me/PocketBudBot" className="text-center underline text-blue-900">Dont have an account? Click here.</a>
      </CardFooter>
    </Card>
  );
};


export default LoginForm;
