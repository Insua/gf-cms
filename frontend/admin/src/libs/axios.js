import React from 'react'
import { notification } from 'antd'
import collect from 'collect.js'
import Axios from 'axios'
import history from '_l/history'

Axios.interceptors.response.use(res => {
  return res
}, error => {
  let response = error.response
  if (response) {
    switch (response.status) {
      case 401:
        history.push(`${process.env.PUBLIC_URL}/login`)
        break
      case 422:
        const description = <div>
          {
            collect(response.data.errors).flatten().map((value, index) => {
              return <p key={index+value}>{`${value}`}</p>
            })
          }
        </div>
        notification['error']({
          message: '发生错误',
          description
        })
        break
      default:
        break
    }
  }
  return Promise.reject(error)
})

export default Axios
