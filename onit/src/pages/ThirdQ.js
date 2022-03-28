import axios from 'axios'
import React, { useEffect, useState } from 'react'
import ReactDOMServer from 'react-dom/server'
export default function ThirdQ() {
  const [json, setJson] = useState('hello')

  const [tumbler, setTumbler] = useState(false)
  const [marka, setMarka] = useState()
  const [model, setModel] = useState()
  const url = 'http://localhost:8080/v3/'
  const [jsonData, setJsonData] = useState()
  const [displayData, setDisplayData] = useState()
  const [globId, setGlobId] = useState()
  const [trTumbler, setTrTumbler] = useState(false)


  useEffect(function stringify() {
    setJson(JSON.stringify({ marka, model }))
  })

  useEffect(() => {
    if (tumbler == false) {
      getAuto()
    }
    // console.log(jsonData)
  })

  async function getAuto() {
    await axios({
      url: url + "GetAuto",
      method: "GET"
    }).then(resp => {
      setJsonData(resp.data)
    }).catch(e => {
      console.log(e)
    })
    getData()
    document.getElementById('btn_edit').style.visibility = 'hidden'
    setTumbler(true)
    
  }


  function editTr(_id) {
    document.getElementById('btn_edit').value = 'Edit Auto id '+_id
    document.getElementById('btn_edit').style.visibility = 'visible'
    setGlobId(_id)
  }

  // function editTr(_id) {
  //   const output = document.getElementById(_id).innerHTML
  //   console.log(output)
  //   Ee(_id, output)
  // }

  // function tumblerTr(_id) {
  //   console.log(trContent)  
  //   document.getElementById(_id).innerHTML = trContent
  // }

  // function Ee(_id, output) {
  //   document.getElementById(_id).innerHTML = "<td> _id </td><td> <input type='text' id='markaEdit' required /> </td>   <td> <input type='text' id='modelEdit' required /> </td>   <td> <input type='button' id='editAuto' value='Edit Auto' /> </td>   <td> <input type='button' onClick={tumblerTr(_id)} value='Back' /> </td>"
  //   console.log(document.getElementById('markaEdit').value)
  //   document.getElementById('editAuto').addEventListener("click", editAuto(_id))
  //   console.log('true')



  //   // return (
  //   //   <tr>
  //       // <td> {_id} </td>
  //       // <td> <input type='text' id='markaEdit' required /> </td>
  //       // <td> <input type='text' id='modelEdit' required /> </td>
  //       // <td> <input type='button' onClick={editAuto} value='Edit Auto' /> </td>
  //       // <td> <input type='button' onClick={tumblerTr(_id)} value='Back' /> </td>
  //   //   </tr>
  //   // )
  // }

  async function editAuto() {
    const _id = globId
      console.log(model, marka)
      await axios({
        url: url + "UpdateAuto/" + _id,
        method: "POST",
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        data: json
      }).catch(e => {
        console.log(e)
      })
    setTumbler(false)
  }

  async function deleteAuto(_id) {
    // if (delTumbler) {
    await axios({
      url: url + "DelAuto/" + _id,
      method: "GET"
    }).catch(e => {
      console.log(e)
    })

    setTumbler(false)
    // }
    // setDelTumbler(false)
  }
  

  async function addAuto() {
    await axios({
      url: url + "AddAuto",
      method: "POST",
      headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
      data: json
    }).catch(e => {
      console.log(e)
    })
    setTumbler(false)
  }

  // useEffect(function getGet(){
  //   getAuto()
  // })

  function getData() {
    if (jsonData !== undefined) {
      setDisplayData(jsonData.map(
        (info) => {
          return (
            <tr id={info.id} >
              <td id={info.id+'id'}>{info.id}</td>
              <td id={info.id+'marka'}>{info.marka}</td>
              <td id={info.id+'model'}>{info.model}</td>
              <td>
                <input type="button" value="Delete" onClick={e => deleteAuto(info.id)} />
              </td>
              <td>
                <input type="button" value="Edit" onClick={e => editTr(info.id)} />
              </td>
            </tr>
          )
        })
      )
    }
    else console.log('undef')
  }

  return (
    <>
      <table className="table table-hover">
        <thead>
          <tr>
            <th scope="col">ID</th>
            <th scope="col">Mark</th>
            <th scope="col">Model</th>
            <th scope='col'>Buttons</th>
            <th scope='col'></th>
          </tr>
        </thead>
        <tbody>
          {displayData}
        </tbody>
      </table>
      <input type='text' onChange={e => setMarka(e.target.value)} />
      <input type='text' onChange={e => setModel(e.target.value)} />
      <input type='button' id='btn_add' onClick={e=>addAuto()} value='Add Auto' /> 
      <input type='button' id='btn_edit' onClick={e=>editAuto()} />
    </>
  )
}
