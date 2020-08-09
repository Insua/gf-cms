import React, { Component,Fragment } from 'react'
import { Link } from 'react-router-dom'
import { Menu } from 'antd'
import {DashboardOutlined, UnorderedListOutlined, SettingOutlined} from '@ant-design/icons'

class SideMenu extends Component {
  render ()
  {
    return (
      <Fragment>
        <Menu  theme="dark" mode="inline">
          <Menu.Item key="1" icon={<DashboardOutlined/>}>
            <Link to={`${process.env.PUBLIC_URL}/`}>首页</Link>
          </Menu.Item>
          <Menu.Item key="2" icon={<UnorderedListOutlined/>}>
            <Link to={`${process.env.PUBLIC_URL}/category`}>文章分类</Link>
          </Menu.Item>
          <Menu.Item key="3" icon={<SettingOutlined/>}>
            <Link to={`${process.env.PUBLIC_URL}/setting`}>系统设置</Link>
          </Menu.Item>
        </Menu>
      </Fragment>
    )
  }
}

export default SideMenu
