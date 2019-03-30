import React, { useState, useEffect } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import './App.scss';
import api from './api';

import Header from './components/Header';
import HomePage from './pages/Home';
import LoginPage from './pages/Login';
import BlogsPage from './pages/Blogs';
import ProjectsPage from './pages/Projects';
import EachProjectPage from './pages/EachProject';
import EachBlogPage from './pages/EachBlog';
import LoadingPage from './pages/Loading';
import NotFoundPage from './pages/NotFound';

import { ProjectProvider } from './contexts/ProjectContext';
import { BlogProvider } from './contexts/BlogContext';

const App = () => {
	const [ auth, setAuth ]: any = useState(null);

	const checkAuth = async () => {
		try {
			const response = await api.get('/auth/current_user');
			setAuth(response.data);
		} catch (error) {
			setAuth(false);
		}
	};

	useEffect(() => {
		checkAuth();
	}, []);

	return (
		<div className="App">
			<ProjectProvider>
				<BlogProvider>
					<BrowserRouter>
						<Header />
						{auth ? (
							<Switch>
								<Route path="/" exact component={HomePage} />
								<Route path="/blogs" exact component={BlogsPage} />
								<Route path="/blogs/:blogId" exact component={EachBlogPage} />
								<Route path="/projects" exact component={ProjectsPage} />
								<Route path="/projects/:projectId" exact component={EachProjectPage} />
								<Route component={NotFoundPage} />
							</Switch>
						) : auth === null ? (
							<Route component={LoadingPage} />
						) : (
							<React.Fragment>
								<Route component={LoginPage} />
							</React.Fragment>
						)}
					</BrowserRouter>
				</BlogProvider>
			</ProjectProvider>
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
