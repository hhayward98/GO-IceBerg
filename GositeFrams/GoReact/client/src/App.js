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

  const [ IsValid, setIsValid ] = useState(false);

  const RouteOne = () => {
    setIsValid(false);
    
    console.log("Running Route One....");
    axios.post(`${"http://localhost:8080"}/RouteOne`, {Message: "Testing"}).then((response) => {

      console.log("RouteOne: ",response);
      setMsg(response.data.Message);

    });
  }

  const RouteTwo = () => {

    
    console.log("Running Route Two....");


    axios.post(`${"http://localhost:8080"}/RouteTwo`, {Message: "Testing"}).then((response) => {

      console.log("RouteTwo: ",response);
      setMsg(response.data.Message);

      if (response.data.Auth == "true") {
        setIsValid(true);
      }

    });

  }

  const RouteThree = () => {
    setIsValid(false);
    
    console.log("Running Route Three....");

    axios.post(`${"http://localhost:8080"}/RouteThree`, {Message: "Testing"}).then((response) => {

      console.log("RouteOne: ",response);
      setMsg(response.data.Message);

    });

  }

  const RouteFour = () => {
    setIsValid(false);
    
    console.log("Running Route Four....");

    setMsg("");
  }


  const RouteP = () => {
    setIsValid(false);
    console.log("Running Route POST....");

    axios.post(`${"http://localhost:8080"}/RouteP`, {Message: "Testing"}).then((response) => {

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
              <button onClick={RouteTwo}>Btn1</button>
            </div>
            <div className="col">
              <p>Col 3</p>
              <button onClick={RouteThree}>Btn1</button>
            </div>
            <div className="col">
              <p>Col 4</p>
              <button onClick={RouteFour}>Btn1</button>
            </div>
          </div>
        </div>
      </header>
      <br/>
      <br/>
      {IsValid ? <button onClick={RouteP}>Post Request</button> : <h1>{Msg}</h1>}
      {BTN1 ? <h1>Btn1</h1> : null }
      {BTN2 ? <h1>Btn2</h1> : null }
      {BTN3 ? <h1>Btn3</h1> : null }
      {BTN4 ? <h1>Btn4</h1> : null }
    </div>
  );
}

export default App;
