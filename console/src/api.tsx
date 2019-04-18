import axios from 'axios';

// Axios Instance
const api = axios.create({
	baseURL: process.env.REACT_APP_API_URL + '/api',
	headers: {
		'Content-Type': 'application/x-www-form-urlencoded',
		Accept: 'application/json'
	},
	withCredentials: true
});

// if (process.env.NODE_ENV === "development") {

// }

export default api;
