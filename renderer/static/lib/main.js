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

// For Document Navigation in Footer
function setNavigation(data, urlBase) {
	let linkSeg = window.location.href.split('/');
	let currentId = linkSeg.pop();
	if (currentId === '') {
		currentId = linkSeg.pop();
	}

	cIndex = data.findIndex(({ _id }) => _id === currentId);

	const prevBtn = document.getElementById('Footer-Navigation-Btn-Prev');
	const nextBtn = document.getElementById('Footer-Navigation-Btn-Next');

	if (cIndex > 0) {
		prevBtn.href = '/' + urlBase + '/' + data[cIndex - 1]._id;
		prevBtn.firstChild.textContent = '< Prev';
		prevBtn.firstChild.disabled = false;
		prevBtn.firstChild.style.padding = '6px 10px';
	}

	if (cIndex < data.length - 1) {
		nextBtn.href = '/' + urlBase + '/' + data[cIndex + 1]._id;
		nextBtn.firstChild.textContent = 'Next >';
		nextBtn.firstChild.disabled = false;
		nextBtn.firstChild.style.padding = '6px 10px';
	}
}

// Register Service Worker
if ('serviceWorker' in navigator) {
	window.addEventListener('load', () => {
		navigator.serviceWorker
			.register('/static/serviceWorker.js')
			.then((reg) => console.log('Service Worker: Registered (Pages)'))
			.catch((err) => console.log(`Service Worker: Error: ${err}`));
	});
}

// Button POP-UP Animation

// For Document Navigation in Footer
setTimeout(() => {
	(function() {
		const conatiner = document.querySelector('.Bottom-Social-Btn-Container-00');
		console.log(conatiner.childNodes);

		cArr = Array.prototype.slice.call(conatiner.childNodes);

		for (let i = 0; i < cArr.length; i++) {
			const element = cArr[i];
			(function(e) {
				setTimeout(() => {
					if (e.style) {
						e.style.display = 'flex';
					}
				}, i * 100);
			})(element);
		}
	})();
}, 600);
