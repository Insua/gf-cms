import React from 'react'
import { Route } from 'react-router-dom'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

NProgress.configure({ showSpinner: false })

const NprogressRoute = props => {
  React.useState(NProgress.start())

  React.useEffect(() => {
    NProgress.done()
    return () => NProgress.start()
  })

  return (
    <Route {...props} />
  )
}

export default NprogressRoute
