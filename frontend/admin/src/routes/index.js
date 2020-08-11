import {lazy} from 'react'

const routes = [
  { path: '', exact: true, name: 'home', component: lazy(() => import('_v/dashboard'))},
  { path: 'dashboard', exact: true, name: 'dashboard', component: lazy(() => import('_v/dashboard'))}
]

export default routes
