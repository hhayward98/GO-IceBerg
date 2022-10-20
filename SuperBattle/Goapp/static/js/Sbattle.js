window.addEventListener("DOMContentLoaded", domLoaded);


function domLoaded() {

	const FightBtn = document.getElementById("fight");

	FightBtn.addEventListener("click", StartFight);

};


function StartFight() {
	console.log("Fight started")
	const selectedHeroValue = document.querySelector("#heroSelect").value;
	const selectedVillainValue = document.querySelector("#villainSelect").value;
	const Winner = document.getElementById("WinData");

	const randnum = Math.floor((Math.random() * 10) + 1);

	console.log(randnum);


	if (randnum > 5) {
		console.log("villain Wins");
		Winner.innerHTML = "<p> " + selectedVillainValue + "!";

	} else if (randnum < 5) {
		console.log("hero Wins");
		Winner.innerHTML = "<p> " + selectedHeroValue + "!";
	}
	console.log(selectedVillainValue);
	console.log(selectedHeroValue);




}
