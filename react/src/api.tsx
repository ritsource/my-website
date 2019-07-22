import axios from 'axios';

export const serverAddress = process.env.NODE_ENV === 'production' ? 'https://ritwiksaha.com' : 'http://localhost:8080';

// Axios Instance
const api = axios.create({
	baseURL: serverAddress + '/api',
	headers: {
		'Content-Type': 'application/x-www-form-urlencoded',
		Accept: 'application/json'
	},
	withCredentials: true
});

export default api;
