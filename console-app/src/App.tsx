import React, { Component } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import './App.scss';

import Header from './Header';
import IndexPage from './pages/Index';
import BlogsPage from './pages/Blogs';
import ProjectsPage from './pages/Projects';
import ImagesPage from './pages/Images';
import NotFoundPage from './pages/NotFound';

class App extends Component {
	render() {
		return (
			<div className="App">
				<BrowserRouter>
					<Header />
					<Switch>
						<Route path="/" exact component={IndexPage} />
						<Route path="/blogs" exact component={BlogsPage} />
						<Route path="/projects" exact component={ProjectsPage} />
						<Route path="/images" exact component={ImagesPage} />
						<Route component={NotFoundPage} />
					</Switch>
				</BrowserRouter>
			</div>
		);
	}
}

export default App;
