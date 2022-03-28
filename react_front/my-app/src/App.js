import React from 'react'
import {api} from './api'



const App = () => {

  const result = () => api.get(`/lol/gg`).then(response => console.log(response.data))
  result()
    return (
        <div>
          good job
        </div>
    )};

export default App;