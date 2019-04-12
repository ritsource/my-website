fetch('http://localhost:8080/public/blog/all')
	.then((response) => {
		return response.json();
	})
	.then((data) => {
		setNavigation(data);
	})
	.catch((error) => {
		console.log(error);
	});

// data
function setNavigation(data) {
	let linkSeg = window.location.href.split('/');
	let currentId = linkSeg.pop();
	if (currentId === '') {
		currentId = linkSeg.pop();
	}

	cIndex = data.findIndex(({ _id }) => _id === currentId);

	const prevBtn = document.getElementById('Footer-Navigation-Btn-Prev');
	const nextBtn = document.getElementById('Footer-Navigation-Btn-Next');

	if (cIndex > 0) {
		console.log(data[cIndex - 1]._id);

		prevBtn.href = 'http://localhost:8081/blog/' + data[cIndex - 1]._id;
		prevBtn.firstChild.disabled = false;
	}

	if (cIndex < data.length - 1) {
		nextBtn.href = 'http://localhost:8081/blog/' + data[cIndex + 1]._id;
		nextBtn.firstChild.disabled = false;
	}

	console.log(cIndex);
}
