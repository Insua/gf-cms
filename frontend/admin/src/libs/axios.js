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
        break
      default:
        break
    }
  }
  return Promise.reject(error)
})

export default Axios
