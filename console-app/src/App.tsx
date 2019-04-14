import React, { useState, useEffect } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import { connect } from 'react-redux';
import './App.scss';
import api from './api';

import { readProjects } from './actions/project_actions';
import { readBlogs } from './actions/blog_actions';

import Header from './components/Header';
import HomePage from './pages/Home';
import LoginPage from './pages/Login';
import EachProjectPage from './pages/EachProject';
import EachBlogPage from './pages/EachBlog';
import LoadingPage from './pages/Loading';
import NotFoundPage from './pages/NotFound';

import { ProjectProvider } from './contexts/ProjectContext';
import { BlogProvider } from './contexts/BlogContext';

const App = (props: any) => {
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
		props.readProjects();
		props.readBlogs();
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
								<Route path="/blogs/:blogId" exact component={EachBlogPage} />
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

const mapDispatchToProps = (dispatch: (x: any) => void) => ({
	readProjects: () => dispatch(readProjects()),
	readBlogs: () => dispatch(readBlogs())
});

export default connect(null, mapDispatchToProps)(App);

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
