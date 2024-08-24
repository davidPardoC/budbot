import { Redirect, Route, Switch } from "wouter"
import LoginPage from "../pages/LoginPage"

const MainRouter = () => {
  return (
    <Switch>
        <Route path="/" component={() => <Redirect to="/login" /> } />
        <Route path="/login" component={LoginPage} />
    </Switch>
  )
}

export default MainRouter