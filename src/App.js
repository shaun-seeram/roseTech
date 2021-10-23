import './App.css';
import {useEffect, useState} from "react"

function App() {

  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    fetch("http://localhost:10000/articles").then((data) => {
      return data.json()
    }).then((jsonData) => {
      setTasks(jsonData)
    })
  }, [])
  
  // useEffect(() => {
    // const params = {
    //   headers: {
    //     "content-type": "application/json; charset=UTF-8"
    //   },
    //   body: JSON.stringify(data),
    //   method: "POST"
    // }

    // fetch("http://www.localhost:10000/article", params).then((data) => {
    //   return data.json()
    // }).then((jsonData) => {
    //   console.log(jsonData)
    // }).catch((err) => {
    //   console.log(err)
    // })
  // }, [])

  const modules = document.querySelectorAll(".moduleContainer")

  const nodes = document.querySelectorAll(".moduleNode")

  let newData;

  nodes.forEach((node) => {
    node.addEventListener("dragstart", (e) => {
      newData = tasks.filter((item) => {
        return item.orderID === e.target.classList[1]
      })
      console.log("hi")
      console.log("NewData:", newData)
    })
  })

  modules.forEach((module) => {
    module.addEventListener("drop", (e) => {

      // newData[0].driver = e.target.id

      // const params = {
      //   headers: {
      //     "content-type": "application/json; charset=UTF-8"
      //   },
      //   body: JSON.stringify(newData[0]),
      //   method: "POST"
      // }
  
      // fetch("http://www.localhost:10000/article", params).then((data) => {
      //   return data.json()
      // }).then((jsonData) => {
      //   setTasks(jsonData)
      // }).catch((err) => {
      //   console.log(err)
      // })

      // console.log("Drop", newData)
      // newData = null
    //   console.log(e)

    })

    module.addEventListener("dragenter", (e) => {
      e.preventDefault()
    })

    module.addEventListener("dragover", (e) => {
      e.preventDefault()
    })
  })


  return (
    <div className="App">
      <header><h1>Rose Rocket Tech Test</h1></header>
      <main>

        <div className="moduleContainer" id="Unassigned">
          <h2>Unassigned Tasks</h2>
          <div className="moduleNode">
            <ul>
              <li className="move"></li>
              <li className="desc">Description</li>
              <li className="rev">Revenue</li>
              <li className="cost">Cost</li>
              <li className="edit"></li>
            </ul>
            </div>

          {
              tasks.map((task) => {
                if (task.driver === "Unassigned") {
                  return (
                    <div className={`moduleNode ${task.orderID}`} key={task.orderID} draggable="true">
                      <ul>
                        <li className="move">M</li>
                        <li className="desc">{task.description}</li>
                        <li className="rev">{task.revenue}</li>
                        <li className="cost">{task.cost}</li>
                        <li className="edit">E</li>
                      </ul>
                    </div>
                  )
                }
              })
          }

        </div>

        <div className="moduleContainer" id="SteveWilliams">
          <h2>Steve Williams</h2>
          <div className="moduleNode">
            <ul>
              <li className="move"></li>
              <li className="desc">Description</li>
              <li className="rev">Revenue</li>
              <li className="cost">Cost</li>
              <li className="edit"></li>
            </ul>
            </div>

          {
              tasks.map((task) => {
                if (task.driver === "SteveWilliams") {
                  return (
                    <div className={`moduleNode ${task.orderID}`} key={task.orderID} draggable="true">
                      <ul>
                        <li className="move">M</li>
                        <li className="desc">{task.description}</li>
                        <li className="rev">{task.revenue}</li>
                        <li className="cost">{task.cost}</li>
                        <li className="edit">E</li>
                      </ul>
                    </div>
                  )
                }
              })
          }

        </div>

        <div className="moduleContainer" id="ChrisHorton">
          <h2>Chris Horton</h2>
          <div className="moduleNode">
            <ul>
              <li className="move"></li>
              <li className="desc">Description</li>
              <li className="rev">Revenue</li>
              <li className="cost">Cost</li>
              <li className="edit"></li>
            </ul>
            </div>

          {
              tasks.map((task) => {
                if (task.driver === "ChrisHorton") {
                  return (
                    <div className={`moduleNode ${task.orderID}`} key={task.orderID} draggable="true">
                      <ul>
                        <li className="move">M</li>
                        <li className="desc">{task.description}</li>
                        <li className="rev">{task.revenue}</li>
                        <li className="cost">{task.cost}</li>
                        <li className="edit">E</li>
                      </ul>
                    </div>
                  )
                }
              })
          }

        </div>

      </main>
      <footer><p>Made by Shaun</p></footer>
    </div>
  );
}

export default App;
