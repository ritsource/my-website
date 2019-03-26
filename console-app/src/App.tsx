import React, { useState, useEffect } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import axios from 'axios';
import './App.scss';

import Header from './components/Header';
import HomePage from './pages/Home';
import LoginPage from './pages/Login';
import BlogsPage from './pages/Blogs';
import ProjectsPage from './pages/Projects';
import ImagesPage from './pages/Images';
import EachProjectPage from './pages/EachProject';
import LoadingPage from './pages/Loading';
import NotFoundPage from './pages/NotFound';

const App = () => {
	const [ auth, setAuth ]: any = useState(null);

	const checkAuth = async () => {
		try {
			// axios.defaults.withCredentials = true;

			const response = await axios.get(process.env.REACT_APP_API_URL + '/auth/current_user', {
				headers: {
					'Content-Type': 'application/x-www-form-urlencoded',
					Accept: 'application/json'
					// cookie: document.cookie,
				},
				withCredentials: true
			});

			setAuth(response.data);
			console.log(response);
		} catch (error) {
			setAuth(false);
			console.log('Error:', error);
		}
	};

	useEffect(() => {
		// setAuth()
		checkAuth();
	}, []);

	return (
		<div className="App">
			<BrowserRouter>
				<Header />
				<Switch>
					{auth ? (
						<React.Fragment>
							<Route path="/" exact component={HomePage} />
							<Route path="/blogs" exact component={BlogsPage} />
							<Route path="/projects" exact component={ProjectsPage} />
							<Route path="/images" exact component={ImagesPage} />
							<Route path="/project/:projectId" exact component={EachProjectPage} />
							<Route component={NotFoundPage} />
						</React.Fragment>
					) : auth === null ? (
						<Route component={LoadingPage} />
					) : (
						<React.Fragment>
							<Route component={LoginPage} />
							{/* <Route component={NotFoundPage} /> */}
						</React.Fragment>
					)}
				</Switch>
			</BrowserRouter>
		</div>
	);
};

export default App;

declare global {
	interface Window {
		getCookie: any;
	}
}

window.getCookie =
	window.getCookie ||
	((key: string) => {
		var val = document.cookie.match('(^|[^;]+)\\s*' + key + '\\s*=\\s*([^;]+)');
		return val ? val.pop() : '';
	});
