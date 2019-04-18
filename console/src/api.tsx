import axios from 'axios';

export const apiBase =
	process.env.NODE_ENV === 'production' ? 'https://ritwiksaha.com/api' : process.env.REACT_APP_API_URL;

export const siteBase =
	process.env.NODE_ENV === 'production' ? 'https://ritwiksaha.com' : process.env.REACT_APP_WEBSITE_URL;

// Axios Instance
const api = axios.create({
	baseURL: apiBase,
	headers: {
		'Content-Type': 'application/x-www-form-urlencoded',
		Accept: 'application/json'
	},
	withCredentials: true
});

// if (process.env.NODE_ENV === "development") {

// }

export default api;
