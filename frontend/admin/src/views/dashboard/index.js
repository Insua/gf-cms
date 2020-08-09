import React, { Component } from 'react'
import axios from 'axios'

class Dashboard extends Component {
  render ()
  {
    return (
      <div>
        dashboard
      </div>
    )
  }

  async componentDidMount () {
    const {data} = await axios.get('/api/user')
    console.log(data)
  }
}

export default Dashboard
