import { useState } from 'react';
import axios from 'axios';
import './App.css';
import NavBar from "./components/NavBar.jsx";


function App() {

  const [ BTN1, setBTN1 ] = useState(false);
  const [ BTN2, setBTN2 ] = useState(false);
  const [ BTN3, setBTN3 ] = useState(false);
  const [ BTN4, setBTN4 ] = useState(false);

  const [ Msg, setMsg ] = useState(""); 

  const RouteOne = () => {


    axios.post(`${"http://localhost:8080"}/RouteOne`, {Message: "Testing"}).then((response) => {

      console.log(response);
      setMsg(response.data.Message);

    });
  }
  



  return (
    <div className="App">
      <header className="App-header">
        <div id="NavBarDiv">
          <div className="row">
            <div className="col">
              <p>Col 1 </p>
              <button onClick={RouteOne}>Btn1</button>
            </div>
            <div className="col">
              <p>Col 2</p>
            </div>
            <div className="col">
              <p>Col 3</p>
            </div>
            <div className="col">
              <p>Col 4</p>
            </div>
          </div>
        </div>
      </header>
      <br/>
      <br/>
      <h1>{Msg}</h1>
      {BTN1 ? <h1>Btn1</h1> : null }
      {BTN2 ? <h1>Btn2</h1> : null }
      {BTN3 ? <h1>Btn3</h1> : null }
      {BTN4 ? <h1>Btn4</h1> : null }
    </div>
  );
}

export default App;
