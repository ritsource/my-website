// Some jaavscript for Header Animation (on Scroll)
var prevScrollpos = window.pageYOffset;

window.onscroll = function() {
	var currentScrollPos = window.pageYOffset;

	if (prevScrollpos > currentScrollPos) {
		document.getElementById('navbar').style.top = '0';
	} else {
		// -68 because header height is -68
		document.getElementById('navbar').style.top = '-68px';
	}

	prevScrollpos = currentScrollPos;
};
