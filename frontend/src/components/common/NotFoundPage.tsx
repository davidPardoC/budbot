
import NotFoundImage from "../../assets/404.png";

export const NotFoundPage = () => {
  return (
    <div className="container flex items-center justify-center min-h-screen">
        <img src={NotFoundImage} alt="404"  className="w-80 h-80" />
    </div>
  )
}
