import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "../../ui/card";

const LoginForm = () => {
  return (
    <Card>
      <TelegramScript />
      <CardHeader>
        <CardTitle>Card Title</CardTitle>
        <CardDescription>Card Description</CardDescription>
      </CardHeader>
      <CardContent>
        <p>Card Content</p>
      </CardContent>
      <CardFooter>
        <p>Card Footer</p>
      </CardFooter>
    </Card>
  );
};

const TelegramScript = () => {
  return (
    <>
      <script
        async
        src="https://telegram.org/js/telegram-widget.js?22"
        data-telegram-login="PocketBudBot"
        data-size="large"
        data-auth-url="https://budbot.suitsoftware.com/api/telegram/outh/callback"
        data-request-access="write"
      ></script>
    </>
  );
};

export default LoginForm;
