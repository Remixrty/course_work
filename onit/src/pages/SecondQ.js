import axios from 'axios'
import React, { useState } from 'react'

export default function SecondQ() {
  // const [tudcp, setTudcp] = useState('TCP')
  
  // const [json, setJson] = useState()

  async function getPorts(boo) {
    let tudcp = ''
    console.log(boo)
    if (boo){
      tudcp = 'TCP'
    }
    else tudcp = 'UDP'
    const url = 'http://localhost:8080/v1/port'+tudcp
    const output = document.getElementById('newData')
    console.log(tudcp)
    // console.log('start')
    output.textContent = 'Пожалуйста, подождите'
    await axios ({
      method: "GET",
      url: url,
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
    }).then(resp => {
      const json = (resp.data)
      // console.log(json)
      output.textContent = JSON.stringify(json)
    })
    // separateJson()
  }

  // useEffect(() => {
  //   console.log(tudcp)
  // }, [tudcp])
  

  return (
    <>
    <input type="button" onClick={e=>getPorts(true)} value='TCP'/>
    <input type="button" onClick={e=>getPorts(false)} value='UDP'/>
    {/* <div>{json}</div> */}
    <p id='newData'></p>
    </>
  )
}
