window.addEventListener("DOMContentLoaded", domLoaded);
var clicks = 0;
const ImgSrc = ["https://live.staticflickr.com/65535/52301739289_b48de2b134_k.jpg", "https://live.staticflickr.com/65535/52300486877_28f72d32b3_k.jpg", "https://live.staticflickr.com/65535/52301467371_7e2c88eb62_k.jpg", "	https://live.staticflickr.com/65535/52301467421_c8add5a99c_k.jpg", "https://live.staticflickr.com/65535/52301480763_f0c8d80b4d_k.jpg" ];

function domLoaded() {
	let RTB = document.getElementById('RotateButton');
	let IMG = document.getElementById('RImg')

	RTB.addEventListener("click", function() {
		OnClick()
		IMG.src = ImgSrc[clicks]
	});
}

function OnClick() {
	clicks += 1;
	if (clicks >= ImgSrc.length) {
		clicks = 0;
	}
	
}
