import React, { Component, lazy, Suspense } from 'react'
import { Route, Switch } from 'react-router-dom'
import { Layout as Container } from 'antd'
import { MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons'
import axios from 'axios'
import SideMenu from './components/side_menu'
import routes from '../../routes'
import './index.scss'

const {Header, Sider, Content} = Container

class Layout extends Component {

  state = {
    collapsed: false
  }

  toggle = () => {
    this.setState({
      collapsed: !this.state.collapsed
    })
  }

  render ()
  {
    return (
      <Container className="out-container">
        <Sider trigger={null} collapsible collapsed={this.state.collapsed}>
          <div className="layout-logo">gf cms</div>
          <SideMenu/>
        </Sider>
        <Container className="layout-container">
          <Header className="layout-header-background" style={{ padding: 0 }}>
            {
              React.createElement(this.state.collapsed ? MenuUnfoldOutlined : MenuFoldOutlined, {
                className: 'trigger',
                onClick: this.toggle
              })
            }
          </Header>
          <Content className="site-layout-content" style={{
            margin: '24px 16px',
            padding: 24,
            minHeight: 280,
          }}>
            <Suspense fallback="">
              <Switch>
                {
                  routes.map((route, index) => {
                    return route.component && (
                      <Route key={index} path={`${process.env.PUBLIC_URL}/${route.path}`} exact={route.exact} name={route.name} component={route.component}/>
                    )
                  })
                }
                <Route component={lazy(() => import('_v/errors/404'))}/>
              </Switch>
            </Suspense>
          </Content>
        </Container>
      </Container>
    )
  }

  async componentDidMount () {
    const {data} = await axios.get('/api/admin/auth/user')
    console.log(data)
  }
}

export default Layout
