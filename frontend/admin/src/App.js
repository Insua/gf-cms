import React, { Component, lazy, Suspense } from 'react'
import { Router, Switch } from 'react-router-dom'
import NprogressRoute from '_c/nprogress-route'
import history from '_l/history'
import './assets/styles/index.scss'

const Login = lazy(() => import('_v/login'))
const Layout = lazy(() => import('_v/layout'))

class App extends Component {
  render ()
  {
    return (
      <Router history={history}>
        <Suspense fallback="">
          <Switch>
            <NprogressRoute {...{exact: true, path: `${process.env.PUBLIC_URL}/login`, component: Login}}/>
            <NprogressRoute {...{path: `${process.env.PUBLIC_URL}/`, component: Layout}}/>
          </Switch>
        </Suspense>
      </Router>
    )
  }
}

export default App
