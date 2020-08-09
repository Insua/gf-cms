import React, { Component, lazy, Suspense } from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'

const Login = lazy(() => import('./views/login'))
const Layout = lazy(() => import('./views/layout'))

class App extends Component {
  render ()
  {
    return (
      <Router>
        <Suspense fallback={<div>Loading...</div>}>
          <Switch>
            <Route exact path={`${process.env.PUBLIC_URL}/login`} component={Login}/>
            <Route path={`${process.env.PUBLIC_URL}/`} exact component={Layout}/>
          </Switch>
        </Suspense>
      </Router>
    )
  }
}

export default App
