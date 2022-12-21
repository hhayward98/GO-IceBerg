import { useState } from 'react';
import axios from 'axios';

const NavBar = () => {

	const RouteOne = () => {


		axios.post(`${"http://localhost:8080"}/RouteOne`, {Message: "Testing"}).then((response) => {

			console.log(response);

		});
	}

	return (

		<div id="NavBarDiv">
			<div className="row">
				<div className="col">
					
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

	);

}; export default NavBar;