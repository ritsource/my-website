const myAppThemes: any = {
	light: {
		'--text-color': '#424242',
		'--text-color-medium': '#666666',
		'--text-color-light': '#909090',
		'--background-color': '#ffffff',
		'--background-color-darker': '#f1f3f3',
		'--border-color': '#dddddd',
		'--theme-color': '#00ab6c'
	},
	dark: {
		'--text-color': '#ffffff',
		'--text-color-medium': '#e0e0e0',
		'--text-color-light': '#bbbbbb',
		'--background-color': '#464646',
		'--background-color-darker': '#323232',
		'--border-color': '#808080',
		'--theme-color': '#00ab6c'
	}
};

export default (isLight: boolean, setIsLight: (isLight: boolean) => void) => {
	const theme = isLight ? 'dark' : 'light';

	if (window) {
		const html = document.getElementsByTagName('html')[0];
		const obj = myAppThemes[theme];

		Object.keys(obj).map((key) => {
			html.style.setProperty(key, obj[key]);
		});
	}

	setIsLight(!isLight);
};
