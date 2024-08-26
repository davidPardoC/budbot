import { NotFoundPage } from "@/components/common/NotFoundPage"
import { ProtectedRoute } from "@/components/common/ProtectedRoute"
import { HomePage } from "@/pages/HomePage"
import { SettingsPage } from "@/pages/SettingsPage"
import { Redirect, Route, Router, Switch } from "wouter"
import LoginPage from "../pages/LoginPage"

const MainRouter = () => {
  return (
    <Switch>
        <Route path="/" component={() => <Redirect to="/login" /> } />
        <Route path="/login" component={LoginPage} />
        <ProtectedRoute path="/home" component={HomePage} />
        <Route path="/settings" component={SettingsPage} />
        <Router>
          <NotFoundPage/>
        </Router>
    </Switch>
  )
}

export default MainRouter