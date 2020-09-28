import React, { Component } from 'react'
import { Button, Card, Form, Input, notification } from 'antd'
import axios from '_l/axios'
import history from '_l/history'
import './index.scss'

const bgImage = require('_a/images/kanna.png')

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
}
const tailLayout = {
  wrapperCol: { offset: 8, span: 16 },
}

const onFinishFailed = errorInfo => {
  notification.error({
    message: '验证未通过',
    description: '请按提示填写'
  })
}

class Login extends Component {
  constructor (props) {
    super(props)
    this.state = {
      loading: false
    }
  }

  finish = async values => {
    console.log(values)
    this.setState({
      loading: true
    })
    try {
      const { data } = await axios.post('/api/admin/auth/login', values)
      if (data.id) {
        notification.success({
          message: '登录成功',
          description: data.body.length > 0 ? `欢迎您${data.body}` : `欢迎您${data.name}`
        })
        history.replace(`${process.env.PUBLIC_URL}/`)
      }
    } finally {
      this.setState({
        loading: false
      })
    }
  }

  render ()
  {
    return (
      <div style={{ backgroundImage: 'url(' + bgImage + ')' }} className="login">
        <div className="login-con">
          <Card title="欢迎登录" style={{ width: 320 }}>
            <Form
              {...layout}
              initialValues={{ remember: true }}
              onFinish={values => this.finish(values)}
              onFinishFailed={onFinishFailed}
              style={{ backgroundColor: 'white' }}
            >
              <Form.Item
                label="用户名"
                name="name"
                rules={[{ required: true, message: '请输入用户名' }]}
              >
                <Input/>
              </Form.Item>

              <Form.Item
                label="密码"
                name="password"
                rules={[{ required: true, message: '请输入密码' }]}
              >
                <Input.Password/>
              </Form.Item>

              <Form.Item {...tailLayout}>
                <Button type="primary" htmlType="submit" block loading={this.state.loading}>登录</Button>
              </Form.Item>
            </Form>
          </Card>
        </div>
      </div>
    )
  }
}

export default Login
