import React, { Component, lazy, Suspense } from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import './assets/styles/index.scss'

const Login = lazy(() => import('_v/login'))
const Layout = lazy(() => import('_v/layout'))

class App extends Component {
  render ()
  {
    return (
      <Router>
        <Suspense fallback="">
          <Switch>
            <Route exact path={`${process.env.PUBLIC_URL}/login`} component={Login}/>
            <Route path={`${process.env.PUBLIC_URL}/`} component={Layout}/>
          </Switch>
        </Suspense>
      </Router>
    )
  }
}

export default App
