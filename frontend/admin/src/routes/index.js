import {lazy} from 'react'

const routes = [
  { path: '', exact: true, name: 'home', component: lazy(() => import('../views/dashboard'))},
  { path: 'dashboard', exact: true, name: 'dashboard', component: lazy(() => import('../views/dashboard'))}
]

export default routes
